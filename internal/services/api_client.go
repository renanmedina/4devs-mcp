package services

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"math"
	"mime/multipart"
	"net/http"
	"reflect"
	"strings"

	"github.com/renanmedina/4devs-mcp/observability"
)

type ApiClient[T any] struct {
	httpClient http.Client
	baseUrl    string
	authToken  string
	encodeType string
	logger     *slog.Logger
}

type ApiConfig struct {
	ApiUrl     string
	AuthToken  string
	EncodeType string
	LogEnabled bool
}

func NewApiClient[T any](config ApiConfig) ApiClient[T] {
	var logger *slog.Logger

	if config.LogEnabled {
		logger = observability.GetLogger()
	}

	encodeType := "json"
	if config.EncodeType != "" {
		encodeType = config.EncodeType
	}

	return ApiClient[T]{
		httpClient: http.Client{},
		encodeType: encodeType,
		baseUrl:    config.ApiUrl,
		authToken:  config.AuthToken,
		logger:     logger,
	}
}

func (c *ApiClient[T]) BuildUrl(path string, params map[string]interface{}, requestMethod string) string {
	url := fmt.Sprintf("%s%s", c.baseUrl, path)
	if requestMethod != http.MethodGet {
		return url
	}

	var paramsStrings []string

	for pname, pvalue := range params {
		paramsStrings = append(paramsStrings, fmt.Sprintf("%s=%s", pname, pvalue))
	}

	return fmt.Sprintf("%s?%s", url, strings.Join(paramsStrings, "&"))
}

func parseResult[T any](data []byte) (*T, error) {
	var resultData T
	err := json.Unmarshal(data, &resultData)

	if err != nil {
		return nil, err
	}

	return &resultData, nil
}

func bytesToType[T any](b []byte, order binary.ByteOrder) (T, error) {
	var value T
	switch v := any(&value).(type) {
	case *int8:
		*v = int8(b[0])
	case *int16:
		*v = int16(order.Uint16(b))
	case *int32:
		*v = int32(order.Uint32(b))
	case *int64:
		*v = int64(order.Uint64(b))
	case *uint8:
		*v = b[0]
	case *uint16:
		*v = order.Uint16(b)
	case *uint32:
		*v = order.Uint32(b)
	case *uint64:
		*v = order.Uint64(b)
	case *float32:
		*v = math.Float32frombits(order.Uint32(b))
	case *float64:
		*v = math.Float64frombits(order.Uint64(b))
	case *string:
		*v = string(b)
	default:
		rv := reflect.ValueOf(&value).Elem()
		if rv.Kind() == reflect.Struct {
			for i := 0; i < rv.NumField(); i++ {
				field := rv.Field(i)
				fieldType := field.Type()
				fieldSize := int(fieldType.Size())

				if len(b) < fieldSize {
					return value, errors.New("not enough bytes for struct field")
				}

				fieldValue, err := bytesToType[any](b[:fieldSize], order)
				if err != nil {
					return value, err
				}

				field.Set(reflect.ValueOf(fieldValue).Convert(fieldType))
				b = b[fieldSize:]
			}
		} else if rv.Kind() == reflect.Slice {
			slice := reflect.MakeSlice(rv.Type(), len(b), len(b))
			reflect.Copy(slice, reflect.ValueOf(b))
			rv.Set(slice)
		} else {
			return value, errors.New("unsupported type")
		}
	}
	return value, nil
}

func (client *ApiClient[T]) Get(path string, params map[string]interface{}, headers map[string]string) (*T, error) {
	response, err := client.performRequest(http.MethodGet, path, params, headers)
	if err != nil {
		return nil, err
	}
	return client.parseResponse(response, err)
}

func (client *ApiClient[T]) Post(path string, params map[string]interface{}, headers map[string]string) (*T, error) {
	response, err := client.performRequest(http.MethodPost, path, params, headers)
	if err != nil {
		return nil, err
	}
	return client.parseResponse(response, err)
}

func (client *ApiClient[T]) Put(path string, params map[string]interface{}, headers map[string]string) (*T, error) {
	response, err := client.performRequest(http.MethodPut, path, params, headers)
	if err != nil {
		return nil, err
	}
	return client.parseResponse(response, err)
}

func (client *ApiClient[T]) parseResponse(response *http.Response, err error) (*T, error) {
	if err != nil {
		return nil, err
	}

	bodyData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if strings.ToLower(client.encodeType) != "json" {
		value, err := bytesToType[T](bodyData, binary.LittleEndian)
		if err != nil {
			return nil, err
		}
		return &value, nil
	}

	parsed, err := parseResult[T](bodyData)

	if err != nil {
		return nil, err
	}

	return parsed, nil
}

func (client *ApiClient[T]) performRequest(requestMethod string, path string, params map[string]interface{}, headers map[string]string) (*http.Response, error) {
	url := client.BuildUrl(path, params, requestMethod)
	paramsBuffer := bytes.NewBuffer(make([]byte, 0))

	headers["Accept"] = "*/*"
	headers["Content-Type"] = "application/json"

	if len(params) > 0 {
		if client.encodeType == "json" {
			bodyParams, err := json.Marshal(params)
			if err != nil {
				return nil, err
			}

			paramsBuffer = bytes.NewBuffer(bodyParams)
		} else {
			writer := multipart.NewWriter(paramsBuffer)
			for param, paramVal := range params {
				_ = writer.WriteField(param, paramVal.(string))
			}

			err := writer.Close()
			if err != nil {
				client.log(err.Error())
				return nil, err
			}

			headers["Content-Type"] = writer.FormDataContentType()
		}
	}

	request, err := http.NewRequest(requestMethod, url, paramsBuffer)

	if err != nil {
		return nil, err
	}

	if client.authToken != "" {
		headers["Authorization"] = client.authToken
	}

	for headerKey, headerValue := range headers {
		request.Header.Add(headerKey, headerValue)
	}

	client.log(fmt.Sprintf("[%s] Sending http request to %s", requestMethod, url))
	response, err := client.httpClient.Do(request)
	client.log(fmt.Sprintf("Response Status: %s", response.Status))
	client.log(fmt.Sprintf("Response StatusCode: %d", response.StatusCode))
	return response, err
}

func (client *ApiClient[T]) log(message string) {
	if client.logger != nil {
		typeName := reflect.TypeFor[T]().Name()
		formattedLog := fmt.Sprintf("[%s] %s", typeName, message)
		client.logger.Info(formattedLog)
	}
}

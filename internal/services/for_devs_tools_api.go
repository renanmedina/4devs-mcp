package services

type ForDevsToolsApi struct {
	ApiClient[string]
}

func NewForDevsToolsClient(logEnabled bool) ForDevsToolsApi {
	return ForDevsToolsApi{
		NewApiClient[string](
			ApiConfig{
				ApiUrl:     "https://www.4devs.com.br/ferramentas_online.php",
				EncodeType: "form-data",
				LogEnabled: logEnabled,
			},
		),
	}
}

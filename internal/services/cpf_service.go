package services

import (
	"log/slog"

	"github.com/renanmedina/4devs-mcp/observability"
)

type CpfService struct {
	logger    *slog.Logger
	apiClient ApiClient[string]
}

func NewCpfService(logEnabled bool) CpfService {
	return CpfService{
		observability.GetLogger(),
		NewApiClient[string](
			ApiConfig{
				ApiUrl:     "https://www.4devs.com.br",
				EncodeType: "form-data",
				LogEnabled: logEnabled,
			},
		),
	}
}

func (s CpfService) Generate(formatted bool, uf string) (*string, error) {
	formatted_string := "N"
	if formatted {
		formatted_string = "S"
	}

	return s.apiClient.Post("/ferramentas_online.php", map[string]interface{}{
		"acao":       "gerar_cpf",
		"pontuacao":  formatted_string,
		"cpf_estado": uf,
	}, map[string]string{})
}

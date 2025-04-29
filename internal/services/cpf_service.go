package services

type CpfService struct {
	client ForDevsToolsApi
}

func NewCpfService(logEnabled bool) CpfService {
	return CpfService{NewForDevsToolsClient(logEnabled)}
}

func (s CpfService) Generate(formatted bool, uf string) (*string, error) {
	formatted_string := "N"
	if formatted {
		formatted_string = "S"
	}

	return s.client.Post("", map[string]interface{}{
		"acao":       "gerar_cpf",
		"pontuacao":  formatted_string,
		"cpf_estado": uf,
	}, map[string]string{})
}

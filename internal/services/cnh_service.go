package services

type CnhService struct {
	client ForDevsToolsApi
}

func NewCnhService(logEnabled bool) CnhService {
	return CnhService{NewForDevsToolsClient(logEnabled)}
}

func (s CnhService) Generate() (*string, error) {
	return s.client.Post("", map[string]interface{}{
		"acao": "gerar_cnh",
	}, map[string]string{})
}

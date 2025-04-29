package services

type CpfService struct {
	client ForDevsToolsApi
}

func NewCpfService(logEnabled bool) CpfService {
	return CpfService{NewForDevsToolsClient(logEnabled)}
}

func (s CpfService) Generate(formatted bool, uf string) (*string, error) {
	return s.client.Post("", map[string]interface{}{
		"acao":       "gerar_cpf",
		"pontuacao":  formattedParam(formatted),
		"cpf_estado": uf,
	}, map[string]string{})
}

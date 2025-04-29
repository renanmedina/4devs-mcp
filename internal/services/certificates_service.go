package services

type CertificatesService struct {
	client ForDevsToolsApi
}

const (
	ANY_CERTIFICATE                = "Indiferente"
	BIRTHDATE_CERTIFICATE          = "nascimento"
	MARRIAGE_CERTIFICATE           = "casamento"
	RELIGIOUS_MARRIAGE_CERTIFICATE = "casamento_religioso"
	DEATH_CERTIFICATE              = "obito"
)

func NewCertificatesService(logEnabled bool) CertificatesService {
	return CertificatesService{NewForDevsToolsClient(logEnabled)}
}

func (s CertificatesService) Generate(formatted bool, certificateType string) (*string, error) {
	if certificateType == "" {
		certificateType = ANY_CERTIFICATE
	}

	return s.client.Post("", map[string]interface{}{
		"acao":          "gerador_certidao",
		"pontuacao":     formattedParam(formatted),
		"tipo_certidao": certificateType,
	}, map[string]string{})
}

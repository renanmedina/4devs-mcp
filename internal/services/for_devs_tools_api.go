package services

type ForDevsToolsApi struct {
	ApiClient[string]
}

const (
	STATE_AC = "AC"
	STATE_AL = "AL"
	STATE_AP = "AP"
	STATE_AM = "AM"
	STATE_BA = "BA"
	STATE_CE = "CE"
	STATE_DF = "DF"
	STATE_ES = "ES"
	STATE_GO = "GO"
	STATE_MA = "MA"
	STATE_MG = "MG"
	STATE_MT = "MT"
	STATE_MS = "MS"
	STATE_PA = "PA"
	STATE_PB = "PB"
	STATE_RJ = "RJ"
	STATE_SP = "SP"
	STATE_PR = "PR"
	STATE_RS = "RS"
	STATE_SC = "SC"
	STATE_SE = "SE"
	STATE_TO = "TO"
	STATE_PI = "PI"
	STATE_RN = "RN"
	STATE_RO = "RO"
	STATE_RR = "RR"
)

var STATES_OPTIONS []string = []string{
	STATE_AC, STATE_AL, STATE_AP, STATE_AM, STATE_BA, STATE_CE, STATE_DF, STATE_ES, STATE_GO, STATE_MA, STATE_MG, STATE_MT, STATE_MS, STATE_PA, STATE_PB, STATE_RJ, STATE_SP, STATE_PR, STATE_RS, STATE_SC, STATE_SE, STATE_TO, STATE_PI, STATE_RN, STATE_RO, STATE_RR,
}

var BANKS_MAP = map[string]string{
	"Banco do Brasil": "2",
	"Bradesco":        "121",
	"Citibank":        "85",
	"Itau":            "120",
	"Santander":       "151",
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

func formattedParam(formatted bool) string {
	if formatted {
		return "S"
	}

	return "N"
}

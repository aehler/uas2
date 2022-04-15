package config

type Config struct {
	KafkaAddressPool []string
	RestAddress string
	RestCreds RestCreds `json:"-"`
}

type RestCreds struct {
	ExtSystem string `json:"externalSystem"`
	Login string `json:"login"`
	Password string `json:"password"`
}


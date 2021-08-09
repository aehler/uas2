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

var Data = Config{
	KafkaAddressPool: []string{"10.204.192.111:9192","10.204.192.112:9192","10.204.192.113:9192"},
	RestAddress: "https://eaist2-f.mos.ru/module/uas-2/api/v1",
	RestCreds :RestCreds{
		ExtSystem: "DIT_test_http",
		Login: "DIT_test_http",
		Password: "DIT_test_http",
	},

}
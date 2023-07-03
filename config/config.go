package config

import "github.com/kelseyhightower/envconfig"

const (
	NAMESPACE = "palpalo"
)

type ValorantAPI struct {
	ClientID    string `envconfig:"client_id"`
	RedirectURI string `envconfig:"redirect_uri"`
	UserAgent   string `envconfig:"user_agent"`
}

type Configuration struct {
	ValorantAPI ValorantAPI `envconfig:"valorant_api"`
}

func Get() Configuration {
	var cfg Configuration

	err := envconfig.Process(NAMESPACE, &cfg)

	if err != nil {
		return Configuration{}
	}

	return cfg
}

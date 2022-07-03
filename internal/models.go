package internal

import "github.com/mitchellh/mapstructure"

type EnvMap map[string]string

type RunningContainer struct {
	Name  string
	AppID string
}

type TraefikConfig struct {
	EnableHttps bool   `mapstructure:"enableHttps"`
	CertEmail   string `mapstructure:"certEmail"`
}

func (p Plugin) ConfigForTraefik() TraefikConfig {
	var parsed TraefikConfig
	if p.Config != nil {
		mapstructure.Decode(p.Config, &parsed)
	}
	return parsed
}

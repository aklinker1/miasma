package internal

import "github.com/mitchellh/mapstructure"

type EnvMap map[string]string

type TraefikConfig struct {
	EnableHttps bool   `mapstructure:"enableHttps"`
	CertsEmail  string `mapstructure:"certsEmail"`
	CertsDir    string `mapstructure:"certsDir"`
}

func (p Plugin) ConfigForTraefik() TraefikConfig {
	var parsed TraefikConfig
	if p.Config != nil {
		mapstructure.Decode(p.Config, &parsed)
	}
	return parsed
}

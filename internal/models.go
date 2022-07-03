package internal

type EnvMap map[string]string

type RunningContainer struct {
	Name  string
	AppID string
}

type TraefikConfig struct {
	EnableHttps bool   `mapstructure:"enableHttps"`
	CertEmail   string `mapstructure:"certEmail"`
}

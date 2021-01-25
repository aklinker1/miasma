package client

import "github.com/go-openapi/strfmt"

func NewClientWith(host string) *Miasma {
	return NewHTTPClientWithConfig(
		strfmt.Default,
		&TransportConfig{
			Host:     host,
			BasePath: "/",
		},
	)
}

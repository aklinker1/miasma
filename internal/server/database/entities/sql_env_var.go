package entities

import "github.com/go-openapi/strfmt"

type SQLEnvVar struct {
	AppID strfmt.UUID4 `gorm:"primaryKey"`
	Key   string       `gorm:"primaryKey"`
	Value string
}

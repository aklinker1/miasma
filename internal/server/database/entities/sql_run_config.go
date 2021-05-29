package entities

import (
	"github.com/aklinker1/miasma/internal/server/utils/custom_formats"
	"github.com/go-openapi/strfmt"
)

type SQLRunConfig struct {
	AppID          strfmt.UUID4               `gorm:"primaryKey"`
	Command        custom_formats.StringArray `gorm:"type:blob"`
	ImageDigest    string
	Networks       custom_formats.StringArray `gorm:"type:blob"`
	Placement      custom_formats.StringArray `gorm:"type:blob"`
	PublishedPorts custom_formats.UInt32Array `gorm:"type:blob"`
	TargetPorts    custom_formats.UInt32Array `gorm:"type:blob"`
	Volumes        []byte                     `gorm:"type:blob"`
}

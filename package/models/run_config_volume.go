// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RunConfigVolume run config volume
//
// swagger:model RunConfigVolume
type RunConfigVolume struct {

	// The volume name or directory on the host that the data is stored in
	Source string `json:"Source,omitempty"`

	// The path inside the container that the data is served from
	Target string `json:"Target,omitempty"`
}

// Validate validates this run config volume
func (m *RunConfigVolume) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RunConfigVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RunConfigVolume) UnmarshalBinary(b []byte) error {
	var res RunConfigVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

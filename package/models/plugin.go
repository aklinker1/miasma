// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Plugin plugin
//
// swagger:model Plugin
type Plugin struct {

	// Whether or not the plugin is installed
	// Required: true
	Installed bool `json:"installed"`

	// The plugin's name. It can be used to install a plugin
	// Required: true
	Name string `json:"name" gorm:"primaryKey"`
}

// Validate validates this plugin
func (m *Plugin) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstalled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Plugin) validateInstalled(formats strfmt.Registry) error {

	if err := validate.Required("installed", "body", bool(m.Installed)); err != nil {
		return err
	}

	return nil
}

func (m *Plugin) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Plugin) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Plugin) UnmarshalBinary(b []byte) error {
	var res Plugin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

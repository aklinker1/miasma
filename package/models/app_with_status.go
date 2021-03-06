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

// AppWithStatus app with status
//
// swagger:model AppWithStatus
type AppWithStatus struct {

	// A simple label to track what apps are related
	Group string `json:"group,omitempty"`

	// The number of instances running vs what should be running
	Instances string `json:"instances,omitempty"`

	// The apps name, used in the CLI with the `-a|--app` flag
	// Required: true
	Name string `json:"name"`

	// The published ports for the app
	// Required: true
	Ports []string `json:"ports"`

	// If the app has routing, a simple string representing that route
	Routing string `json:"routing,omitempty"`

	// Whether or not the application is running, stopped, or starting up
	// Required: true
	Status string `json:"status"`
}

// Validate validates this app with status
func (m *AppWithStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppWithStatus) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *AppWithStatus) validatePorts(formats strfmt.Registry) error {

	if err := validate.Required("ports", "body", m.Ports); err != nil {
		return err
	}

	return nil
}

func (m *AppWithStatus) validateStatus(formats strfmt.Registry) error {

	if err := validate.RequiredString("status", "body", string(m.Status)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppWithStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppWithStatus) UnmarshalBinary(b []byte) error {
	var res AppWithStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

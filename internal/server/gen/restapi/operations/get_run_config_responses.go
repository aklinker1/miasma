// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/aklinker1/miasma/package/models"
)

// GetRunConfigOKCode is the HTTP code returned for type GetRunConfigOK
const GetRunConfigOKCode int = 200

/*GetRunConfigOK OK

swagger:response getRunConfigOK
*/
type GetRunConfigOK struct {

	/*
	  In: Body
	*/
	Payload *models.RunConfig `json:"body,omitempty"`
}

// NewGetRunConfigOK creates GetRunConfigOK with default headers values
func NewGetRunConfigOK() *GetRunConfigOK {

	return &GetRunConfigOK{}
}

// WithPayload adds the payload to the get run config o k response
func (o *GetRunConfigOK) WithPayload(payload *models.RunConfig) *GetRunConfigOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get run config o k response
func (o *GetRunConfigOK) SetPayload(payload *models.RunConfig) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRunConfigOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRunConfigNotFoundCode is the HTTP code returned for type GetRunConfigNotFound
const GetRunConfigNotFoundCode int = 404

/*GetRunConfigNotFound Not Found

swagger:response getRunConfigNotFound
*/
type GetRunConfigNotFound struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetRunConfigNotFound creates GetRunConfigNotFound with default headers values
func NewGetRunConfigNotFound() *GetRunConfigNotFound {

	return &GetRunConfigNotFound{}
}

// WithPayload adds the payload to the get run config not found response
func (o *GetRunConfigNotFound) WithPayload(payload string) *GetRunConfigNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get run config not found response
func (o *GetRunConfigNotFound) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRunConfigNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetRunConfigDefault Unknown Error

swagger:response getRunConfigDefault
*/
type GetRunConfigDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetRunConfigDefault creates GetRunConfigDefault with default headers values
func NewGetRunConfigDefault(code int) *GetRunConfigDefault {
	if code <= 0 {
		code = 500
	}

	return &GetRunConfigDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get run config default response
func (o *GetRunConfigDefault) WithStatusCode(code int) *GetRunConfigDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get run config default response
func (o *GetRunConfigDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get run config default response
func (o *GetRunConfigDefault) WithPayload(payload string) *GetRunConfigDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get run config default response
func (o *GetRunConfigDefault) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRunConfigDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetAppEnvOKCode is the HTTP code returned for type GetAppEnvOK
const GetAppEnvOKCode int = 200

/*GetAppEnvOK OK

swagger:response getAppEnvOK
*/
type GetAppEnvOK struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetAppEnvOK creates GetAppEnvOK with default headers values
func NewGetAppEnvOK() *GetAppEnvOK {

	return &GetAppEnvOK{}
}

// WithPayload adds the payload to the get app env o k response
func (o *GetAppEnvOK) WithPayload(payload interface{}) *GetAppEnvOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get app env o k response
func (o *GetAppEnvOK) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAppEnvOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetAppEnvNotFoundCode is the HTTP code returned for type GetAppEnvNotFound
const GetAppEnvNotFoundCode int = 404

/*GetAppEnvNotFound Not Found

swagger:response getAppEnvNotFound
*/
type GetAppEnvNotFound struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetAppEnvNotFound creates GetAppEnvNotFound with default headers values
func NewGetAppEnvNotFound() *GetAppEnvNotFound {

	return &GetAppEnvNotFound{}
}

// WithPayload adds the payload to the get app env not found response
func (o *GetAppEnvNotFound) WithPayload(payload string) *GetAppEnvNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get app env not found response
func (o *GetAppEnvNotFound) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAppEnvNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetAppEnvDefault Unknown Error

swagger:response getAppEnvDefault
*/
type GetAppEnvDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetAppEnvDefault creates GetAppEnvDefault with default headers values
func NewGetAppEnvDefault(code int) *GetAppEnvDefault {
	if code <= 0 {
		code = 500
	}

	return &GetAppEnvDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get app env default response
func (o *GetAppEnvDefault) WithStatusCode(code int) *GetAppEnvDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get app env default response
func (o *GetAppEnvDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get app env default response
func (o *GetAppEnvDefault) WithPayload(payload string) *GetAppEnvDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get app env default response
func (o *GetAppEnvDefault) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAppEnvDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

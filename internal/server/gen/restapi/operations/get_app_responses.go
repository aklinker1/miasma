// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/aklinker1/miasma/package/models"
)

// GetAppOKCode is the HTTP code returned for type GetAppOK
const GetAppOKCode int = 200

/*GetAppOK OK

swagger:response getAppOK
*/
type GetAppOK struct {

	/*
	  In: Body
	*/
	Payload *models.App `json:"body,omitempty"`
}

// NewGetAppOK creates GetAppOK with default headers values
func NewGetAppOK() *GetAppOK {

	return &GetAppOK{}
}

// WithPayload adds the payload to the get app o k response
func (o *GetAppOK) WithPayload(payload *models.App) *GetAppOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get app o k response
func (o *GetAppOK) SetPayload(payload *models.App) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAppOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAppNotFoundCode is the HTTP code returned for type GetAppNotFound
const GetAppNotFoundCode int = 404

/*GetAppNotFound Not Found

swagger:response getAppNotFound
*/
type GetAppNotFound struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetAppNotFound creates GetAppNotFound with default headers values
func NewGetAppNotFound() *GetAppNotFound {

	return &GetAppNotFound{}
}

// WithPayload adds the payload to the get app not found response
func (o *GetAppNotFound) WithPayload(payload string) *GetAppNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get app not found response
func (o *GetAppNotFound) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAppNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

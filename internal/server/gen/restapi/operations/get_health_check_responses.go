// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/aklinker1/miasma/internal/server/gen/models"
)

// GetHealthCheckOKCode is the HTTP code returned for type GetHealthCheckOK
const GetHealthCheckOKCode int = 200

/*GetHealthCheckOK OK

swagger:response getHealthCheckOK
*/
type GetHealthCheckOK struct {

	/*
	  In: Body
	*/
	Payload *models.Health `json:"body,omitempty"`
}

// NewGetHealthCheckOK creates GetHealthCheckOK with default headers values
func NewGetHealthCheckOK() *GetHealthCheckOK {

	return &GetHealthCheckOK{}
}

// WithPayload adds the payload to the get health check o k response
func (o *GetHealthCheckOK) WithPayload(payload *models.Health) *GetHealthCheckOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get health check o k response
func (o *GetHealthCheckOK) SetPayload(payload *models.Health) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHealthCheckOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

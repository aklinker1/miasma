// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/aklinker1/miasma/package/models"
)

// ListAppsReader is a Reader for the ListApps structure.
type ListAppsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAppsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAppsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListAppsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAppsOK creates a ListAppsOK with default headers values
func NewListAppsOK() *ListAppsOK {
	return &ListAppsOK{}
}

/*ListAppsOK handles this case with default header values.

OK
*/
type ListAppsOK struct {
	Payload []*models.AppWithStatus
}

func (o *ListAppsOK) Error() string {
	return fmt.Sprintf("[GET /api/apps][%d] listAppsOK  %+v", 200, o.Payload)
}

func (o *ListAppsOK) GetPayload() []*models.AppWithStatus {
	return o.Payload
}

func (o *ListAppsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAppsDefault creates a ListAppsDefault with default headers values
func NewListAppsDefault(code int) *ListAppsDefault {
	return &ListAppsDefault{
		_statusCode: code,
	}
}

/*ListAppsDefault handles this case with default header values.

Unknown Error
*/
type ListAppsDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the list apps default response
func (o *ListAppsDefault) Code() int {
	return o._statusCode
}

func (o *ListAppsDefault) Error() string {
	return fmt.Sprintf("[GET /api/apps][%d] listApps default  %+v", o._statusCode, o.Payload)
}

func (o *ListAppsDefault) GetPayload() string {
	return o.Payload
}

func (o *ListAppsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

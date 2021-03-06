// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// GetOperatingSystemsReader is a Reader for the GetOperatingSystems structure.
type GetOperatingSystemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOperatingSystemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOperatingSystemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetOperatingSystemsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOperatingSystemsOK creates a GetOperatingSystemsOK with default headers values
func NewGetOperatingSystemsOK() *GetOperatingSystemsOK {
	return &GetOperatingSystemsOK{}
}

/*GetOperatingSystemsOK handles this case with default header values.

dummy
*/
type GetOperatingSystemsOK struct {
	Payload *models.OperatingSystemCollection
}

func (o *GetOperatingSystemsOK) Error() string {
	return fmt.Sprintf("[GET /operating_systems][%d] getOperatingSystemsOK  %+v", 200, o.Payload)
}

func (o *GetOperatingSystemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OperatingSystemCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOperatingSystemsInternalServerError creates a GetOperatingSystemsInternalServerError with default headers values
func NewGetOperatingSystemsInternalServerError() *GetOperatingSystemsInternalServerError {
	return &GetOperatingSystemsInternalServerError{}
}

/*GetOperatingSystemsInternalServerError handles this case with default header values.

error
*/
type GetOperatingSystemsInternalServerError struct {
	Payload *models.Riaaserror
}

func (o *GetOperatingSystemsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /operating_systems][%d] getOperatingSystemsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOperatingSystemsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

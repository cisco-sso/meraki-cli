// Code generated by go-swagger; DO NOT EDIT.

package change_log

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetOrganizationConfigurationChangesReader is a Reader for the GetOrganizationConfigurationChanges structure.
type GetOrganizationConfigurationChangesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationConfigurationChangesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationConfigurationChangesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOrganizationConfigurationChangesOK creates a GetOrganizationConfigurationChangesOK with default headers values
func NewGetOrganizationConfigurationChangesOK() *GetOrganizationConfigurationChangesOK {
	return &GetOrganizationConfigurationChangesOK{}
}

/*GetOrganizationConfigurationChangesOK handles this case with default header values.

Successful operation
*/
type GetOrganizationConfigurationChangesOK struct {
	Payload interface{}
}

func (o *GetOrganizationConfigurationChangesOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/configurationChanges][%d] getOrganizationConfigurationChangesOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationConfigurationChangesOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetOrganizationConfigurationChangesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
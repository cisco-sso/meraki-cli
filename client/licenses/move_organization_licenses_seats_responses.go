// Code generated by go-swagger; DO NOT EDIT.

package licenses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// MoveOrganizationLicensesSeatsReader is a Reader for the MoveOrganizationLicensesSeats structure.
type MoveOrganizationLicensesSeatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MoveOrganizationLicensesSeatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMoveOrganizationLicensesSeatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMoveOrganizationLicensesSeatsOK creates a MoveOrganizationLicensesSeatsOK with default headers values
func NewMoveOrganizationLicensesSeatsOK() *MoveOrganizationLicensesSeatsOK {
	return &MoveOrganizationLicensesSeatsOK{}
}

/*MoveOrganizationLicensesSeatsOK handles this case with default header values.

Successful operation
*/
type MoveOrganizationLicensesSeatsOK struct {
	Payload interface{}
}

func (o *MoveOrganizationLicensesSeatsOK) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationId}/licenses/moveSeats][%d] moveOrganizationLicensesSeatsOK  %+v", 200, o.Payload)
}

func (o *MoveOrganizationLicensesSeatsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *MoveOrganizationLicensesSeatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
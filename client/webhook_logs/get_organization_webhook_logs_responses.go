// Code generated by go-swagger; DO NOT EDIT.

package webhook_logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetOrganizationWebhookLogsReader is a Reader for the GetOrganizationWebhookLogs structure.
type GetOrganizationWebhookLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationWebhookLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationWebhookLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOrganizationWebhookLogsOK creates a GetOrganizationWebhookLogsOK with default headers values
func NewGetOrganizationWebhookLogsOK() *GetOrganizationWebhookLogsOK {
	return &GetOrganizationWebhookLogsOK{}
}

/*GetOrganizationWebhookLogsOK handles this case with default header values.

Successful operation
*/
type GetOrganizationWebhookLogsOK struct {
	Payload interface{}
}

func (o *GetOrganizationWebhookLogsOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/webhookLogs][%d] getOrganizationWebhookLogsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationWebhookLogsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetOrganizationWebhookLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

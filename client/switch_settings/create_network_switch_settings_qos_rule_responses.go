// Code generated by go-swagger; DO NOT EDIT.

package switch_settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateNetworkSwitchSettingsQosRuleReader is a Reader for the CreateNetworkSwitchSettingsQosRule structure.
type CreateNetworkSwitchSettingsQosRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNetworkSwitchSettingsQosRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateNetworkSwitchSettingsQosRuleCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateNetworkSwitchSettingsQosRuleCreated creates a CreateNetworkSwitchSettingsQosRuleCreated with default headers values
func NewCreateNetworkSwitchSettingsQosRuleCreated() *CreateNetworkSwitchSettingsQosRuleCreated {
	return &CreateNetworkSwitchSettingsQosRuleCreated{}
}

/*CreateNetworkSwitchSettingsQosRuleCreated handles this case with default header values.

Successful operation
*/
type CreateNetworkSwitchSettingsQosRuleCreated struct {
	Payload interface{}
}

func (o *CreateNetworkSwitchSettingsQosRuleCreated) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/switch/settings/qosRules][%d] createNetworkSwitchSettingsQosRuleCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkSwitchSettingsQosRuleCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateNetworkSwitchSettingsQosRuleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
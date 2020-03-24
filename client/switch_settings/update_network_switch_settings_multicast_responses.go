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

// UpdateNetworkSwitchSettingsMulticastReader is a Reader for the UpdateNetworkSwitchSettingsMulticast structure.
type UpdateNetworkSwitchSettingsMulticastReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkSwitchSettingsMulticastReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkSwitchSettingsMulticastOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateNetworkSwitchSettingsMulticastOK creates a UpdateNetworkSwitchSettingsMulticastOK with default headers values
func NewUpdateNetworkSwitchSettingsMulticastOK() *UpdateNetworkSwitchSettingsMulticastOK {
	return &UpdateNetworkSwitchSettingsMulticastOK{}
}

/*UpdateNetworkSwitchSettingsMulticastOK handles this case with default header values.

Successful operation
*/
type UpdateNetworkSwitchSettingsMulticastOK struct {
	Payload interface{}
}

func (o *UpdateNetworkSwitchSettingsMulticastOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/switch/settings/multicast][%d] updateNetworkSwitchSettingsMulticastOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkSwitchSettingsMulticastOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkSwitchSettingsMulticastOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
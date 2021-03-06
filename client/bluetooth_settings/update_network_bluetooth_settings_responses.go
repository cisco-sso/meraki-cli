// Code generated by go-swagger; DO NOT EDIT.

package bluetooth_settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateNetworkBluetoothSettingsReader is a Reader for the UpdateNetworkBluetoothSettings structure.
type UpdateNetworkBluetoothSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkBluetoothSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkBluetoothSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateNetworkBluetoothSettingsOK creates a UpdateNetworkBluetoothSettingsOK with default headers values
func NewUpdateNetworkBluetoothSettingsOK() *UpdateNetworkBluetoothSettingsOK {
	return &UpdateNetworkBluetoothSettingsOK{}
}

/*UpdateNetworkBluetoothSettingsOK handles this case with default header values.

Successful operation
*/
type UpdateNetworkBluetoothSettingsOK struct {
	Payload interface{}
}

func (o *UpdateNetworkBluetoothSettingsOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/bluetoothSettings][%d] updateNetworkBluetoothSettingsOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkBluetoothSettingsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkBluetoothSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

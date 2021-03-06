// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkDeviceUplinkReader is a Reader for the GetNetworkDeviceUplink structure.
type GetNetworkDeviceUplinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkDeviceUplinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkDeviceUplinkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNetworkDeviceUplinkOK creates a GetNetworkDeviceUplinkOK with default headers values
func NewGetNetworkDeviceUplinkOK() *GetNetworkDeviceUplinkOK {
	return &GetNetworkDeviceUplinkOK{}
}

/*GetNetworkDeviceUplinkOK handles this case with default header values.

Successful operation
*/
type GetNetworkDeviceUplinkOK struct {
	Payload interface{}
}

func (o *GetNetworkDeviceUplinkOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/devices/{serial}/uplink][%d] getNetworkDeviceUplinkOK  %+v", 200, o.Payload)
}

func (o *GetNetworkDeviceUplinkOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkDeviceUplinkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

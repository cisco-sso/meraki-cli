// Code generated by go-swagger; DO NOT EDIT.

package s_m

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// LockNetworkSmDevicesReader is a Reader for the LockNetworkSmDevices structure.
type LockNetworkSmDevicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LockNetworkSmDevicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLockNetworkSmDevicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLockNetworkSmDevicesOK creates a LockNetworkSmDevicesOK with default headers values
func NewLockNetworkSmDevicesOK() *LockNetworkSmDevicesOK {
	return &LockNetworkSmDevicesOK{}
}

/*LockNetworkSmDevicesOK handles this case with default header values.

Successful operation
*/
type LockNetworkSmDevicesOK struct {
	Payload interface{}
}

func (o *LockNetworkSmDevicesOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/sm/devices/lock][%d] lockNetworkSmDevicesOK  %+v", 200, o.Payload)
}

func (o *LockNetworkSmDevicesOK) GetPayload() interface{} {
	return o.Payload
}

func (o *LockNetworkSmDevicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateNetworkSiteToSiteVpnReader is a Reader for the UpdateNetworkSiteToSiteVpn structure.
type UpdateNetworkSiteToSiteVpnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkSiteToSiteVpnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkSiteToSiteVpnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateNetworkSiteToSiteVpnOK creates a UpdateNetworkSiteToSiteVpnOK with default headers values
func NewUpdateNetworkSiteToSiteVpnOK() *UpdateNetworkSiteToSiteVpnOK {
	return &UpdateNetworkSiteToSiteVpnOK{}
}

/*UpdateNetworkSiteToSiteVpnOK handles this case with default header values.

Successful operation
*/
type UpdateNetworkSiteToSiteVpnOK struct {
	Payload interface{}
}

func (o *UpdateNetworkSiteToSiteVpnOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/siteToSiteVpn][%d] updateNetworkSiteToSiteVpnOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkSiteToSiteVpnOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkSiteToSiteVpnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
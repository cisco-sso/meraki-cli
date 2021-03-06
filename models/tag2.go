// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Tag2 Tag2
//
// swagger:model Tag2
type Tag2 struct {

	// The privilege of the SAML administrator on the tag
	Access string `json:"access,omitempty"`

	// The name of the tag
	Tag string `json:"tag,omitempty"`
}

// Validate validates this tag2
func (m *Tag2) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Tag2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Tag2) UnmarshalBinary(b []byte) error {
	var res Tag2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// EncryptionMode EncryptionMode
//
// The psk encryption mode for the SSID ('wep' or 'wpa'). This param is only valid if the authMode is 'psk'
//
// swagger:model EncryptionMode
type EncryptionMode string

const (

	// EncryptionModeWep captures enum value "wep"
	EncryptionModeWep EncryptionMode = "wep"

	// EncryptionModeWpa captures enum value "wpa"
	EncryptionModeWpa EncryptionMode = "wpa"
)

// for schema
var encryptionModeEnum []interface{}

func init() {
	var res []EncryptionMode
	if err := json.Unmarshal([]byte(`["wep","wpa"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		encryptionModeEnum = append(encryptionModeEnum, v)
	}
}

func (m EncryptionMode) validateEncryptionModeEnum(path, location string, value EncryptionMode) error {
	if err := validate.Enum(path, location, value, encryptionModeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this encryption mode
func (m EncryptionMode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEncryptionModeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
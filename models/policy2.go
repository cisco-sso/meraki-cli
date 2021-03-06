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

// Policy2 Policy2
//
// The policy applied to matching traffic. Must be 'deny'.
//
// swagger:model Policy2
type Policy2 string

const (

	// Policy2Deny captures enum value "deny"
	Policy2Deny Policy2 = "deny"
)

// for schema
var policy2Enum []interface{}

func init() {
	var res []Policy2
	if err := json.Unmarshal([]byte(`["deny"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		policy2Enum = append(policy2Enum, v)
	}
}

func (m Policy2) validatePolicy2Enum(path, location string, value Policy2) error {
	if err := validate.Enum(path, location, value, policy2Enum); err != nil {
		return err
	}
	return nil
}

// Validate validates this policy2
func (m Policy2) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePolicy2Enum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

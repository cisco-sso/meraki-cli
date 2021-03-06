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

// Access Access
//
// A string indicating the rule for which IPs are allowed to use the specified service. Acceptable values are "blocked" (no remote IPs can access the service), "restricted" (only whitelisted IPs can access the service), and "unrestriced" (any remote IP can access the service). This field is required
//
// swagger:model Access
type Access string

const (

	// AccessBlocked captures enum value "blocked"
	AccessBlocked Access = "blocked"

	// AccessRestricted captures enum value "restricted"
	AccessRestricted Access = "restricted"

	// AccessUnrestricted captures enum value "unrestricted"
	AccessUnrestricted Access = "unrestricted"
)

// for schema
var accessEnum []interface{}

func init() {
	var res []Access
	if err := json.Unmarshal([]byte(`["blocked","restricted","unrestricted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accessEnum = append(accessEnum, v)
	}
}

func (m Access) validateAccessEnum(path, location string, value Access) error {
	if err := validate.Enum(path, location, value, accessEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this access
func (m Access) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAccessEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

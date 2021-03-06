// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BlockedURLCategories BlockedUrlCategories
//
// Settings for blacklisted URL categories
//
// swagger:model BlockedUrlCategories
type BlockedURLCategories struct {

	// A list of URL categories to block
	Categories []string `json:"categories"`

	// settings
	Settings Settings4 `json:"settings,omitempty"`
}

// Validate validates this blocked Url categories
func (m *BlockedURLCategories) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BlockedURLCategories) validateSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.Settings) { // not required
		return nil
	}

	if err := m.Settings.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("settings")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BlockedURLCategories) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlockedURLCategories) UnmarshalBinary(b []byte) error {
	var res BlockedURLCategories
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

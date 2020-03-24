// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateNetworkSwitchSettingsQosRulesOrder updateNetworkSwitchSettingsQosRulesOrder
//
// swagger:model updateNetworkSwitchSettingsQosRulesOrder
type UpdateNetworkSwitchSettingsQosRulesOrder struct {

	// A list of quality of service rule IDs arranged in order in which they should be processed by the switch.
	// Required: true
	RuleIds []string `json:"ruleIds"`
}

// Validate validates this update network switch settings qos rules order
func (m *UpdateNetworkSwitchSettingsQosRulesOrder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRuleIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateNetworkSwitchSettingsQosRulesOrder) validateRuleIds(formats strfmt.Registry) error {

	if err := validate.Required("ruleIds", "body", m.RuleIds); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateNetworkSwitchSettingsQosRulesOrder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateNetworkSwitchSettingsQosRulesOrder) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchSettingsQosRulesOrder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
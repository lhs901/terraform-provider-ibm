// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ListenerPolicyRuleReference listener policy rule reference
// swagger:model ListenerPolicyRuleReference
type ListenerPolicyRuleReference struct {

	// The condition of the rule.
	// Enum: [contains equals matches_regex]
	Condition string `json:"condition,omitempty"`

	// HTTP header field. This is only applicable to "header" rule type.
	Field string `json:"field,omitempty"`

	// The rule's canonical URL.
	// Pattern: ^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$
	Href string `json:"href,omitempty"`

	// The policy's unique identifier.
	// Pattern: ^[-0-9a-z_]+$
	ID string `json:"id,omitempty"`

	// The type of the rule.
	// Enum: [path hostname header]
	Type string `json:"type,omitempty"`

	// Value to be matched for rule condition
	Value string `json:"value,omitempty"`
}

// Validate validates this listener policy rule reference
func (m *ListenerPolicyRuleReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCondition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var listenerPolicyRuleReferenceTypeConditionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["contains","equals","matches_regex"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listenerPolicyRuleReferenceTypeConditionPropEnum = append(listenerPolicyRuleReferenceTypeConditionPropEnum, v)
	}
}

const (

	// ListenerPolicyRuleReferenceConditionContains captures enum value "contains"
	ListenerPolicyRuleReferenceConditionContains string = "contains"

	// ListenerPolicyRuleReferenceConditionEquals captures enum value "equals"
	ListenerPolicyRuleReferenceConditionEquals string = "equals"

	// ListenerPolicyRuleReferenceConditionMatchesRegex captures enum value "matches_regex"
	ListenerPolicyRuleReferenceConditionMatchesRegex string = "matches_regex"
)

// prop value enum
func (m *ListenerPolicyRuleReference) validateConditionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, listenerPolicyRuleReferenceTypeConditionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ListenerPolicyRuleReference) validateCondition(formats strfmt.Registry) error {

	if swag.IsZero(m.Condition) { // not required
		return nil
	}

	// value enum
	if err := m.validateConditionEnum("condition", "body", m.Condition); err != nil {
		return err
	}

	return nil
}

func (m *ListenerPolicyRuleReference) validateHref(formats strfmt.Registry) error {

	if swag.IsZero(m.Href) { // not required
		return nil
	}

	if err := validate.Pattern("href", "body", string(m.Href), `^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$`); err != nil {
		return err
	}

	return nil
}

func (m *ListenerPolicyRuleReference) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.Pattern("id", "body", string(m.ID), `^[-0-9a-z_]+$`); err != nil {
		return err
	}

	return nil
}

var listenerPolicyRuleReferenceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["path","hostname","header"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listenerPolicyRuleReferenceTypeTypePropEnum = append(listenerPolicyRuleReferenceTypeTypePropEnum, v)
	}
}

const (

	// ListenerPolicyRuleReferenceTypePath captures enum value "path"
	ListenerPolicyRuleReferenceTypePath string = "path"

	// ListenerPolicyRuleReferenceTypeHostname captures enum value "hostname"
	ListenerPolicyRuleReferenceTypeHostname string = "hostname"

	// ListenerPolicyRuleReferenceTypeHeader captures enum value "header"
	ListenerPolicyRuleReferenceTypeHeader string = "header"
)

// prop value enum
func (m *ListenerPolicyRuleReference) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, listenerPolicyRuleReferenceTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ListenerPolicyRuleReference) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListenerPolicyRuleReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListenerPolicyRuleReference) UnmarshalBinary(b []byte) error {
	var res ListenerPolicyRuleReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
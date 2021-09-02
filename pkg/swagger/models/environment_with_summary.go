// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EnvironmentWithSummary A managed environment with its latest scan summary.
//
// swagger:model EnvironmentWithSummary
type EnvironmentWithSummary struct {
	Environment

	// resource summary
	ResourceSummary *ResourceSummary `json:"resource_summary,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *EnvironmentWithSummary) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Environment
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Environment = aO0

	// AO1
	var dataAO1 struct {
		ResourceSummary *ResourceSummary `json:"resource_summary,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ResourceSummary = dataAO1.ResourceSummary

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m EnvironmentWithSummary) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Environment)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		ResourceSummary *ResourceSummary `json:"resource_summary,omitempty"`
	}

	dataAO1.ResourceSummary = m.ResourceSummary

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this environment with summary
func (m *EnvironmentWithSummary) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Environment
	if err := m.Environment.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceSummary(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnvironmentWithSummary) validateResourceSummary(formats strfmt.Registry) error {

	if swag.IsZero(m.ResourceSummary) { // not required
		return nil
	}

	if m.ResourceSummary != nil {
		if err := m.ResourceSummary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource_summary")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EnvironmentWithSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnvironmentWithSummary) UnmarshalBinary(b []byte) error {
	var res EnvironmentWithSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
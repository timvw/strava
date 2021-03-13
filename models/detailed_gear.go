// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DetailedGear detailed gear
//
// swagger:model detailedGear
type DetailedGear struct {
	SummaryGear

	// The gear's brand name.
	BrandName string `json:"brand_name,omitempty"`

	// The gear's description.
	Description string `json:"description,omitempty"`

	// The gear's frame type (bike only).
	FrameType int64 `json:"frame_type,omitempty"`

	// The gear's model name.
	ModelName string `json:"model_name,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *DetailedGear) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 SummaryGear
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.SummaryGear = aO0

	// AO1
	var dataAO1 struct {
		BrandName string `json:"brand_name,omitempty"`

		Description string `json:"description,omitempty"`

		FrameType int64 `json:"frame_type,omitempty"`

		ModelName string `json:"model_name,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.BrandName = dataAO1.BrandName

	m.Description = dataAO1.Description

	m.FrameType = dataAO1.FrameType

	m.ModelName = dataAO1.ModelName

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m DetailedGear) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.SummaryGear)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		BrandName string `json:"brand_name,omitempty"`

		Description string `json:"description,omitempty"`

		FrameType int64 `json:"frame_type,omitempty"`

		ModelName string `json:"model_name,omitempty"`
	}

	dataAO1.BrandName = m.BrandName

	dataAO1.Description = m.Description

	dataAO1.FrameType = m.FrameType

	dataAO1.ModelName = m.ModelName

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this detailed gear
func (m *DetailedGear) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SummaryGear
	if err := m.SummaryGear.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this detailed gear based on the context it is used
func (m *DetailedGear) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SummaryGear
	if err := m.SummaryGear.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DetailedGear) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DetailedGear) UnmarshalBinary(b []byte) error {
	var res DetailedGear
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

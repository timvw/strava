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

// MovingStream moving stream
//
// swagger:model movingStream
type MovingStream struct {
	BaseStream

	// The sequence of moving values for this stream, as boolean values
	Data []bool `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *MovingStream) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BaseStream
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BaseStream = aO0

	// AO1
	var dataAO1 struct {
		Data []bool `json:"data"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Data = dataAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m MovingStream) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BaseStream)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Data []bool `json:"data"`
	}

	dataAO1.Data = m.Data

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this moving stream
func (m *MovingStream) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseStream
	if err := m.BaseStream.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this moving stream based on the context it is used
func (m *MovingStream) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseStream
	if err := m.BaseStream.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *MovingStream) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MovingStream) UnmarshalBinary(b []byte) error {
	var res MovingStream
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

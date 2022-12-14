// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ModelsPostInitiatedNotificationResponse models post initiated notification response
//
// swagger:model models.PostInitiatedNotificationResponse
type ModelsPostInitiatedNotificationResponse struct {

	// external Id
	ExternalID string `json:"externalId,omitempty"`
}

// Validate validates this models post initiated notification response
func (m *ModelsPostInitiatedNotificationResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this models post initiated notification response based on context it is used
func (m *ModelsPostInitiatedNotificationResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelsPostInitiatedNotificationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsPostInitiatedNotificationResponse) UnmarshalBinary(b []byte) error {
	var res ModelsPostInitiatedNotificationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

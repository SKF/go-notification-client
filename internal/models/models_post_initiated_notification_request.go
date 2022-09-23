// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ModelsPostInitiatedNotificationRequest models post initiated notification request
//
// swagger:model models.PostInitiatedNotificationRequest
type ModelsPostInitiatedNotificationRequest struct {

	// body
	Body string `json:"body,omitempty"`

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// header
	Header string `json:"header,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// notification type name
	NotificationTypeName string `json:"notificationTypeName,omitempty"`

	// resource Id
	ResourceID string `json:"resourceId,omitempty"`
}

// Validate validates this models post initiated notification request
func (m *ModelsPostInitiatedNotificationRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this models post initiated notification request based on context it is used
func (m *ModelsPostInitiatedNotificationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelsPostInitiatedNotificationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsPostInitiatedNotificationRequest) UnmarshalBinary(b []byte) error {
	var res ModelsPostInitiatedNotificationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
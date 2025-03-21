// Package defs_internal_item_db provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package defs_internal_item_db

import (
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

// PgItem defines model for PgItem.
type PgItem struct {
	Category     string             `json:"category"`
	Comment      *string            `json:"comment,omitempty"`
	CreatedAt    time.Time          `json:"created_at"`
	CustomerUUID openapi_types.UUID `json:"customer_uuid"`
	Description  string             `json:"description"`
	DishUUID     openapi_types.UUID `json:"dish_uuid"`
	IsDraft      bool               `json:"is_draft"`
	Name         string             `json:"name"`
	OrderUUID    openapi_types.UUID `json:"order_uuid"`
	PictureLink  string             `json:"picture_link"`
	Price        float32            `json:"price"`
	Resolution   *string            `json:"resolution,omitempty"`
	Status       string             `json:"status"`
	UpdatedAt    time.Time          `json:"updated_at"`
	UUID         openapi_types.UUID `json:"uuid"`
	Weight       int                `json:"weight"`
}

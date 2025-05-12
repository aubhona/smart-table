package views

import (
	"context"

	viewsAdminRestaurant "github.com/smart-table/src/views/codegen/admin_restaurant"
)

func (h *AdminV1RestaurantHandler) PostAdminV1RestaurantDishEdit(
	ctx context.Context,
	request viewsAdminRestaurant.PostAdminV1RestaurantDishEditRequestObject,
) (viewsAdminRestaurant.PostAdminV1RestaurantDishEditResponseObject, error) {
	return viewsAdminRestaurant.PostAdminV1RestaurantDishEdit204Response{}, nil
}

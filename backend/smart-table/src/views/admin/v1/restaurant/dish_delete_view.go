package views

import (
	"context"

	viewsAdminRestaurant "github.com/smart-table/src/views/codegen/admin_restaurant"
)

func (h *AdminV1RestaurantHandler) PostAdminV1RestaurantDishDelete(
	ctx context.Context,
	request viewsAdminRestaurant.PostAdminV1RestaurantDishDeleteRequestObject,
) (viewsAdminRestaurant.PostAdminV1RestaurantDishDeleteResponseObject, error) {
	return viewsAdminRestaurant.PostAdminV1RestaurantDishDelete204Response{}, nil
}

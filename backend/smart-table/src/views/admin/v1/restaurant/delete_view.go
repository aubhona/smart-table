package views

import (
	"context"

	viewsAdminRestaurant "github.com/smart-table/src/views/codegen/admin_restaurant"
)

func (h *AdminV1RestaurantHandler) PostAdminV1RestaurantDelete(
	ctx context.Context,
	request viewsAdminRestaurant.PostAdminV1RestaurantDeleteRequestObject,
) (viewsAdminRestaurant.PostAdminV1RestaurantDeleteResponseObject, error) {
	return viewsAdminRestaurant.PostAdminV1RestaurantDelete204Response{}, nil
}

package views

import (
	"context"

	viewsAdminRestaurant "github.com/smart-table/src/views/codegen/admin_restaurant"
)

func (h *AdminV1RestaurantHandler) PostAdminV1RestaurantEdit(
	ctx context.Context,
	request viewsAdminRestaurant.PostAdminV1RestaurantEditRequestObject,
) (viewsAdminRestaurant.PostAdminV1RestaurantEditResponseObject, error) {
	return viewsAdminRestaurant.PostAdminV1RestaurantEdit204Response{}, nil
}

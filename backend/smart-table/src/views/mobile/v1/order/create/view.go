package views

import (
	"context"

	app "github.com/es-debug/backend-academy-2024-go-template/src/domains/orders/app/use_cases"
	"github.com/es-debug/backend-academy-2024-go-template/src/domains/orders/di"
	"github.com/es-debug/backend-academy-2024-go-template/src/utils"
	views_order "github.com/es-debug/backend-academy-2024-go-template/src/views/codegen/order"
)

type MobileV1OrderCreateHandler struct{}

func (h *MobileV1OrderCreateHandler) PostMobileV1OrderCreate(ctx context.Context, request views_order.PostMobileV1OrderCreateRequestObject) (views_order.PostMobileV1OrderCreateResponseObject, error) {
	handler, err := di.GetFromContainer[app.OrderCreateCommandHandler](ctx)
	if err != nil {
		return nil, err
	}

	result, err := handler.Handle(&app.OrderCreateCommand{
		TableId:      request.Body.TableID,
		CustomerUuid: request.Body.CustomerUUID,
		RoomCode:     utils.OptionalFromPointer(request.Body.RoomCode),
	})

	if err != nil {
		return nil, err
	}

	return views_order.PostMobileV1OrderCreate200JSONResponse{
		OrderUUID:    result.OrderUUID,
		SkipRoomCode: result.SkipRoomCode.ToPointer(),
	}, nil
}

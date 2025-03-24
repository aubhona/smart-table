package views

import (
	"context"
	// app "github.com/es-debug/backend-academy-2024-go-template/src/domains/admin/app/use_cases"
	// "github.com/es-debug/backend-academy-2024-go-template/src/domains/admin/di"
	// "github.com/es-debug/backend-academy-2024-go-template/src/utils"
	views_order "github.com/es-debug/backend-academy-2024-go-template/src/views/codegen/admin_user"
)

func (h *V1AdminUserHandler) PostV1AdminUserSignUp(ctx context.Context, request views_order.PostV1AdminUserSignUpRequestObject) (views_order.PostV1AdminUserSignUpResponseObject, error) {
	// handler, err := di.GetFromContainer[app.UserSingUpCommandHandler](ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// result, err := handler.Handle(&app.OrderCreateCommand{
	// 	TableId:      request.Body.TableID,
	// 	CustomerUuid: request.Body.CustomerUUID,
	// 	RoomCode:     utils.OptionalFromPointer(request.Body.RoomCode),
	// })
	return nil, nil
}

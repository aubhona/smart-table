package views

import (
	"context"
	"fmt"
	"time"

	app "github.com/smart-table/src/domains/admin/app/use_cases"
	appErrors "github.com/smart-table/src/domains/admin/app/use_cases/errors"
	domainErrors "github.com/smart-table/src/domains/admin/domain/errors"
	"github.com/smart-table/src/logging"
	"github.com/smart-table/src/utils"
	viewsAdminPlace "github.com/smart-table/src/views/codegen/admin_place"
)

func (h *AdminV1PlaceHandler) PostAdminV1PlaceCreate(
	ctx context.Context,
	request viewsAdminPlace.PostAdminV1PlaceCreateRequestObject,
) (viewsAdminPlace.PostAdminV1PlaceCreateResponseObject, error) {
	handler, err := utils.GetFromContainer[*app.PlaceCreateCommandHandler](ctx)
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while getting command handler: %v", err))
		return nil, err
	}

	openingTime, err := time.Parse("15:04", request.Body.OpeningTime)
	if err != nil {
		return viewsAdminPlace.PostAdminV1PlaceCreate400JSONResponse{
			Code:    viewsAdminPlace.IncorrectTimeFormat,
			Message: err.Error(),
		}, nil
	}

	closingTime, err := time.Parse("15:04", request.Body.ClosingTime)
	if err != nil {
		return viewsAdminPlace.PostAdminV1PlaceCreate400JSONResponse{
			Code:    viewsAdminPlace.IncorrectTimeFormat,
			Message: err.Error(),
		}, nil
	}

	result, err := handler.Handle(&app.PlaceCreateCommand{
		OwnerUUID:      request.Params.UserUUID,
		RestaurantUUID: request.Body.RestaurantUUID,
		Address:        request.Body.Address,
		TableCount:     request.Body.TableCount,
		OpeningTime:    openingTime,
		ClosingTime:    closingTime,
	})
	if err != nil {
		switch {
		case utils.IsTheSameErrorType[domainErrors.InvalidTableCount](err):
			return viewsAdminPlace.PostAdminV1PlaceCreate400JSONResponse{
				Code:    viewsAdminPlace.IncorrectTableCount,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[domainErrors.RestaurantNotFound](err):
			return viewsAdminPlace.PostAdminV1PlaceCreate404JSONResponse{
				Code:    viewsAdminPlace.RestaurantNotFound,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[appErrors.RestaurantAccessDenied](err):
			return viewsAdminPlace.PostAdminV1PlaceCreate403JSONResponse{
				Code:    viewsAdminPlace.AccessDenied,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[appErrors.PlaceAddressAlreadyExists](err):
			return viewsAdminPlace.PostAdminV1PlaceCreate403JSONResponse{
				Code:    viewsAdminPlace.AlreadyExist,
				Message: err.Error(),
			}, nil
		}

		logging.GetLogger().Error(fmt.Sprintf("Error while getting result from command handler: %v", err))

		return nil, err
	}

	return viewsAdminPlace.PostAdminV1PlaceCreate200JSONResponse{
		PlaceUUID: result.PlaceUUID,
	}, nil
}

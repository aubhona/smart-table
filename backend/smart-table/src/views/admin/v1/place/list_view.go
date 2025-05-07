package views

import (
	"context"
	"fmt"

	app "github.com/smart-table/src/domains/admin/app/use_cases"
	appErrors "github.com/smart-table/src/domains/admin/app/use_cases/errors"
	"github.com/smart-table/src/domains/admin/domain"
	domainErrors "github.com/smart-table/src/domains/admin/domain/errors"
	"github.com/smart-table/src/logging"
	"github.com/smart-table/src/utils"
	viewsAdminPlace "github.com/smart-table/src/views/codegen/admin_place"
)

func convertPlaceToPlaceInfo(
	place utils.SharedRef[domain.Place],
) viewsAdminPlace.PlaceInfo {
	return viewsAdminPlace.PlaceInfo{
		UUID:        place.Get().GetUUID(),
		Address:     place.Get().GetAddress(),
		TableCount:  place.Get().GetTableCount(),
		OpeningTime: place.Get().GetOpeningTime().Format("15:04"),
		ClosingTime: place.Get().GetClosingTime().Format("15:04"),
	}
}

func (h *AdminV1PlaceHandler) PostAdminV1PlaceList(
	ctx context.Context,
	request viewsAdminPlace.PostAdminV1PlaceListRequestObject,
) (viewsAdminPlace.PostAdminV1PlaceListResponseObject, error) {
	handler, err := utils.GetFromContainer[*app.PlaceListCommandHandler](ctx)
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while getting command handler: %v", err))
		return nil, err
	}

	result, err := handler.Handle(&app.PlaceListCommand{
		OwnerUUID:      request.Params.UserUUID,
		RestaurantUUID: request.Body.RestaurantUUID,
	})
	if err != nil {
		switch {
		case utils.IsTheSameErrorType[domainErrors.RestaurantNotFound](err):
			return viewsAdminPlace.PostAdminV1PlaceList404JSONResponse{
				Code:    viewsAdminPlace.RestaurantNotFound,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[appErrors.RestaurantAccessDenied](err):
			return viewsAdminPlace.PostAdminV1PlaceList403JSONResponse{
				Code:    viewsAdminPlace.AccessDenied,
				Message: err.Error(),
			}, nil
		}

		logging.GetLogger().Error(fmt.Sprintf("Error while getting result from command handler: %v", err))

		return nil, err
	}

	placeInfoList := make([]viewsAdminPlace.PlaceInfo, 0, len(result.PlaceList))

	for _, place := range result.PlaceList {
		placeInfoList = append(placeInfoList, convertPlaceToPlaceInfo(place))
	}

	return viewsAdminPlace.PostAdminV1PlaceList200JSONResponse{
		PlaceList: placeInfoList,
	}, nil
}

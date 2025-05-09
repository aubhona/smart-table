package views

import (
	"context"
	"fmt"

	app "github.com/smart-table/src/domains/admin/app/use_cases"
	appErrors "github.com/smart-table/src/domains/admin/app/use_cases/errors"
	domainErrors "github.com/smart-table/src/domains/admin/domain/errors"
	"github.com/smart-table/src/logging"
	"github.com/smart-table/src/utils"
	viewsAdminPlace "github.com/smart-table/src/views/codegen/admin_place"
)

func (h *AdminV1PlaceHandler) PostAdminV1PlaceTableDeeplinksList(
	ctx context.Context,
	request viewsAdminPlace.PostAdminV1PlaceTableDeeplinksListRequestObject,
) (viewsAdminPlace.PostAdminV1PlaceTableDeeplinksListResponseObject, error) {
	handler, err := utils.GetFromContainer[*app.TableDeepLinksListCommandHandler](ctx)
	if err != nil {
		logging.GetLogger().Error(fmt.Sprintf("Error while getting command handler: %v", err))
		return nil, err
	}

	result, err := handler.Handle(&app.TableDeepLinksListCommand{
		UserUUID:  request.Params.UserUUID,
		PlaceUUID: request.Body.PlaceUUID,
	})
	if err != nil {
		switch {
		case utils.IsTheSameErrorType[appErrors.PlaceAccessDenied](err):
			return viewsAdminPlace.PostAdminV1PlaceTableDeeplinksList403JSONResponse{
				Code:    viewsAdminPlace.AccessDenied,
				Message: err.Error(),
			}, nil
		case utils.IsTheSameErrorType[domainErrors.PlaceNotFound](err):
			return viewsAdminPlace.PostAdminV1PlaceTableDeeplinksList404JSONResponse{
				Code:    viewsAdminPlace.PlaceNotFound,
				Message: err.Error(),
			}, nil
		}
	}

	return &viewsAdminPlace.PostAdminV1PlaceTableDeeplinksList200JSONResponse{
		Deeplinks: result.TableDeepLinksList,
	}, nil
}

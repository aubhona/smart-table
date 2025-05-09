package app

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/smart-table/src/dependencies"
	appErrors "github.com/smart-table/src/domains/admin/app/services/errors"
	"github.com/smart-table/src/domains/admin/domain"
	"github.com/smart-table/src/utils"
)

type PlaceTableService struct {
	webAppURL string
}

func NewPlaceTableService(dependencies *dependencies.Dependencies) *PlaceTableService {
	return &PlaceTableService{webAppURL: dependencies.Config.Bot.WebAppURL}
}

func (p *PlaceTableService) GetPlaceUUIDFromTableID(tableID string) (uuid.UUID, error) {
	parts := strings.Split(tableID, "_")
	if len(parts) != 2 {
		return uuid.Nil, appErrors.InvalidTableID{TableID: tableID}
	}

	return uuid.Parse(parts[0])
}

func (p *PlaceTableService) GetTableNumberFromTableID(tableID string) (int, error) {
	parts := strings.Split(tableID, "_")
	if len(parts) != 2 {
		return 0, appErrors.InvalidTableID{TableID: tableID}
	}

	return strconv.Atoi(parts[1])
}

func (p *PlaceTableService) GetTableDeepLinkForQR(place utils.SharedRef[domain.Place]) []string {
	tableIDs := place.Get().GetTableIDs()

	return lo.Map(tableIDs, func(tableID string, _ int) string {
		return fmt.Sprintf("%s=%s", p.webAppURL, tableID)
	})
}

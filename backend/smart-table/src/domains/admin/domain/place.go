package domain

import (
	"time"

	"github.com/google/uuid"
	domainErrors "github.com/smart-table/src/domains/admin/domain/errors"
	domainServices "github.com/smart-table/src/domains/admin/domain/services"
	"github.com/smart-table/src/utils"
)

type Place struct {
	uuid        uuid.UUID
	restaurant  utils.SharedRef[Restaurant]
	address     string
	tableCount  int
	openingTime time.Time
	closingTime time.Time
	createdAt   time.Time
	updatedAt   time.Time
}

func NewPlace(
	restaurant utils.SharedRef[Restaurant],
	address string,
	tableCount int,
	openingTime,
	closingTime time.Time,
	uuidGenerator *domainServices.UUIDGenerator,
) (utils.SharedRef[Place], error) {
	if tableCount <= 0 {
		return utils.SharedRef[Place]{}, domainErrors.InvalidTableCount{TableCount: tableCount}
	}

	place := Place{
		restaurant:  restaurant,
		address:     address,
		tableCount:  tableCount,
		openingTime: openingTime,
		closingTime: closingTime,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
	}

	shardID := uuidGenerator.GetShardID()
	place.uuid = uuidGenerator.GenerateShardedUUID(shardID)

	placeRef, _ := utils.NewSharedRef(&place)

	return placeRef, nil
}

func RestorePlace(
	id uuid.UUID,
	restaurant utils.SharedRef[Restaurant],
	address string,
	tableCount int,
	openingTime,
	closingTime,
	createdAt,
	updatedAt time.Time,
) utils.SharedRef[Place] {
	place := Place{
		uuid:        id,
		restaurant:  restaurant,
		address:     address,
		tableCount:  tableCount,
		openingTime: openingTime,
		closingTime: closingTime,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}

	placeRef, _ := utils.NewSharedRef(&place)

	return placeRef
}

func (p *Place) GetUUID() uuid.UUID                         { return p.uuid }
func (p *Place) GetRestaurant() utils.SharedRef[Restaurant] { return p.restaurant }
func (p *Place) GetAddress() string                         { return p.address }
func (p *Place) GetTableCount() int                         { return p.tableCount }
func (p *Place) GetOpeningTime() time.Time                  { return p.openingTime }
func (p *Place) GetClosingTime() time.Time                  { return p.closingTime }
func (p *Place) GetCreatedAt() time.Time                    { return p.createdAt }
func (p *Place) GetUpdatedAt() time.Time                    { return p.updatedAt }

package domain

import (
	"time"

	"github.com/google/uuid"
	domainErrors "github.com/smart-table/src/domains/admin/domain/errors"
	domainServices "github.com/smart-table/src/domains/admin/domain/services"
	"github.com/smart-table/src/utils"
)

type Place struct {
	uuid           uuid.UUID
	restaurantUUID uuid.UUID
	address        string
	tableCount     int
	openingTime    time.Time
	closingTime    time.Time
	createdAt      time.Time
	updatedAt      time.Time
}

func NewPlace(
	restaurantUUID uuid.UUID,
	address string,
	tableCount int,
	openingTime time.Time,
	closingTime time.Time,
	uuidGenerator *domainServices.UUIDGenerator,
) (utils.SharedRef[Place], error) {
	if tableCount <= 0 {
		return utils.SharedRef[Place]{}, domainErrors.InvalidTableCount{TableCount: tableCount}
	}

	shardID := uuidGenerator.GetShardID()

	return RestorePlace(
		uuidGenerator.GenerateShardedUUID(shardID),
		restaurantUUID,
		address,
		tableCount,
		openingTime,
		closingTime,
		time.Now(),
		time.Now(),
	), nil
}

func RestorePlace(
	id uuid.UUID,
	restaurantUUID uuid.UUID,
	address string,
	tableCount int,
	openingTime time.Time,
	closingTime time.Time,
	createdAt time.Time,
	updatedAt time.Time,
) utils.SharedRef[Place] {
	place := Place{
		uuid:           id,
		restaurantUUID: restaurantUUID,
		address:        address,
		tableCount:     tableCount,
		openingTime:    openingTime,
		closingTime:    closingTime,
		createdAt:      createdAt,
		updatedAt:      updatedAt,
	}

	placeRef, _ := utils.NewSharedRef(&place)

	return placeRef
}

func (p *Place) GetUUID() uuid.UUID          { return p.uuid }
func (p *Place) GetRestauranUUID() uuid.UUID { return p.restaurantUUID }
func (p *Place) GetAddress() string          { return p.address }
func (p *Place) GetTableCount() int          { return p.tableCount }
func (p *Place) GetOpeningTime() time.Time   { return p.openingTime }
func (p *Place) GetClosingTime() time.Time   { return p.closingTime }
func (p *Place) GetCreatedAt() time.Time     { return p.createdAt }
func (p *Place) GetUpdatedAt() time.Time     { return p.updatedAt }

package pg

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/samber/lo"
	"github.com/smart-table/src/domains/admin/domain"
	domainErrors "github.com/smart-table/src/domains/admin/domain/errors"
	db "github.com/smart-table/src/domains/admin/infra/pg/codegen"
	"github.com/smart-table/src/domains/admin/infra/pg/mapper"
	"github.com/smart-table/src/utils"
)

type PlaceRepository struct {
	coonPool *pgxpool.Pool
}

func NewPlaceRepository(pool *pgxpool.Pool) *PlaceRepository {
	return &PlaceRepository{pool}
}

func (p *PlaceRepository) Begin() (domain.Transaction, error) {
	ctx := context.Background()
	tx, err := p.coonPool.Begin(ctx)

	if err != nil {
		return nil, err
	}

	return &pgTx{tx: tx, ctx: ctx}, nil
}

func (p *PlaceRepository) Commit(tx domain.Transaction) error {
	return tx.Commit()
}

func (p *PlaceRepository) Rollback(tx domain.Transaction) error {
	return tx.Rollback()
}

func (p *PlaceRepository) Save(tx domain.Transaction, place utils.SharedRef[domain.Place]) error {
	ctx := context.Background()
	trx := tx.(*pgTx)
	queries := db.New(p.coonPool).WithTx(trx.tx)

	pgPlace, err := mapper.ConvertToPgPlace(place)
	if err != nil {
		return err
	}

	_, err = queries.InsertPlace(ctx, pgPlace)

	return err
}

func (p *PlaceRepository) Update(tx domain.Transaction, place utils.SharedRef[domain.Place]) error {
	ctx := context.Background()
	trx := tx.(*pgTx)
	queries := db.New(p.coonPool).WithTx(trx.tx)

	pgEmployees, err := mapper.ConvertToPgEmployees(place.Get().GetEmployees())
	if err != nil {
		return err
	}

	pgMenuDishes, err := mapper.ConvertToPgMenuDishes(place.Get().GetMenuDishes())
	if err != nil {
		return err
	}

	err = queries.UpsertEmployees(ctx, pgEmployees)
	if err != nil {
		return err
	}

	err = queries.UpsertMenuDishes(ctx, pgMenuDishes)
	if err != nil {
		return err
	}

	err = queries.DeleteMenuDishesByUUID(ctx, place.Get().GetDeletedMenuDishUUIDs())
	if err != nil {
		return err
	}

	return nil
}

func getPlaceNotFoundError(placeUUIDs []uuid.UUID, places []utils.SharedRef[domain.Place]) error {
	placeUUIDSet := lo.SliceToMap(places, func(place utils.SharedRef[domain.Place]) (uuid.UUID, interface{}) {
		return place.Get().GetUUID(), nil
	})

	for _, placeUUID := range placeUUIDs {
		if _, found := placeUUIDSet[placeUUID]; !found {
			return domainErrors.PlaceNotFound{UUID: placeUUID}
		}
	}

	return nil
}

func (p *PlaceRepository) FindPlace(id uuid.UUID) (utils.SharedRef[domain.Place], error) {
	places, err := p.FindPlaces([]uuid.UUID{id})
	if err != nil {
		return utils.SharedRef[domain.Place]{}, err
	}

	return places[0], nil
}

func (p *PlaceRepository) FindPlaces(uuids []uuid.UUID) ([]utils.SharedRef[domain.Place], error) {
	ctx := context.Background()
	queries := db.New(p.coonPool)

	pgResults, err := queries.FetchPlacesByUUID(ctx, uuids)
	if err != nil {
		return nil, err
	}

	places, err := mapper.ConvertPgPlacesToModel(pgResults)
	if err != nil {
		return nil, err
	}

	if len(places) == len(uuids) {
		return places, nil
	}

	return nil, getPlaceNotFoundError(uuids, places)
}

func (p *PlaceRepository) FindPlaceByAddress(address string, restaurantUUID uuid.UUID) (utils.SharedRef[domain.Place], error) {
	ctx := context.Background()
	queries := db.New(p.coonPool)

	placeUUID, err := queries.GetPlaceUUIDByAddress(ctx, db.GetPlaceUUIDByAddressParams{
		Column1: address,
		Column2: restaurantUUID,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return utils.SharedRef[domain.Place]{}, domainErrors.PlaceNotFoundByAddress{Address: address, RestaurantUUID: restaurantUUID}
		}

		return utils.SharedRef[domain.Place]{}, err
	}

	return p.FindPlace(placeUUID)
}

func (p *PlaceRepository) FindPlacesByRestaurantUUID(uuid uuid.UUID) ([]utils.SharedRef[domain.Place], error) {
	ctx := context.Background()
	queries := db.New(p.coonPool)

	placeUUIDs, err := queries.GetPlaceUUIDsByRestaurantUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return p.FindPlaces(placeUUIDs)
}

func (p *PlaceRepository) FindPlacesForUpdate(tx domain.Transaction, uuids []uuid.UUID) ([]utils.SharedRef[domain.Place], error) {
	ctx := context.Background()
	trx := tx.(*pgTx)
	queries := db.New(p.coonPool).WithTx(trx.tx)
	pgResults, err := queries.FetchPlacesByUUID(ctx, uuids)

	if err != nil {
		return nil, err
	}

	places, err := mapper.ConvertPgPlacesToModel(pgResults)
	if err != nil {
		return nil, err
	}

	if len(places) == len(uuids) {
		return places, nil
	}

	return nil, getPlaceNotFoundError(uuids, places)
}

func (p *PlaceRepository) FindPlaceForUpdate(tx domain.Transaction, id uuid.UUID) (utils.SharedRef[domain.Place], error) {
	places, err := p.FindPlacesForUpdate(tx, []uuid.UUID{id})
	if err != nil {
		return utils.SharedRef[domain.Place]{}, err
	}

	return places[0], nil
}

func (p *PlaceRepository) FindPlacesByEmployeeUserUUID(userUUID uuid.UUID) ([]utils.SharedRef[domain.Place], error) {
	ctx := context.Background()
	queries := db.New(p.coonPool)

	placeUUIDs, err := queries.GetPlaceUUIDsByEmployeeUserUUID(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	return p.FindPlaces(placeUUIDs)
}

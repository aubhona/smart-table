package pg

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/smart-table/src/domains/admin/domain"
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

func (p *PlaceRepository) Begin(ctx context.Context) (pgx.Tx, error) {
	return p.coonPool.Begin(ctx)
}

func (p *PlaceRepository) Commit(ctx context.Context, tx pgx.Tx) error {
	return tx.Commit(ctx)
}

func (p *PlaceRepository) Save(ctx context.Context, tx pgx.Tx, place utils.SharedRef[domain.Place]) error {
	queries := db.New(p.coonPool).WithTx(tx)

	pgPlace, err := mapper.ConvertToPgPlace(place)
	if err != nil {
		return err
	}

	_, err = queries.InsertPlace(ctx, pgPlace)

	return err
}

func (p *PlaceRepository) CheckAddressExist(ctx context.Context, address string, restaurantUUID uuid.UUID) (bool, error) {
	queries := db.New(p.coonPool)

	params := db.CheckPlaceAddressExistParams{
		Column1: address,
		Column2: restaurantUUID,
	}

	placeExists, err := queries.CheckPlaceAddressExist(ctx, params)
	if err != nil {
		return false, err
	}

	return placeExists, nil
}

func (p *PlaceRepository) FindPlaceListByRestaurantUUID(
	ctx context.Context,
	restaurantUUID uuid.UUID,
) ([]utils.SharedRef[domain.Place], error) {
	queries := db.New(p.coonPool)

	pgResult, err := queries.FetchPlaceListByRestaurantUUID(ctx, restaurantUUID)
	if err != nil || pgResult == nil {
		return []utils.SharedRef[domain.Place]{}, err
	}

	placeList := make([]utils.SharedRef[domain.Place], 0, len(pgResult))

	for i := range pgResult {
		place, err := mapper.ConvertPgPlaceToModel(pgResult[i])
		if err != nil {
			return []utils.SharedRef[domain.Place]{}, err
		}

		placeList = append(placeList, place)
	}

	return placeList, nil
}

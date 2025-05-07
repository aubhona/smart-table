package pg

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/smart-table/src/domains/admin/domain"
	domainErrors "github.com/smart-table/src/domains/admin/domain/errors"
	db "github.com/smart-table/src/domains/admin/infra/pg/codegen"
	"github.com/smart-table/src/domains/admin/infra/pg/mapper"
	"github.com/smart-table/src/utils"
	"github.com/thoas/go-funk"
)

type RestaurantRepository struct {
	coonPool *pgxpool.Pool
}

func NewRestaurantRepository(pool *pgxpool.Pool) *RestaurantRepository {
	return &RestaurantRepository{pool}
}

func (r *RestaurantRepository) Begin() (domain.Transaction, error) {
	ctx := context.Background()
	tx, err := r.coonPool.Begin(ctx)

	if err != nil {
		return nil, err
	}

	return &pgTx{tx: tx, ctx: ctx}, nil
}

func (r *RestaurantRepository) Commit(tx domain.Transaction) error {
	return tx.Commit()
}

func (r *RestaurantRepository) Rollback(tx domain.Transaction) error {
	return tx.Rollback()
}

func (r *RestaurantRepository) Save(tx domain.Transaction, restaurant utils.SharedRef[domain.Restaurant]) error {
	ctx := context.Background()
	trx := tx.(*pgTx)
	queries := db.New(r.coonPool).WithTx(trx.tx)

	pgRestaurant, err := mapper.ConvertToPgRestaurant(restaurant)
	if err != nil {
		return err
	}

	_, err = queries.InsertRestaurant(ctx, pgRestaurant)

	return err
}

func (r *RestaurantRepository) FindRestaurantByName(name string) (utils.SharedRef[domain.Restaurant], error) {
	ctx := context.Background()
	queries := db.New(r.coonPool)

	restaurantUUID, err := queries.GetRestaurantUUIDByName(ctx, name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return utils.SharedRef[domain.Restaurant]{}, domainErrors.RestaurantNotFoundByName{Name: name}
		}

		return utils.SharedRef[domain.Restaurant]{}, err
	}

	return r.FindRestaurant(restaurantUUID)
}

func (r *RestaurantRepository) FindRestaurants(uuids []uuid.UUID) ([]utils.SharedRef[domain.Restaurant], error) {
	ctx := context.Background()
	queries := db.New(r.coonPool)

	pgResult, err := queries.FetchRestaurantsByUUID(ctx, uuids)
	if err != nil {
		return nil, err
	}

	restaurants, err := mapper.ConvertPgRestaurantsToModel(pgResult)
	if err != nil {
		return nil, err
	}

	if len(restaurants) == len(uuids) {
		return restaurants, nil
	}

	return nil, getRestaurantNotFoundError(uuids, restaurants)
}

func (r *RestaurantRepository) FindRestaurant(id uuid.UUID) (utils.SharedRef[domain.Restaurant], error) {
	restaurants, err := r.FindRestaurants([]uuid.UUID{id})
	if err != nil {
		return utils.SharedRef[domain.Restaurant]{}, err
	}

	return restaurants[0], nil
}

func (r *RestaurantRepository) Update(tx domain.Transaction, restaurant utils.SharedRef[domain.Restaurant]) error {
	ctx := context.Background()
	trx := tx.(*pgTx)
	queries := db.New(r.coonPool).WithTx(trx.tx)

	pgDishes, err := mapper.ConvertToPgDishes(restaurant.Get().GetDishes())
	if err != nil {
		return err
	}

	_, err = queries.UpsertDishes(ctx, pgDishes)
	if err != nil {
		return err
	}

	return nil
}

func (r *RestaurantRepository) FindRestaurantForUpdate(tx domain.Transaction, id uuid.UUID) (utils.SharedRef[domain.Restaurant], error) {
	restaurants, err := r.FindRestaurantsForUpdate(tx, []uuid.UUID{id})
	if err != nil {
		return utils.SharedRef[domain.Restaurant]{}, err
	}

	return restaurants[0], nil
}

func (r *RestaurantRepository) FindRestaurantsForUpdate(
	tx domain.Transaction,
	uuids []uuid.UUID,
) ([]utils.SharedRef[domain.Restaurant], error) {
	ctx := context.Background()
	trx := tx.(*pgTx)
	queries := db.New(r.coonPool).WithTx(trx.tx)

	pgResult, err := queries.FetchRestaurantsForUpdateByUUID(ctx, uuids)
	if err != nil {
		return nil, err
	}

	restaurants, err := mapper.ConvertPgRestaurantsToModel(pgResult)
	if err != nil {
		return nil, err
	}

	if len(restaurants) == len(uuids) {
		return restaurants, nil
	}

	return nil, getRestaurantNotFoundError(uuids, restaurants)
}

func (r *RestaurantRepository) FindRestaurantsByOwnerUUID(
	ownerUUID uuid.UUID,
) ([]utils.SharedRef[domain.Restaurant], error) {
	ctx := context.Background()
	queries := db.New(r.coonPool)

	pgResult, err := queries.GetRestaurantUUIDsByOwnerUUID(ctx, ownerUUID)
	if err != nil || pgResult == nil {
		return []utils.SharedRef[domain.Restaurant]{}, err
	}

	return r.FindRestaurants(pgResult)
}

func getRestaurantNotFoundError(restaurantUUIDs []uuid.UUID, restaurants []utils.SharedRef[domain.Restaurant]) error {
	restaurantUUIDSet := funk.Map(restaurants, func(restaurant utils.SharedRef[domain.Place]) (uuid.UUID, interface{}) {
		return restaurant.Get().GetUUID(), nil
	}).(map[uuid.UUID]interface{})

	for _, restaurantUUID := range restaurantUUIDs {
		if _, found := restaurantUUIDSet[restaurantUUID]; !found {
			return domainErrors.PlaceNotFound{UUID: restaurantUUID}
		}
	}

	return nil
}

package pg

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/smart-table/src/domains/admin/domain"
	db "github.com/smart-table/src/domains/admin/infra/pg/codegen"
	"github.com/smart-table/src/domains/admin/infra/pg/mapper"
	"github.com/smart-table/src/utils"
)

type RestaurantRepository struct {
	coonPool *pgxpool.Pool
}

func NewRestaurantRepository(pool *pgxpool.Pool) *RestaurantRepository {
	return &RestaurantRepository{pool}
}

func (r *RestaurantRepository) Begin(ctx context.Context) (pgx.Tx, error) {
	return r.coonPool.Begin(ctx)
}

func (r *RestaurantRepository) Commit(ctx context.Context, tx pgx.Tx) error {
	return tx.Commit(ctx)
}

func (r *RestaurantRepository) Save(ctx context.Context, tx pgx.Tx, restaurant utils.SharedRef[domain.Restaurant]) error {
	queries := db.New(r.coonPool).WithTx(tx)

	pgRestaurant, err := mapper.ConvertToPgRestaurant(restaurant)
	if err != nil {
		return err
	}

	_, err = queries.InsertRestaurant(ctx, pgRestaurant)

	return err
}

func (r *RestaurantRepository) CheckNameExist(ctx context.Context, name string) (bool, error) {
	queries := db.New(r.coonPool)

	restaurantExists, err := queries.CheckNameExist(ctx, name)
	if err != nil {
		return false, err
	}

	return restaurantExists, nil
}

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

func (o *RestaurantRepository) Begin(ctx context.Context) (pgx.Tx, error) {
	return o.coonPool.Begin(ctx)
}

func (o *RestaurantRepository) Commit(ctx context.Context, tx pgx.Tx) error {
	return tx.Commit(ctx)
}

func (o *RestaurantRepository) Save(ctx context.Context, tx pgx.Tx, Restaurant utils.SharedRef[domain.Restaurant]) error {
	queries := db.New(o.coonPool).WithTx(tx)

	pgRestaurant, err := mapper.ConvertToPgRestaurant(Restaurant)
	if err != nil {
		return err
	}

	_, err = queries.InsertRestaurant(ctx, pgRestaurant)

	return err
}

func (o *RestaurantRepository) CheckNameExist(ctx context.Context, name string) (bool, error) {
	queries := db.New(o.coonPool)

	RestaurantExists, err := queries.CheckNameExist(ctx, name)
	if err != nil {
		return false, err
	}

	return RestaurantExists, nil
}

package pg

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/smart-table/src/domains/customer/domain"
	domainErrors "github.com/smart-table/src/domains/customer/domain/errors"
	db "github.com/smart-table/src/domains/customer/infra/pg/codegen"
	"github.com/smart-table/src/domains/customer/infra/pg/mapper"
	"github.com/smart-table/src/utils"
	"github.com/thoas/go-funk"
)

type OrderRepository struct {
	coonPool *pgxpool.Pool
	trx      *pgx.Tx
}

func NewOrderRepository(pool *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{pool, nil}
}

func (o *OrderRepository) Save(ctx context.Context, order utils.SharedRef[domain.Order]) error {
	queries := db.New(o.coonPool)

	pgOrder, err := mapper.ConvertToPgOrder(order)
	if err != nil {
		return err
	}

	_, err = queries.InsertOrder(ctx, pgOrder)
	if err != nil {
		return err
	}

	pgItems, err := mapper.ConvertToPgItems(order.Get().GetItems())
	if err != nil {
		return err
	}

	_, err = queries.UpsertItems(ctx, pgItems)

	return err
}

func (o *OrderRepository) Begin(ctx context.Context) error {
	if o.trx != nil {
		return fmt.Errorf("transaction already started")
	}

	trx, err := o.coonPool.Begin(ctx)
	if err != nil {
		return err
	}

	*o.trx = trx

	return nil
}

func (o *OrderRepository) Commit(ctx context.Context) error {
	if o.trx == nil {
		return fmt.Errorf("transaction doesn't exist")
	}

	err := (*o.trx).Commit(ctx)
	if err != nil {
		return err
	}

	*o.trx = nil

	return nil
}

func (o *OrderRepository) FindOrders(ctx context.Context, orderUuids []uuid.UUID) ([]utils.SharedRef[domain.Order], error) {
	queries := db.New(o.coonPool)

	pgResult, err := queries.FetchOrders(ctx, orderUuids)
	if err != nil {
		return nil, err
	}

	orders, err := mapper.ConvertPgOrderAggregatesToModels(pgResult)
	if err != nil {
		return nil, err
	}

	if len(orderUuids) == len(orders) {
		return orders, nil
	}

	return nil, getNotFoundError(orderUuids, orders)
}

func (o *OrderRepository) FindOrder(ctx context.Context, orderUUID uuid.UUID) (utils.SharedRef[domain.Order], error) {
	orders, err := o.FindOrders(ctx, []uuid.UUID{orderUUID})

	return orders[0], err
}

func (o *OrderRepository) FindActiveOrderByTableID(ctx context.Context, tableID string) (utils.SharedRef[domain.Order], error) {
	queries := db.New(o.coonPool)

	pgResult, err := queries.GetActiveOrderUuidByTableId(ctx, tableID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return utils.SharedRef[domain.Order]{}, domainErrors.OrderNotFoundByTableID{TableID: tableID}
		}

		return utils.SharedRef[domain.Order]{}, err
	}

	return o.FindOrder(ctx, pgResult)
}

func getNotFoundError(orderUuids []uuid.UUID, orders []utils.SharedRef[domain.Order]) error {
	orderUUIDSet := funk.Map(orders, func(order utils.SharedRef[domain.Order]) (uuid.UUID, interface{}) {
		return order.Get().GetUUID(), nil
	}).(map[uuid.UUID]interface{})

	for _, orderUUID := range orderUuids {
		if _, found := orderUUIDSet[orderUUID]; !found {
			return domainErrors.OrderNotFound{UUID: orderUUID}
		}
	}

	return nil
}

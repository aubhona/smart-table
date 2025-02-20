package pg

import (
	"context"
	"errors"
	"fmt"
	"github.com/es-debug/backend-academy-2024-go-template/src/domains/orders/domain"
	domain_errors "github.com/es-debug/backend-academy-2024-go-template/src/domains/orders/domain/errors"
	db "github.com/es-debug/backend-academy-2024-go-template/src/domains/orders/infra/pg/codegen"
	"github.com/es-debug/backend-academy-2024-go-template/src/domains/orders/infra/pg/mapper"
	"github.com/es-debug/backend-academy-2024-go-template/src/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
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

func (o *OrderRepository) FindOrder(ctx context.Context, orderUuid uuid.UUID) (utils.SharedRef[domain.Order], error) {
	orders, err := o.FindOrders(ctx, []uuid.UUID{orderUuid})

	return orders[0], err
}

func (o *OrderRepository) FindActiveOrderByTableId(ctx context.Context, tableId string) (utils.SharedRef[domain.Order], error) {
	queries := db.New(o.coonPool)

	pgResult, err := queries.GetActiveOrderUuidByTableId(ctx, tableId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return utils.SharedRef[domain.Order]{}, domain_errors.OrderNotFoundByTableId{TableId: tableId}
		}

		return utils.SharedRef[domain.Order]{}, err
	}

	return o.FindOrder(ctx, pgResult)
}

func getNotFoundError(orderUuids []uuid.UUID, orders []utils.SharedRef[domain.Order]) error {
	orderUuidSet := funk.Map(orders, func(order utils.SharedRef[domain.Order]) (uuid.UUID, interface{}) {
		return order.Get().GetUUID(), nil
	}).(map[uuid.UUID]interface{})

	for _, orderUuid := range orderUuids {
		if _, found := orderUuidSet[orderUuid]; !found {
			return domain_errors.OrderNotFound{Uuid: orderUuid}
		}
	}

	return nil
}

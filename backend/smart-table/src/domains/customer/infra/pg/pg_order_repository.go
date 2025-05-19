package pg

import (
	"context"
	"errors"

	"github.com/samber/lo"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/smart-table/src/domains/customer/domain"
	domainErrors "github.com/smart-table/src/domains/customer/domain/errors"
	db "github.com/smart-table/src/domains/customer/infra/pg/codegen"
	"github.com/smart-table/src/domains/customer/infra/pg/mapper"
	"github.com/smart-table/src/utils"
)

type OrderRepository struct {
	coonPool *pgxpool.Pool
}

func NewOrderRepository(pool *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{pool}
}

func (o *OrderRepository) Begin() (domain.Transaction, error) {
	ctx := context.Background()
	tx, err := o.coonPool.Begin(ctx)

	if err != nil {
		return nil, err
	}

	return &pgTx{tx: tx, ctx: ctx}, nil
}

func (o *OrderRepository) Commit(tx domain.Transaction) error {
	return tx.Commit()
}

func (o *OrderRepository) Rollback(tx domain.Transaction) error {
	return tx.Rollback()
}

func (o *OrderRepository) Save(tx domain.Transaction, order utils.SharedRef[domain.Order]) error {
	ctx := context.Background()
	trx := tx.(*pgTx)
	queries := db.New(o.coonPool).WithTx(trx.tx)

	pgOrder, err := mapper.ConvertToPgOrder(order)
	if err != nil {
		return err
	}

	err = queries.InsertOrder(ctx, pgOrder)
	if err != nil {
		return err
	}

	pgItems, err := mapper.ConvertToPgItems(order.Get().GetItems())
	if err != nil {
		return err
	}

	err = queries.UpsertItems(ctx, pgItems)
	if err != nil {
		return err
	}

	return nil
}

func (o *OrderRepository) Update(tx domain.Transaction, order utils.SharedRef[domain.Order]) error {
	ctx := context.Background()
	trx := tx.(*pgTx)
	queries := db.New(o.coonPool).WithTx(trx.tx)

	pgOrder, err := mapper.ConvertToPgOrder(order)
	if err != nil {
		return err
	}

	pgItemsToInsert, err := mapper.ConvertToPgItems(order.Get().GetItems())
	if err != nil {
		return err
	}

	pgItemsToDelete, err := mapper.ConvertToPgItems(order.Get().GetDeletesItems())
	if err != nil {
		return err
	}

	err = queries.UpdateOrder(ctx, pgOrder)
	if err != nil {
		return err
	}

	err = queries.UpsertItems(ctx, pgItemsToInsert)
	if err != nil {
		return err
	}

	err = queries.DeleteItems(ctx, pgItemsToDelete)
	if err != nil {
		return err
	}

	return nil
}

func getNotFoundError(orderUUIDs []uuid.UUID, orders []utils.SharedRef[domain.Order]) error {
	orderUUIDSet := lo.SliceToMap(orders, func(order utils.SharedRef[domain.Order]) (uuid.UUID, interface{}) {
		return order.Get().GetUUID(), nil
	})

	for _, orderUUID := range orderUUIDs {
		if _, found := orderUUIDSet[orderUUID]; !found {
			return domainErrors.OrderNotFound{UUID: orderUUID}
		}
	}

	return nil
}

func (o *OrderRepository) FindOrder(orderUUID uuid.UUID) (utils.SharedRef[domain.Order], error) {
	orders, err := o.FindOrders([]uuid.UUID{orderUUID})
	if err != nil {
		return utils.SharedRef[domain.Order]{}, err
	}

	return orders[0], err
}

func (o *OrderRepository) FindOrders(orderUUIDs []uuid.UUID) ([]utils.SharedRef[domain.Order], error) {
	ctx := context.Background()
	queries := db.New(o.coonPool)

	pgResult, err := queries.FetchOrders(ctx, orderUUIDs)
	if err != nil {
		return nil, err
	}

	orders, err := mapper.ConvertPgOrderAggregatesToModels(pgResult)
	if err != nil {
		return nil, err
	}

	if len(orderUUIDs) == len(orders) {
		return orders, nil
	}

	return nil, getNotFoundError(orderUUIDs, orders)
}

func (o *OrderRepository) FindActiveOrderByTableIDForUpdate(tx domain.Transaction, tableID string) (utils.SharedRef[domain.Order], error) {
	ctx := context.Background()
	queries := db.New(o.coonPool)

	pgResult, err := queries.GetActiveOrderUUIDdByTableID(ctx, tableID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return utils.SharedRef[domain.Order]{}, domainErrors.OrderNotFoundByTableID{TableID: tableID}
		}

		return utils.SharedRef[domain.Order]{}, err
	}

	return o.FindOrderForUpdate(tx, pgResult)
}

func (o *OrderRepository) FindActiveOrderByCustomerUUID(customerUUID uuid.UUID) (utils.SharedRef[domain.Order], error) {
	ctx := context.Background()
	queries := db.New(o.coonPool)

	pgResult, err := queries.GetActiveOrderUUIDByCustomerUUID(ctx, customerUUID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return utils.SharedRef[domain.Order]{}, domainErrors.OrderNotFoundByCustomerUUID{CustomerUUID: customerUUID}
		}

		return utils.SharedRef[domain.Order]{}, err
	}

	return o.FindOrder(pgResult)
}

func (o *OrderRepository) FindOrderForUpdate(tx domain.Transaction, orderUUID uuid.UUID) (utils.SharedRef[domain.Order], error) {
	orders, err := o.FindOrdersForUpdate(tx, []uuid.UUID{orderUUID})
	if err != nil {
		return utils.SharedRef[domain.Order]{}, err
	}

	return orders[0], err
}

func (o *OrderRepository) FindOrdersForUpdate(tx domain.Transaction, orderUUIDs []uuid.UUID) ([]utils.SharedRef[domain.Order], error) {
	ctx := context.Background()
	trx := tx.(*pgTx)
	queries := db.New(o.coonPool).WithTx(trx.tx)

	pgResult, err := queries.FetchOrdersForUpdate(ctx, orderUUIDs)
	if err != nil {
		return nil, err
	}

	orders, err := mapper.ConvertPgOrderAggregatesToModels(pgResult)
	if err != nil {
		return nil, err
	}

	if len(orderUUIDs) == len(orders) {
		return orders, nil
	}

	return nil, getNotFoundError(orderUUIDs, orders)
}

func (o *OrderRepository) FindOrdersByPlaceUUID(placeUUID uuid.UUID, isActive bool) ([]utils.SharedRef[domain.Order], error) {
	ctx := context.Background()
	queries := db.New(o.coonPool)

	orderUUIDs, err := queries.GetOrderUUIDsByPlaceUUID(ctx, db.GetOrderUUIDsByPlaceUUIDParams{
		Column1: placeUUID,
		Column2: isActive,
	})
	if err != nil {
		return nil, err
	}

	return o.FindOrders(orderUUIDs)
}

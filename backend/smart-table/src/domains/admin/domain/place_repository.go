package domain

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/google/uuid"
	"github.com/smart-table/src/utils"
)

type PlaceRepository interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Commit(ctx context.Context, tx pgx.Tx) error

	Save(ctx context.Context, tx pgx.Tx, place utils.SharedRef[Place]) error

	CheckAddressExist(ctx context.Context, address string, restaurantUUID uuid.UUID) (bool, error)
}

package utils

import (
	"github.com/smart-table/src/logging"
	"go.uber.org/zap"
)

type TransactionManager[T any] interface {
	Commit(tx T) error
	Rollback(tx T) error
}

func Rollback[T any](
	repo TransactionManager[T],
	tx T,
) {
	rollbackErr := repo.Rollback(tx)
	if rollbackErr != nil {
		logging.GetLogger().Error("rollback failed due to error", zap.Error(rollbackErr))
	}
}

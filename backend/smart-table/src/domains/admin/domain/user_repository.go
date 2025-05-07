package domain

import (
	"github.com/google/uuid"
	"github.com/smart-table/src/utils"
)

type UserRepository interface {
	Begin() (Transaction, error)
	Commit(tx Transaction) error
	Rollback(tx Transaction) error

	Save(tx Transaction, user utils.SharedRef[User]) error

	FindUserByUUID(uuid uuid.UUID) (utils.SharedRef[User], error)
	FindUserByLogin(userLogin string) (utils.SharedRef[User], error)
	FindUserByLoginOrTgLogin(userLogin, tgLogin string) (utils.SharedRef[User], error)
}

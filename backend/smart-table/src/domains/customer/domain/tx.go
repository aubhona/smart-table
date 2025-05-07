package domain

type Transaction interface {
	Commit() error
	Rollback() error
}

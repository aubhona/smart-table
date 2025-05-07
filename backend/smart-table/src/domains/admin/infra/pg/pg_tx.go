package pg

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type pgTx struct {
	tx  pgx.Tx
	ctx context.Context
}

func (p *pgTx) Commit() error {
	if p.tx == nil {
		return nil
	}

	err := p.tx.Commit(p.ctx)
	if err == nil {
		p.tx = nil
	}

	return err
}

func (p *pgTx) Rollback() error {
	if p.tx == nil {
		return nil
	}

	err := p.tx.Rollback(p.ctx)
	if err == nil {
		p.tx = nil
	}

	return err
}

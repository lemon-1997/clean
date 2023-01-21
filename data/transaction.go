package data

import (
	"context"
	"database/sql"
	"github.com/lemon-1997/clean/usecase"
)

func NewTransaction(d *Data) usecase.Transaction {
	return d
}

type DbTx interface {
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

package data

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lemon-1997/clean/config"
)

type contextTxKey struct{}

type Data struct {
	db *sql.DB
}

func NewData(conf *config.DB) (*Data, error) {
	db, err := sql.Open(conf.Driver, conf.Source)
	if err != nil {
		return nil, err
	}
	defer func() { _ = db.Close() }()
	return &Data{
		db: db,
	}, nil
}

func (d *Data) InTx(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	err = fn(context.WithValue(ctx, contextTxKey{}, tx))

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (d *Data) DB(ctx context.Context) DbTx {
	tx, ok := ctx.Value(contextTxKey{}).(*sql.Tx)
	if ok {
		return tx
	}
	return d.db
}

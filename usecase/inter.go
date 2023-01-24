package usecase

import (
	"context"
	"github.com/lemon-1997/clean/entity"
)

type Transaction interface {
	InTx(context.Context, func(ctx context.Context) error) error
}

type OrderRepo interface {
	CreateOrder(context.Context, *entity.Order) (int64, error)
	FindOrderByID(context.Context, int64) (*entity.Order, error)
	UpdateOrderByID(context.Context, *entity.Order, int64) error
	DeleteOrderByID(context.Context, int64) error
}

type PayRepo interface {
	CreatePay(context.Context, *entity.Pay) (int64, error)
	DeletePayByID(context.Context, int64) error
}

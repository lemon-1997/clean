package data

import (
	"context"
	"github.com/lemon-1997/clean/entity"
	"github.com/lemon-1997/clean/usecase"
)

type orderRepo struct {
	data *Data
}

func NewOrderRepo(data *Data) usecase.OrderRepo {
	return &orderRepo{data: data}
}

func (r *orderRepo) CreateOrder(ctx context.Context, order *entity.Order) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (r *orderRepo) FindOrderByID(ctx context.Context, i int64) (*entity.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (r *orderRepo) UpdateOrderByID(ctx context.Context, order *entity.Order, i int64) error {
	//TODO implement me
	panic("implement me")
}

func (r *orderRepo) DeleteOrderByID(ctx context.Context, i int64) error {
	//TODO implement me
	panic("implement me")
}

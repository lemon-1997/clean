package data

import (
	"context"
	"github.com/lemon-1997/clean/entity"
	"github.com/lemon-1997/clean/usecase"
)

type payRepo struct {
	data *Data
}

func NewPayRepo(data *Data) usecase.PayRepo {
	return &payRepo{data: data}
}

func (payRepo) CreatePay(ctx context.Context, pay *entity.Pay) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (payRepo) DeletePayByID(ctx context.Context, i int64) error {
	//TODO implement me
	panic("implement me")
}

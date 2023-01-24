package usecase

import (
	"context"
	"github.com/lemon-1997/clean/entity"
)

type OrderUseCase struct {
	order OrderRepo
	pay   PayRepo
	tm    Transaction
}

func NewOrder(order OrderRepo, pay PayRepo, tm Transaction) *OrderUseCase {
	return &OrderUseCase{
		order: order,
		pay:   pay,
		tm:    tm,
	}
}

func (uc *OrderUseCase) CreateOrder(ctx context.Context, order *entity.Order, pay *entity.Pay) (orderID int64, err error) {
	err = uc.tm.InTx(ctx, func(ctx context.Context) error {
		if orderID, err = uc.order.CreateOrder(ctx, order); err != nil {
			return err
		}
		if _, err = uc.pay.CreatePay(ctx, pay); err != nil {
			return err
		}
		return nil
	})
	return
}

func (uc *OrderUseCase) GetOrder(ctx context.Context, id int64) (*entity.Order, error) {
	return uc.order.FindOrderByID(ctx, id)
}

func (uc *OrderUseCase) UpdateOrderStatus(ctx context.Context, id int64, status int) error {
	order, err := uc.order.FindOrderByID(ctx, id)
	if err != nil {
		return err
	}
	order.OrderStatus = status
	return uc.order.UpdateOrderByID(ctx, order, id)
}

func (uc *OrderUseCase) DeleteOrder(ctx context.Context, id int64) error {
	return uc.tm.InTx(ctx, func(ctx context.Context) error {
		if err := uc.order.DeleteOrderByID(ctx, id); err != nil {
			return err
		}
		if err := uc.pay.DeletePayByID(ctx, id); err != nil {
			return err
		}
		return nil
	})
}

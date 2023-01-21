package entity

import "time"

type Order struct {
	ID            int64     `json:"id"`
	OrderNo       string    `json:"order_no"`       //  订单编号
	OrderStatus   int       `json:"order_status"`   //  订单状态 0未付款,1已付款,2已发货,3已签收,-1退货申请,-2退货中,-3已退货,-4取消交易
	ProductCount  int64     `json:"product_count"`  //  商品数量
	ProductAmount float64   `json:"product_amount"` //  商品总价
	OrderAmount   float64   `json:"order_amount"`   //  实际付款金额
	PayTime       time.Time `json:"pay_time"`       //  付款时间
	DeliveryTime  time.Time `json:"delivery_time"`  //  发货时间
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

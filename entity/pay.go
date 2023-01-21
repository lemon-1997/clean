package entity

import "time"

type Pay struct {
	ID        int64     `json:"id"`
	OrderNo   string    `json:"order_no"` //  交易单号
	Amount    float64   `json:"amount"`   //  交易金额
	Source    string    `json:"source"`   //  支付来源 1.微信 2.支付宝
	Status    int64     `json:"status"`   //  支付状态 -1：取消 0 未完成 1已完成 -2:异常
	Remark    string    `json:"remark"`   //  备注
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

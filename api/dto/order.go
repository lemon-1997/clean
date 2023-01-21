package dto

type OrderCreateReq struct {
	OrderNo       string  `json:"order_no"`
	OrderStatus   int     `json:"order_status"`
	ProductCount  int64   `json:"product_count"`
	ProductAmount float64 `json:"product_amount"`
	OrderAmount   float64 `json:"order_amount"`
	DeliveryTime  int64   `json:"delivery_time"`
	PayTime       int64   `json:"pay_time"`
	PayAmount     float64 `json:"pay_amount"`
	PaySource     string  `json:"pay_source"`
	PayStatus     int64   `json:"status"`
	PayRemark     string  `json:"remark"`
}

type OrderUpdateReq struct {
	ID          int64 `json:"id"`
	OrderStatus int   `json:"order_status"`
}

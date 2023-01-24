package handler

import (
	"encoding/json"
	"github.com/lemon-1997/clean/api/dto"
	"github.com/lemon-1997/clean/entity"
	"github.com/lemon-1997/clean/usecase"
	"io"
	"net/http"
	"strconv"
	"time"
)

type Order struct {
	order *usecase.OrderUseCase
}

func NewOrder(uc *usecase.OrderUseCase) *Order {
	return &Order{order: uc}
}

func (h *Order) Handler() map[string]http.HandlerFunc {
	return map[string]http.HandlerFunc{
		"/order": h.handler,
	}
}

func (h *Order) handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.orderGet(w, r)
	case http.MethodPost:
		h.orderCreate(w, r)
	case http.MethodPut:
		h.orderUpdate(w, r)
	case http.MethodDelete:
		h.orderDelete(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (h *Order) orderGet(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	order, err := h.order.GetOrder(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func (h *Order) orderCreate(w http.ResponseWriter, r *http.Request) {
	var req dto.OrderCreateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	id, err := h.order.CreateOrder(r.Context(), transCreateReqToOrder(&req), transCreateReqToPay(&req))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	reply := &dto.OrderCreateReply{OrderID: id}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reply)
}

func (h *Order) orderUpdate(w http.ResponseWriter, r *http.Request) {
	var req dto.OrderUpdateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if err := h.order.UpdateOrderStatus(r.Context(), req.ID, req.OrderStatus); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	io.WriteString(w, "ok")
}

func (h *Order) orderDelete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err = h.order.DeleteOrder(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	io.WriteString(w, "ok")
}

func transCreateReqToOrder(req *dto.OrderCreateReq) *entity.Order {
	return &entity.Order{
		OrderNo:       req.OrderNo,
		OrderStatus:   req.OrderStatus,
		ProductCount:  req.ProductCount,
		ProductAmount: req.ProductAmount,
		OrderAmount:   req.OrderAmount,
		PayTime:       time.Unix(req.PayTime, 0),
		DeliveryTime:  time.Unix(req.DeliveryTime, 0),
	}
}

func transCreateReqToPay(req *dto.OrderCreateReq) *entity.Pay {
	return &entity.Pay{
		OrderNo: req.OrderNo,
		Amount:  req.PayAmount,
		Source:  req.PaySource,
		Status:  req.PayStatus,
		Remark:  req.PayRemark,
	}
}

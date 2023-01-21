package main

import (
	"github.com/lemon-1997/clean/api"
	"github.com/lemon-1997/clean/api/handler"
	"github.com/lemon-1997/clean/config"
	"github.com/lemon-1997/clean/data"
	"github.com/lemon-1997/clean/usecase"
)

func main() {
	c := config.New()
	d, err := data.NewData(c.DB)
	if err != nil {
		panic(err)
	}
	orderRepo := data.NewOrderRepo(d)
	payRepo := data.NewPayRepo(d)
	tm := data.NewTransaction(d)
	uc := usecase.NewOrder(orderRepo, payRepo, tm)
	orderHandler := handler.NewOrder(uc)
	s := api.NewServer(c.Server, orderHandler)
	if err = s.Run(); err != nil {
		panic(err)
	}
}

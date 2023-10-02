package main

import (
	"github.com/rdj68/marketplace-project/server/db"
	service "github.com/rdj68/marketplace-project/server/services"
)

type Inject struct {
	AuthDb *db.Db

	ShopService    *service.ShopService
	ProductService *service.ProductService
}

func NewInject() *Inject {
	inj := &Inject{}
	inj.AuthDb = &db.Db{}

	inj.ShopService = service.NewShopService(inj.AuthDb)
	inj.ProductService = service.NewProductService(inj.AuthDb)

	return inj
}

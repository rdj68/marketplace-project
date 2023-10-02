package db

import (
	"github.com/SaiNageswarS/go-api-boot/odm"
	"github.com/rdj68/marketplace-project/server/model"
)

type Db struct{}

func (a *Db) ProductRepo() *productRepository {
	baseRepo := odm.AbstractRepository[model.Product]{
		Database:       "marketplace",
		CollectionName: "product",
	}

	return &productRepository{baseRepo}
}

func (a *Db) ShopRepo() *shopRepository {
	baseRepo := odm.AbstractRepository[model.Shop]{
		Database:       "marketplace",
		CollectionName: "shop",
	}

	return &shopRepository{baseRepo}
}

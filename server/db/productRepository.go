package db

import (
	"strings"

	"github.com/SaiNageswarS/go-api-boot/odm"
	pb "github.com/rdj68/marketplace-project/proto"
	"github.com/rdj68/marketplace-project/server/model"
	"go.mongodb.org/mongo-driver/bson"
)

type productRepository struct {
	odm.AbstractRepository[model.Product]
}

func (p *productRepository) AddProduct(product *model.Product) (*model.Product, error) {

	err := <-p.Save(product)
	return product, err

}

func (p *productRepository) FindProduct(in *pb.ProductSearchParam) ([]model.Product, error) {

	filter := bson.D{
		{Key: "$and",
			Value: bson.A{
				bson.D{{Key: "product_name", Value: strings.ToLower(in.GetName())}},
				bson.D{{Key: "category", Value: strings.ToLower(in.GetCategory())}},
				bson.D{{Key: "price", Value: bson.D{{Key: "$lt", Value: in.GetMaxPrice()}}}},
				bson.D{{Key: "price", Value: bson.D{{Key: "$gt", Value: in.GetMinPrice()}}}},
			},
		},
	}
	resultChan, errorChan := p.Find(filter.Map(), nil, 20, 0)
	product := []model.Product{}
	select {
	case productRes := <-resultChan:
		product = productRes
	case err := <-errorChan:
		return nil, err
	}
	return product, nil
}

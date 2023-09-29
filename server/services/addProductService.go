package service

import (
	"context"
	"strings"

	pb "github.com/rdj68/marketplace-project/proto"
	db "github.com/rdj68/marketplace-project/server/db"
	"github.com/rdj68/marketplace-project/server/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
*
Add product details to the database and return the product id
*
*/
func (s *ProductServiceServer) AddProduct(ctx context.Context, in *pb.NewProductData) (*pb.ProductId, error) {

	shopId, err := primitive.ObjectIDFromHex(in.ShopId)
	if err != nil {
		panic(err)
	}

	product := model.Product{
		Name:        strings.ToLower(in.Name),
		Description: in.Description,
		Category:    strings.ToLower(in.Category),
		Price:       in.Price,
		ImageURL:    in.ImageURL,
		ShopID:      shopId,
		Attributes:  in.Attributes,
	}
	res, err := db.AddProduct(&product)
	if err != nil {
		return nil, err
	}
	return res, nil

}

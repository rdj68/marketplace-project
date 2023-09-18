package controllers

import (
	"context"
	"fmt"

	pb "github.com/rdj68/marketplace-project/proto"
	"github.com/rdj68/marketplace-project/server/model"
	odm "github.com/rdj68/marketplace-project/server/odm"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
*
function to insert product data into the database
*
*/
func (s *MarketplaceServer) AddProduct(ctx context.Context, in *pb.NewProductData) (*pb.Id, error) {
	client := odm.GetClient()
	coll := client.Database("marketplace").Collection("products")

	shopId, err := primitive.ObjectIDFromHex(in.ShopId)
	if err != nil {
		panic(err)
	}
	product := model.Product{
		Name:        in.Name,
		Description: in.Description,
		Category:    in.Category,
		Price:       in.Price,
		ImageURL:    in.ImageURL,
		ShopID:      shopId,
		Attributes:  in.Attributes,
	}

	res, err := coll.InsertOne(context.TODO(), product)
	if err != nil {
		return nil, err
	}
	fmt.Println("", res)
	return &pb.Id{Id: res.InsertedID.(primitive.ObjectID).Hex()}, nil
}

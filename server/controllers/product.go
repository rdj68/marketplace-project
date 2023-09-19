package controllers

import (
	"context"
	"fmt"
	"log"

	pb "github.com/rdj68/marketplace-project/proto"
	"github.com/rdj68/marketplace-project/server/model"
	odm "github.com/rdj68/marketplace-project/server/odm"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
*
Add product details to the database and return the product id
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

func (s *MarketplaceServer) FindProduct(ctx context.Context, in *pb.ProductSearchParam) (*pb.MultipleProductData, error) {
	client := odm.GetClient()
	coll := client.Database("marketplace").Collection("products")

	filter := bson.D{
		{Key: "$and",
			Value: bson.A{
				bson.D{{Key: "product_name", Value: in.GetName()}},
				bson.D{{Key: "category", Value: in.GetCategory()}},
				bson.D{{Key: "price", Value: bson.D{{Key: "$lt", Value: in.GetMaxPrice()}}}},
				bson.D{{Key: "price", Value: bson.D{{Key: "$gt", Value: in.GetMinPrice()}}}},
			},
		},
	}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	var result []model.Product
	if err := cursor.All(context.TODO(), &result); err != nil {
		log.Print(err)
		return nil, err
	}
	var resultProductData []*pb.ProductData
	for _, x := range result {
		resultProductData = append(resultProductData,
			&pb.ProductData{Id: x.ID.Hex(), Name: x.Name, Description: x.Description,
				Category: x.Category, Price: x.Price, ImageURL: x.ImageURL,
				Attributes: x.Attributes, ShopId: x.ShopID.Hex()})
	}
	return &pb.MultipleProductData{ProductData: resultProductData}, nil
}

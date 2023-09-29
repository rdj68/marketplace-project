package db

import (
	"context"
	"log"
	"strings"

	"github.com/SaiNageswarS/go-api-boot/odm"
	pb "github.com/rdj68/marketplace-project/proto"
	"github.com/rdj68/marketplace-project/server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type productRepository struct {
	odm.AbstractRepository[model.Product]
}

func AddProduct(product *model.Product) (*pb.ProductId, error) {
	client := odm.GetClient()
	coll := client.Database("marketplace").Collection("products")

	res, err := coll.InsertOne(context.TODO(), product)
	if err != nil {
		return nil, err
	}
	return &pb.ProductId{ProductId: res.InsertedID.(primitive.ObjectID).Hex()}, nil

}

func FindProduct(in *pb.ProductSearchParam) (*pb.MultipleProductData, error) {
	client := odm.GetClient()
	coll := client.Database("marketplace").Collection("products")

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

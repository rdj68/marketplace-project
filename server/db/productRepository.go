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
	"go.mongodb.org/mongo-driver/mongo"
)

type productRepository struct {
	odm.AbstractRepository[model.Product]
}

func AddProduct(in *pb.NewProductData) (*mongo.InsertOneResult, error) {

	client := odm.GetClient()
	coll := client.Database("marketplace").Collection("products")

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

	res, err := coll.InsertOne(context.TODO(), product)

	return res, nil

}

func FindProduct(in *pb.ProductSearchParam) ([]model.Product, error) {
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
		return nil, err
	}

	return result, nil
}

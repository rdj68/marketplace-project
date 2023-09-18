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
	"go.mongodb.org/mongo-driver/mongo"
)

type MarketplaceServer struct {
	pb.UnimplementedMarketplaceServer
}

/*
*
function to insert shop data into the database
*
*/
func (s *MarketplaceServer) AddShop(ctx context.Context, in *pb.NewShopData) (*pb.Id, error) {
	client := odm.GetClient()
	coll := client.Database("marketplace").Collection("shop")

	//convert the string user list to object id user list
	var userList []primitive.ObjectID
	for _, x := range in.GetUsers() {
		id, err := primitive.ObjectIDFromHex(x)
		if err != nil {
			continue
		}
		userList = append(userList, id)
	}

	shop := &model.Shop{
		Name: in.GetName(),
		Location: model.Location{
			Type:        "Point",
			Coordinates: in.GetLocation(),
		},
		Users:            userList,
		OperationsTiming: in.GetOperationTiming(),
	}

	res, err := coll.InsertOne(context.TODO(), shop)
	if err != nil {
		return nil, err
	}
	fmt.Println("", res)
	return &pb.Id{Id: res.InsertedID.(primitive.ObjectID).Hex()}, nil
}

/*
*
function to find a shop by its id
*
*/
func (s *MarketplaceServer) FindShopById(ctx context.Context, in *pb.Id) (*pb.ShopData, error) {
	client := odm.GetClient()
	coll := client.Database("marketplace").Collection("shop")

	id, err := primitive.ObjectIDFromHex(in.GetId())
	if err != nil {
		log.Print(err)
		return nil, err
	}

	filter := bson.D{{Key: "_id", Value: id}}
	var result model.Shop
	err = coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No document Found")
		}
		return nil, err
	}
	return &pb.ShopData{}, nil
}

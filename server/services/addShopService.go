package service

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/SaiNageswarS/go-api-boot/odm"
	pb "github.com/rdj68/marketplace-project/proto"
	"github.com/rdj68/marketplace-project/server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ShopServiceServer struct {
	pb.UnimplementedShopServiceServer
}

// removing auth interceptor
func (u *ShopServiceServer) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

/*
*
Add shop details to databaase and return the objectid of the shop
*
*/
func (s *ShopServiceServer) AddShop(ctx context.Context, in *pb.NewShopData) (*pb.Id, error) {
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
		Name: strings.ToLower(in.GetName()),
		Location: model.Location{
			Type:        "Point",
			Coordinates: []float64{in.Location.GetLongitude(), in.Location.GetLattitude()},
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
Find shop details, map it to shopData and returns shopData. If no details found returns nil, error
*
*/
func (s *ShopServiceServer) FindShopById(ctx context.Context, in *pb.Id) (*pb.ShopData, error) {
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
	return &pb.ShopData{Name: result.Name, Id: result.ID.String(), Location: &pb.Location{Longitude: result.Location.Coordinates[0], Lattitude: result.Location.Coordinates[1]}, OperationTiming: result.OperationsTiming}, nil
}

/*
*
filter shops that are close to user's coordinates by max 100km
*
*/
func (s *ShopServiceServer) FindShopNearMe(ctx context.Context, in *pb.Location) (*pb.MultipleShopData, error) {
	client := odm.GetClient()
	coll := client.Database("marketplace").Collection("shop")
	diameter := 1_000_000

	//index database for geospatial queries
	indexModel := mongo.IndexModel{
		Keys: bson.D{{Key: "location", Value: "2dsphere"}},
	}
	_, err := coll.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		panic(err)
	}

	//query to find shops near the location
	filter := bson.M{
		"location": bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{in.GetLongitude(), in.GetLattitude()},
				},
				"$maxDistance": diameter, // Maximum distance in meters
			},
		},
	}
	cursor, err := coll.Find(context.Background(), filter)
	if err != nil {
		panic(err)
	}

	var result []model.Shop
	if err := cursor.All(context.Background(), &result); err != nil {
		panic(err)
	}
	//store the results in pb.shopdata array to share over grpc
	var resultsShop []*pb.ShopData
	for _, x := range result {
		resultsShop = append(resultsShop, &pb.ShopData{Name: x.Name, Location: &pb.Location{Longitude: x.Location.Coordinates[0], Lattitude: x.Location.Coordinates[1]}, OperationTiming: x.OperationsTiming})
	}
	return &pb.MultipleShopData{ShopData: resultsShop}, nil
}

package service

import (
	"context"

	"github.com/SaiNageswarS/go-api-boot/logger"
	pb "github.com/rdj68/marketplace-project/proto"
	db "github.com/rdj68/marketplace-project/server/db"
	"go.uber.org/zap"
)

/*
*
Find shop details, map it to shopData and returns shopData. If no details found returns nil, error
*
*/
func (s *ShopServiceServer) FindShopById(ctx context.Context, in *pb.Id) (*pb.ShopData, error) {

	res, err := db.FindShopById(in)
	if err != nil {
		logger.Error("Cannot Find Shop", zap.Error(err))
	}

	return &pb.ShopData{Name: res.Name, Id: res.ID.String(), Location: &pb.Location{Longitude: res.Location.Coordinates[0], Lattitude: res.Location.Coordinates[1]}, OperationTiming: res.OperationsTiming}, nil
}

/*
*
filter shops that are close to user's coordinates by max 100km
*
*/
func (s *ShopServiceServer) FindShopNearMe(ctx context.Context, in *pb.Location) (*pb.MultipleShopData, error) {

	res, err := db.FindShopNearMe(in)
	if err != nil {
		logger.Error("Error fetching shop", zap.Error(err))
	}

	//store the results in pb.shopdata array to share over grpc
	var resultsShop []*pb.ShopData
	for _, x := range res {
		resultsShop = append(resultsShop, &pb.ShopData{Name: x.Name, Location: &pb.Location{Longitude: x.Location.Coordinates[0], Lattitude: x.Location.Coordinates[1]}, OperationTiming: x.OperationsTiming})
	}
	return &pb.MultipleShopData{ShopData: resultsShop}, nil
}

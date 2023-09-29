package service

import (
	"context"

	"github.com/SaiNageswarS/go-api-boot/logger"
	pb "github.com/rdj68/marketplace-project/proto"
	db "github.com/rdj68/marketplace-project/server/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
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

	res, err := db.AddShop(in)
	if err != nil {
		logger.Error("Couldn't insert data", zap.Error(err))
	}
	return &pb.Id{Id: res.InsertedID.(primitive.ObjectID).Hex()}, nil
}

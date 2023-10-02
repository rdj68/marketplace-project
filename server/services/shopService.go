package service

import (
	"context"
	"strings"

	"github.com/SaiNageswarS/go-api-boot/logger"
	pb "github.com/rdj68/marketplace-project/proto"
	db "github.com/rdj68/marketplace-project/server/db"
	"github.com/rdj68/marketplace-project/server/model"
	"go.uber.org/zap"
)

type ShopService struct {
	pb.UnimplementedShopServiceServer
	db *db.Db
}

// removing auth interceptor
func (u *ShopService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

func NewShopService(db *db.Db) *ShopService {
	return &ShopService{
		db: db,
	}
}

/*
*
Add shop details to databaase and return the objectid of the shop
*
*/
func (s *ShopService) AddShop(ctx context.Context, shopProto *pb.NewShopData) (*pb.Id, error) {

	shop := getShopModel(shopProto)
	id, err := s.db.ShopRepo().AddShop(shop)

	if err != nil {
		logger.Error("Error inserting data", zap.Error(err))
	}
	return &pb.Id{Id: *id}, nil
}

/*
*
Find shop details, map it to shopData and returns shopData. If no details found returns nil, error
*
*/
func (s *ShopService) FindShopById(ctx context.Context, idProto *pb.Id) (*pb.ShopData, error) {

	shop, err := s.db.ShopRepo().FindShopById(idProto.Id)
	return getShopProto(shop), err
}

/*
*
filter shops that are close to user's coordinates by max 100km
*
*/
func (s *ShopService) FindShopNearMe(ctx context.Context, locationProto *pb.Location) (*pb.MultipleShopData, error) {
	diameter := 1_000_000
	shopsData, err := s.db.ShopRepo().FindShopNearMe(locationProto, diameter)

	if err != nil {
		logger.Error("Error fetching shop", zap.Error(err))
	}
	return getMultipleShopProto(shopsData), nil
}

func getShopModel(shopProto *pb.NewShopData) *model.Shop {

	shop := &model.Shop{
		Name:  strings.ToLower(shopProto.GetName()),
		Phone: shopProto.Phone,
		Location: model.Location{
			Type:        "Point",
			Coordinates: []float64{shopProto.Location.GetLongitude(), shopProto.Location.GetLattitude()},
		},
		Users:            shopProto.GetUsers(),
		OperationsTiming: shopProto.GetOperationTiming(),
	}
	return shop
}

func getShopProto(shopModel *model.Shop) *pb.ShopData {
	shopProto := &pb.ShopData{
		Name: shopModel.Name,
		Id:   shopModel.ID,
		Location: &pb.Location{
			Longitude: shopModel.Location.Coordinates[0],
			Lattitude: shopModel.Location.Coordinates[1],
		},
		OperationTiming: shopModel.OperationsTiming,
	}
	return shopProto
}

func getMultipleShopProto(shopData []model.Shop) *pb.MultipleShopData {
	var resultShops []*pb.ShopData
	for _, x := range shopData {
		resultShops = append(resultShops, &pb.ShopData{
			Name: x.Name,
			Location: &pb.Location{
				Longitude: x.Location.Coordinates[0],
				Lattitude: x.Location.Coordinates[1],
			},
			OperationTiming: x.OperationsTiming,
		},
		)
	}
	return &pb.MultipleShopData{ShopData: resultShops}
}

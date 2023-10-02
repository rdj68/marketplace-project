package db

import (
	"github.com/SaiNageswarS/go-api-boot/logger"
	"github.com/SaiNageswarS/go-api-boot/odm"
	pb "github.com/rdj68/marketplace-project/proto"
	"github.com/rdj68/marketplace-project/server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

// TODO implement abstract repository
type shopRepository struct {
	odm.AbstractRepository[model.Shop]
}

func (s *shopRepository) AddShop(shop *model.Shop) (*string, error) {

	err := <-s.Save(shop)

	return &shop.ID, err
}

/*
*
Find shop details, map it to shopData and returns shopData. If no details found returns nil, error
*
*/
func (s *shopRepository) FindShopById(id string) (*model.Shop, error) {

	shopResChan, errChan := s.FindOneById(id)

	shop := &model.Shop{}

	select {
	case shopRes := <-shopResChan:
		shop = shopRes
	case err := <-errChan:
		if err == mongo.ErrNoDocuments {
			logger.Info("No shop with the given id exists", zap.String("shopId: ", id))
			return nil, err
		}

		logger.Error("Failed getting shop", zap.String("userId", id), zap.Error(err))
	}
	return shop, nil
}

/*
*
filter shops that are close to user's coordinates by max 100km
*
*/
func (s *shopRepository) FindShopNearMe(in *pb.Location, diameter int) ([]model.Shop, error) {

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
	resultChan, errChan := s.Find(filter, nil, 20, 0)
	shop := []model.Shop{}
	select {
	case shopRes := <-resultChan:
		shop = shopRes
	case err := <-errChan:
		if err == mongo.ErrNoDocuments {
			logger.Info("No shop nearby")
			return nil, err
		}

		logger.Error("Failed getting shop", zap.Error(err))
	}

	return shop, nil
}

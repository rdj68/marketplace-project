package service

import (
	"context"

	"github.com/SaiNageswarS/go-api-boot/logger"
	pb "github.com/rdj68/marketplace-project/proto"
	db "github.com/rdj68/marketplace-project/server/db"
	"go.uber.org/zap"
)

type ProductServiceServer struct {
	pb.UnimplementedProductServiceServer
}

// removing auth interceptor
func (u *ProductServiceServer) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

func (s *ProductServiceServer) FindProduct(ctx context.Context, in *pb.ProductSearchParam) (*pb.MultipleProductData, error) {
	res, err := db.FindProduct(in)
	if err != nil {
		logger.Error("Error fetching product", zap.Error(err))
	}
	var resultProductData []*pb.ProductData
	for _, x := range res {
		resultProductData = append(resultProductData,
			&pb.ProductData{Id: x.ID.Hex(), Name: x.Name, Description: x.Description,
				Category: x.Category, Price: x.Price, ImageURL: x.ImageURL,
				Attributes: x.Attributes, ShopId: x.ShopID.Hex()})
	}
	return &pb.MultipleProductData{ProductData: resultProductData}, err
}

package service

import (
	"context"
	"strings"

	"github.com/SaiNageswarS/go-api-boot/logger"
	pb "github.com/rdj68/marketplace-project/proto"
	db "github.com/rdj68/marketplace-project/server/db"
	"github.com/rdj68/marketplace-project/server/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type ProductService struct {
	pb.UnimplementedProductServiceServer
	db *db.Db
}

func NewProductService(db *db.Db) *ProductService {
	return &ProductService{
		db: db,
	}
}

// removing auth interceptor
func (s *ProductService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

/*
*
Add product details to the database and return the product id
*
*/
func (s *ProductService) AddProduct(ctx context.Context, productProto *pb.NewProductData) (*pb.ProductId, error) {

	productModel := getProductModel(productProto)
	productAdded, err := s.db.ProductRepo().AddProduct(productModel)
	if err != nil {
		logger.Info("Error adding product", zap.Error(err))
		return nil, err
	}
	return &pb.ProductId{ProductId: productAdded.Id()}, nil
}

func (s *ProductService) FindProduct(ctx context.Context, searchParameters *pb.ProductSearchParam) (*pb.MultipleProductData, error) {
	productsData, err := s.db.ProductRepo().FindProduct(searchParameters)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Info("No shop nearby")
			return nil, err
		}

		logger.Error("Failed getting shop", zap.Error(err))
	}

	return getMultipleProductProto(productsData), nil

}

func getProductModel(productProto *pb.NewProductData) *model.Product {
	product := &model.Product{
		Name:        strings.ToLower(productProto.Name),
		Description: productProto.Description,
		Category:    strings.ToLower(productProto.Category),
		Price:       productProto.Price,
		ImageURL:    productProto.ImageURL,
		ShopID:      productProto.ShopId,
		Attributes:  productProto.Attributes,
	}
	return product
}
func getMultipleProductProto(productModel []model.Product) *pb.MultipleProductData {
	productProto := []*pb.ProductData{}
	for _, product := range productModel {
		productProto = append(productProto,
			&pb.ProductData{Id: product.ID, Name: product.Name, Description: product.Description,
				Category: product.Category, Price: product.Price, ImageURL: product.ImageURL,
				Attributes: product.Attributes, ShopId: product.ShopID})
	}
	return &pb.MultipleProductData{ProductData: productProto}
}

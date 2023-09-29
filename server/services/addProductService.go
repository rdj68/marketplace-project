package service

import (
	"context"

	pb "github.com/rdj68/marketplace-project/proto"
	db "github.com/rdj68/marketplace-project/server/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
*
Add product details to the database and return the product id
*
*/
func (s *ProductServiceServer) AddProduct(ctx context.Context, in *pb.NewProductData) (*pb.ProductId, error) {

	res, err := db.AddProduct(in)
	if err != nil {
		return nil, err
	}
	return &pb.ProductId{ProductId: res.InsertedID.(primitive.ObjectID).Hex()}, nil

}

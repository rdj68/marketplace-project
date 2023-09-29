package service

import (
	"context"

	pb "github.com/rdj68/marketplace-project/proto"

	db "github.com/rdj68/marketplace-project/server/db"
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

	return res, err
}

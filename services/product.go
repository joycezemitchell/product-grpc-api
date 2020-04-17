package main

import (
	"context"

	productpb "github.com/joycezemitchell/product-grpc-api/proto"
)

// Server - gRPC product server
type Server struct {
}

func (*server) CreateProduct(ctx context.Context, req *productpb.CreateProductRequest) (*productpb.CreateProductResponse, error) {
	return nil, nil
}

func (*server) UpdateProduct(ctx context.Context, req *productpb.UpdateProductRequest) (*productpb.UpdateProductResponse, error) {
	return nil, nil
}

func (*server) DeleteProduct(ctx context.Context, req *productpb.DeleteProductRequest) (*productpb.DeleteProductResponse, error) {
	return nil, nil
}

func (*server) ReadProduct(ctx context.Context, req *productpb.ReadProductRequest) (*productpb.ReadProductResponse, error) {
	return nil, nil
}

func (*server) ListProduct(req *productpb.ListProductRequest, stream productpb.ProductService_ListProductServer) error {

}

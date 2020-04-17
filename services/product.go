package services

import (
	"context"

	productpb "github.com/joycezemitchell/product-grpc-api/proto"
)

// Server - gRPC product server
type Server struct {
}

// CreateProduct Method
func (*Server) CreateProduct(ctx context.Context, req *productpb.CreateProductRequest) (*productpb.CreateProductResponse, error) {
	return nil, nil
}

// UpdateProduct Method
func (*Server) UpdateProduct(ctx context.Context, req *productpb.UpdateProductRequest) (*productpb.UpdateProductResponse, error) {
	return nil, nil
}

// DeleteProduct Method
func (*Server) DeleteProduct(ctx context.Context, req *productpb.DeleteProductRequest) (*productpb.DeleteProductResponse, error) {
	return nil, nil
}

// ReadProduct Method
func (*Server) ReadProduct(ctx context.Context, req *productpb.ReadProductRequest) (*productpb.ReadProductResponse, error) {
	return nil, nil
}

// ListProduct Method
func (*Server) ListProduct(req *productpb.ListProductRequest, stream productpb.ProductService_ListProductServer) error {

}

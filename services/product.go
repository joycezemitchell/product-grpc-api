package services

import (
	"context"
	"fmt"

	config "github.com/joycezemitchell/product-grpc-api/config"
	models "github.com/joycezemitchell/product-grpc-api/models"
	productpb "github.com/joycezemitchell/product-grpc-api/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server - gRPC product server
type Server struct {
}

// CreateProduct Method
func (*Server) CreateProduct(ctx context.Context, req *productpb.CreateProductRequest) (*productpb.CreateProductResponse, error) {
	fmt.Println("Create product request")
	p := req.GetProduct()

	categories := []models.Category{}
	variants := []models.Variant{}
	medias := []models.Media{}
	options := []models.ProductOption{}

	for _, v := range p.GetVariant() {

		for _, o := range v.GetProductOption() {
			options = append(options, models.ProductOption{
				ID:    primitive.NewObjectID(),
				Name:  o.GetName(),
				Type:  o.GetType(),
				Value: o.GetValue(),
			})
		}

		for _, m := range v.GetMedia() {
			medias = append(medias, models.Media{
				ID:  primitive.NewObjectID(),
				Src: m.GetSrc(),
			})
		}

		inventory := models.Inventory{
			Quantity:        v.GetInventory().GetQuantity(),
			ContinueSelling: v.GetInventory().GetContinueSelling(),
		}

		shipping := models.Shipping{
			Weight:     v.GetShipping().GetWeight(),
			WeightUnit: v.GetShipping().GetWeightUnit(),
		}

		variants = append(variants, models.Variant{
			ID:         primitive.NewObjectID(),
			Title:      v.GetTitle(),
			Option:     options,
			Media:      medias,
			Inventory:  inventory,
			Shipping:   shipping,
			Price:      v.GetPrice(),
			SalesPrice: v.GetSalesPrice(),
			Sku:        v.GetSku(),
			Taxable:    v.GetTaxable(),
			Positon:    v.GetPositon(),
			CreatedAt:  v.GetCreatedAt(),
			UpdatedAt:  v.GetUpdatedAt(),
		})
	}

	for _, c := range p.GetCategory() {
		categories = append(categories, models.Category{
			ID:   primitive.NewObjectID(),
			Name: c.GetName(),
		})
	}

	data := models.Product{
		Title:     p.GetTitle(),
		Seo:       p.GetSeo(),
		URL:       p.GetUrl(),
		Vendor:    p.GetVendor(),
		Tags:      p.GetTags(),
		Variant:   variants,
		Category:  categories,
		CreatedAt: "xxxxx",
		UpdatedAt: "xxxxx",
	}

	res, err := config.Collection.InsertOne(context.Background(), data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot convert to OID"),
		)
	}
	return &productpb.CreateProductResponse{
		Product: &productpb.Product{
			Id:     oid.Hex(),
			Title:  p.GetTitle(),
			Seo:    p.GetSeo(),
			Vendor: p.GetVendor(),
			Tags:   p.GetTags(),
		},
	}, nil

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
	return nil
}

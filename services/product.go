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
	"gopkg.in/mgo.v2/bson"
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

	fmt.Println("Product was created:", oid.Hex())
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
	fmt.Println("Updated product request")
	p := req.GetProduct()

	oid, err := primitive.ObjectIDFromHex(p.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID"),
		)
	}

	// create an empty struct
	data := &models.Product{}
	filter := bson.M{"_id": oid}

	res := config.Collection.FindOne(context.Background(), filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find blog with specified ID: %v", err),
		)
	}

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

	data.Title = p.GetTitle()
	data.Seo = p.GetSeo()
	data.URL = p.GetUrl()
	data.Vendor = p.GetVendor()
	data.Tags = p.GetTags()
	data.Variant = variants
	data.Category = categories
	data.UpdatedAt = "xxxxx"

	_, updateErr := config.Collection.ReplaceOne(context.Background(), filter, data)
	if updateErr != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot update object in MongoDB: %v", updateErr),
		)
	}

	return &productpb.UpdateProductResponse{
		Product: &productpb.Product{
			Id:     oid.Hex(),
			Title:  p.GetTitle(),
			Seo:    p.GetSeo(),
			Vendor: p.GetVendor(),
			Tags:   p.GetTags(),
		},
	}, nil

}

// DeleteProduct Method
func (*Server) DeleteProduct(ctx context.Context, req *productpb.DeleteProductRequest) (*productpb.DeleteProductResponse, error) {
	fmt.Println("Delete product request")
	oid, err := primitive.ObjectIDFromHex(req.GetProductId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID"),
		)
	}

	filter := bson.M{"_id": oid}

	res, err := config.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot delete object in MongoDB: %v", err),
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find product in MongoDB: %v", err),
		)
	}

	return &productpb.DeleteProductResponse{ProductId: req.GetProductId()}, nil
}

// ReadProduct Method
func (*Server) ReadProduct(ctx context.Context, req *productpb.ReadProductRequest) (*productpb.ReadProductResponse, error) {
	fmt.Println("Read Product request")

	ProductID := req.GetProductId()
	oid, err := primitive.ObjectIDFromHex(ProductID)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID"),
		)
	}

	// create an empty struct
	data := &models.Product{}
	filter := bson.M{"_id": oid}

	res := config.Collection.FindOne(context.Background(), filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find product with specified ID: %v", err),
		)
	}

	categories := []*productpb.Product_Category{}
	variants := []*productpb.Product_Variant{}
	medias := []*productpb.Product_Media{}
	options := []*productpb.Product_ProductOption{}

	for _, v := range data.Variant {

		for _, o := range v.Option {
			options = append(options, &productpb.Product_ProductOption{
				Id:    o.ID.Hex(),
				Name:  o.Name,
				Type:  o.Type,
				Value: o.Value,
			})
		}

		for _, m := range v.Media {
			medias = append(medias, &productpb.Product_Media{
				Id:  m.ID.Hex(),
				Src: m.Src,
			})
		}

		inventory := &productpb.Product_Inventory{
			Quantity:        v.Inventory.Quantity,
			ContinueSelling: v.Inventory.ContinueSelling,
		}

		shipping := &productpb.Product_Shipping{
			Weight:     v.Shipping.Weight,
			WeightUnit: v.Shipping.WeightUnit,
		}

		variants = append(variants, &productpb.Product_Variant{
			Title:         v.Title,
			ProductOption: options,
			Media:         medias,
			Inventory:     inventory,
			Shipping:      shipping,
			Price:         v.Price,
			SalesPrice:    v.SalesPrice,
			Sku:           v.Sku,
			Taxable:       v.Taxable,
			Positon:       v.Positon,
			CreatedAt:     v.CreatedAt,
			UpdatedAt:     v.UpdatedAt,
		})
	}

	for _, c := range data.Category {
		categories = append(categories, &productpb.Product_Category{
			Id:   c.ID.Hex(),
			Name: c.Name,
		})
	}

	return &productpb.ReadProductResponse{
		Product: &productpb.Product{
			Id:        oid.Hex(),
			Title:     data.Title,
			Seo:       data.Seo,
			Url:       data.URL,
			Vendor:    data.Vendor,
			Tags:      data.Tags,
			Variant:   variants,
			Category:  categories,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
		},
	}, nil
}

// ListProduct Method
func (*Server) ListProduct(req *productpb.ListProductRequest, stream productpb.ProductService_ListProductServer) error {
	fmt.Println("List product request")

	cur, err := config.Collection.Find(context.Background(), primitive.D{{}})
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &models.Product{}
		err := cur.Decode(data)
		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding data from MongoDB: %v", err),
			)

		}

		stream.Send(&productpb.ListProductResponse{
			Product: &productpb.Product{
				Id:    data.ID.Hex(),
				Title: data.Title,
			},
		})
	}

	if err := cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	return nil
}

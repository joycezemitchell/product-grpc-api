package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Product Data Structure
type Product struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title"`
	Seo       string             `bson:"seo"`
	URL       string             `bson:"url"`
	Vendor    string             `bson:"vendor"`
	Tags      string             `bson:"tags"`
	Variant   []Variant          `bson:"variant"`
	Category  []Category         `bson:"category"`
	CreatedAt string             `bson:"createdAt"`
	UpdatedAt string             `bson:"updatedAt"`
}

// Variant Data Structure
type Variant struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Title      string             `bson:"title"`
	Option     []ProductOption    `bson:"options"`
	Media      []Media            `bson:"media"`
	Inventory  Inventory          `bson:"inventory"`
	Shipping   Shipping           `bson:"shipping"`
	Price      string             `bson:"price"`
	SalesPrice string             `bson:"salesPrice"`
	Sku        string             `bson:"sku"`
	Taxable    string             `bson:"taxable"`
	Positon    string             `bson:"position"`
	CreatedAt  string             `bson:"createdAt"`
	UpdatedAt  string             `bson:"updatedAt"`
}

// ProductOption  Data Structure
type ProductOption struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Type  string             `bson:"type"`
	Value string             `bson:"value"`
}

// Media  Data Structure
type Media struct {
	ID  primitive.ObjectID `bson:"_id,omitempty"`
	Src string             `bson:"src"`
}

// Inventory Data Structure
type Inventory struct {
	Quantity        string `bson:"Quantity"`
	ContinueSelling string `bson:"continueSelling"`
}

// Shipping Data Structure
type Shipping struct {
	Weight     string `bson:"weight"`
	WeightUnit string `bson:"weightUnit"`
}

// Category Data Structure
type Category struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"title"`
}

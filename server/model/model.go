package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Shop model
type Shop struct {
	ID               primitive.ObjectID   `bson:"_id,omitempty"`
	Name             string               `bson:"shop_name"`
	Location         Location             `bson:"location"`
	OperationsTiming []string             `bson:"operations_timing"`
	Users            []primitive.ObjectID `bson:"users"`
}

type Location struct {
	Type        string    `bson:"type"`
	Coordinates []float64 `bson:"coordinates"`
}

// type OperationsTime struct {
// 	OpeningTime string `bson:"opening_time"`
// 	ClosingTime string `bson:"closing_time"`
// }

// Product represents product entity
type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Category    string             `bson:"category"`
	Price       float32            `bson:"price"`
	ImageURL    string             `bson:"image_url"`
	Attributes  map[string]string  `bson:"attributes"`
	ShopID      primitive.ObjectID `bson:"shop_id"`
}

// InventoryItem represents a product in the shop's inventory
type InventoryItem struct {
	ProductID primitive.ObjectID `bson:"product_id"`
	Quantity  int                `bson:"quantity"`
	InStock   bool               `bson:"in_stock"`
}

// ProductServiceability represents product serviceability
type ProductServiceability struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	ProductID   primitive.ObjectID `bson:"product_id"`
	Serviceable bool               `bson:"serviceable"`
}

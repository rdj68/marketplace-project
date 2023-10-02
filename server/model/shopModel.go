package model

import (
	"encoding/base64"
)

// Shop model
type Shop struct {
	ID               string   `bson:"_id"`
	Phone            string   `bson:"phone"`
	Name             string   `bson:"shop_name"`
	Location         Location `bson:"location"`
	OperationsTiming []string `bson:"operations_timing"`
	Users            []string `bson:"users"`
}

func (m *Shop) Id() string {
	return base64.StdEncoding.EncodeToString([]byte(m.Phone))

}

type Location struct {
	Type        string    `bson:"type"`
	Coordinates []float64 `bson:"coordinates"`
}

// type OperationsTime struct {
// 	OpeningTime string `bson:"opening_time"`
// 	ClosingTime string `bson:"closing_time"`
// }

// InventoryItem represents a product in the shop's inventory
type InventoryItem struct {
	ProductID string `bson:"product_id"`
	Quantity  int    `bson:"quantity"`
	InStock   bool   `bson:"in_stock"`
}

type ProductServiceability struct {
	ID          string `bson:"_id"`
	ProductID   string `bson:"product_id"`
	Serviceable bool   `bson:"serviceable"`
}

type User struct {
	ID       string   `bson:"_id"`
	Name     string   `bson:"user_name"`
	Location Location `bson:"location"`
}

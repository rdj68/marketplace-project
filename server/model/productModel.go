package model

import "encoding/base64"

type Product struct {
	ID          string            `bson:"_id"`
	Name        string            `bson:"product_name"`
	Description string            `bson:"description"`
	Category    string            `bson:"category"`
	Price       int32             `bson:"price"`
	ImageURL    string            `bson:"image_url"`
	Attributes  map[string]string `bson:"attributes"`
	ShopID      string            `bson:"shop_id"`
}

// Id function to assign unique id to each product
func (p *Product) Id() string {
	return base64.StdEncoding.EncodeToString([]byte(p.ShopID + p.Name + string(p.Price)))
}

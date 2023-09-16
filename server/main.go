package main

import (
	odm "github.com/market-place/server/odm"
)

func main() {
	coll := odm.GetClient("mongodb://localhost:27017")
	_ = coll
}

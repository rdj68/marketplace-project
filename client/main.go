package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/rdj68/marketplace-project/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewMarketplaceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	// shopData := &pb.NewShopData{Name: "rdjshop", Users: []string{"65082c66e9a0c3dd7b76cb9e"}, Location: []float64{4.2, 5.5}, OperationTiming: []string{"12:00", "1:00"}}
	// AddShop(ctx, c, shopData)
	// FindShopById(ctx, c, "65082c66e9a0c3dd7b76cb9e")
	FindShopNearMe(ctx, c, 4.2, 5.10005)

}

func AddShop(ctx context.Context, c pb.MarketplaceClient, shopData *pb.NewShopData) {
	r, err := c.AddShop(ctx, shopData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data inserted", r)
}

func FindShopById(ctx context.Context, c pb.MarketplaceClient, idString string) {
	id := &pb.Id{Id: idString}
	r, err := c.FindShopById(ctx, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data Found", r.GetId())
}

func FindShopNearMe(ctx context.Context, c pb.MarketplaceClient, lattitude float64, longitude float64) {
	loc := &pb.Location{Lattitude: lattitude, Longitude: longitude}
	r, err := c.FindShopNearMe(ctx, loc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Shop near you are: ")
	for _, x := range r.ShopData {
		fmt.Printf("Shop name: %v  operation timing: %v\n", x.Name, x.OperationTiming)
	}
}
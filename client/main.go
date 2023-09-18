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

	AddShop(ctx, c)

}

func AddShop(ctx context.Context, c pb.MarketplaceClient) {
	shopData := &pb.NewShopData{Name: "rdjshop", Users: []string{"alice", "bob"}, Location: []float64{4.2, 5.5}, OperationTiming: []string{"12:00", "1:00"}}
	r, err := c.AddShop(ctx, shopData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data inserted", r)
}

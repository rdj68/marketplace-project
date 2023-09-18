package main

import (
	"log"
	"net"

	"github.com/joho/godotenv"
	//"github.com/rdj68/marketplace-project/server/model"
	//odm "github.com/rdj68/marketplace-project/server/odm"
	pb "github.com/rdj68/marketplace-project/proto"
	ct "github.com/rdj68/marketplace-project/server/controllers"
	"google.golang.org/grpc"
)

const (
	port    = ":50051"
	network = "tcp"
)

func main() {
	//Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	//Listen at port 50051
	lis, err := net.Listen(network, port)
	if err != nil {
		log.Fatal(err)
	}

	//setup grpc server
	s := grpc.NewServer()
	pb.RegisterMarketplaceServer(s, &ct.MarketplaceServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}

}

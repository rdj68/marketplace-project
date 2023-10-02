package main

import (
	"log"

	"github.com/SaiNageswarS/go-api-boot/server"
	"github.com/joho/godotenv"
	pb "github.com/rdj68/marketplace-project/proto"
	"github.com/rs/cors"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPort = ":50051"
	webPort  = ":8081"
)

func main() {
	//Load environment variables
	if err := godotenv.Load("./.env"); err != nil {
		log.Println("No .env file found")
	}
	// Load secrets from Keyvault and config through godotenv.
	server.LoadSecretsIntoEnv(true)

	inject := NewInject()

	//create geospatial index for shop collection
	CreateGeospatialIndex("marketplace", "shop")

	corsConfig := cors.New(
		cors.Options{
			AllowedHeaders: []string{"*"},
		})
	bootServer := server.NewGoApiBoot(corsConfig)

	pb.RegisterProductServiceServer(bootServer.GrpcServer, inject.ProductService)
	pb.RegisterShopServiceServer(bootServer.GrpcServer, inject.ShopService)
	reflection.Register(bootServer.GrpcServer)
	bootServer.Start(grpcPort, webPort)

}

package main

import (
	"log"

	"github.com/SaiNageswarS/go-api-boot/server"
	"github.com/joho/godotenv"
	pb "github.com/rdj68/marketplace-project/proto"
	ct "github.com/rdj68/marketplace-project/server/services"
	"github.com/rs/cors"
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
	corsConfig := cors.New(
		cors.Options{
			AllowedHeaders: []string{"*"},
		})
	bootServer := server.NewGoApiBoot(corsConfig)

	pb.RegisterProductServiceServer(bootServer.GrpcServer, &ct.ProductServiceServer{})
	pb.RegisterShopServiceServer(bootServer.GrpcServer, &ct.ShopServiceServer{})
	bootServer.Start(grpcPort, webPort)

}

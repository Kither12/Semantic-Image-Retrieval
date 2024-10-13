package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"search-engine/config"
	"search-engine/database"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", viper.GetString("port")))
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()

	grpcServer := grpc.NewServer()
	log.Println("GRPC server started")

	db := database.CreateDatabase(context.Background())
	NewImageSevice(grpcServer, db)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err.Error())
	}
}

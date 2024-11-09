package main

import (
	"fmt"
	"log"
	"net"
	"search-engine/config"
	"search-engine/services"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {

	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", viper.GetString("PORT")))
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()

	grpcServer := grpc.NewServer()

	go func() {
		log.Printf("GRPC server started at port: %s\n", viper.GetString("PORT"))
		services.NewImageSevice(grpcServer)

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal(err.Error())
		}
	}()

	router := gin.Default()
	router.Static("/images", "./images")

	go func() {
		log.Printf("Image server started at port: 8009")
		if err := router.Run(":8009"); err != nil {
			log.Fatal(err.Error())
		}
	}()

	select {}
}

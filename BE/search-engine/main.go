package main

import (
	"fmt"
	"log"
	"net"
	"search-engine/config"

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
	log.Printf("GRPC server started at port: %s\n", viper.GetString("port"))

	NewImageSevice(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err.Error())
	}
}

package main

import (
	"bytes"
	pb "common/api"
	"context"
	"io"
	"log"
	"search-engine/database"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type service struct {
	pb.UnimplementedImageServiceServer
	db database.Database
	model_client pb.ModelServiceClient
}

func NewImageSevice(grpcServer *grpc.Server, db database.Database) {
	
	conn, err := grpc.NewClient(viper.GetString("ModelServerAddr"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatal(err)
	}
	model_client := pb.NewModelServiceClient(conn)

	sv := service{db: db, model_client: model_client}
	pb.RegisterImageServiceServer(grpcServer, sv)
}

func (s service) Upload(upload_stream pb.ImageService_UploadServer) error {
	data := bytes.Buffer{}
	size := 0

	model_stream, err := s.model_client.ImageEmbedding(context.Background())
	if err != nil{
		return status.Error(codes.Internal, err.Error())
	}

	req, err := upload_stream.Recv()
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	info := req.GetInfo()
	if info == nil {
		return status.Error(codes.Internal, "The first stream must contains file info")
	}

	for {
		req, err := upload_stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		chunk := req.GetChunk()
		if chunk == nil {
			return status.Error(codes.Internal, "The stream must contains the buffer chunk")
		}

		size += len(chunk)

		_, err = data.Write(chunk)
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		if err := model_stream.Send(&pb.ImageEmbeddingRequest{Chunk: chunk}); err != nil {
			return status.Error(codes.Internal, err.Error())
		}

	}
	if err := upload_stream.SendAndClose(&pb.UploadResponse{Id: "", Size: uint32(size)}); err != nil{
		return status.Error(codes.Internal, err.Error())
	}
	_, err = model_stream.CloseAndRecv()
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

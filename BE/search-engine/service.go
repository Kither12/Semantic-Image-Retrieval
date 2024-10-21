package main

import (
	"bytes"
	pb "common/api"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type service struct {
	pb.UnimplementedImageServiceServer
	model_client pb.ModelServiceClient
}

func NewImageSevice(grpcServer *grpc.Server) {
	
	conn, err := grpc.NewClient(viper.GetString("ModelServerAddr"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatal(err)
	}
	model_client := pb.NewModelServiceClient(conn)

	sv := service{model_client: model_client}
	pb.RegisterImageServiceServer(grpcServer, sv)
}

func (s service) Upload(upload_stream pb.ImageService_UploadServer) error {

	fmt.Println("Receiving upload image request");
	
	data := bytes.Buffer{}
	size := 0

	model_stream, err := s.model_client.ImageEmbedding(context.Background())
	if err != nil{
		fmt.Println(fmt.Sprintf("Failed to open model stream: %v", err))
		return status.Error(codes.Internal, err.Error())
	}


	// The first recv will receive the file info from upload stream
	req, err := upload_stream.Recv()
	if err != nil {
		fmt.Println(fmt.Sprintf("Failed to receive file info: %v", err))
		return status.Error(codes.Internal, err.Error())
	}

	info := req.GetInfo()
	if info == nil {
		fmt.Println("Error: Missing file info");
		return status.Error(codes.Internal, "Missing file info")
	}

	//The following recvs will be the data chunks of the file
	//Then each chunk will be sending to model server to get the embedding vector

	for {
		req, err := upload_stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(fmt.Sprintf("Failed to receive the stream data: %v", err));
			return status.Error(codes.Internal, fmt.Sprintf("Failed to receive the stream data: %v", err))
		}

		chunk := req.GetChunk()
		if chunk == nil {
			fmt.Println(fmt.Sprintf("Chunk data cannot be null"));
			return status.Error(codes.Internal, fmt.Sprintf("Chunk data cannot be null"))
		}
		
		size += len(chunk)
		
		_, err = data.Write(chunk)
		if err != nil {
			fmt.Println(fmt.Sprintf("Failed to write chunk data: %v", err));
			return status.Error(codes.Internal, fmt.Sprintf("Failed to write chunk data: %v", err))
		}
		
		if err := model_stream.Send(&pb.ImageEmbeddingRequest{Chunk: chunk}); err != nil {
			fmt.Println(fmt.Sprintf("Failed to send chunk data to model server: %v", err));
			return status.Error(codes.Internal, fmt.Sprintf("Failed to send chunk data to model server: %v", err))
		}

	}
	fmt.Println("Sucessfully upload the image")

	/* 
		After finish sending all the chunks to the model server, we
		- Store the images the the local disk
		- Store the embedding vector to the vector database
	*/

	imageId := uuid.NewString()

	if err := upload_stream.SendAndClose(&pb.UploadResponse{Id: imageId, Size: uint32(size)}); err != nil{
		fmt.Println(fmt.Sprintf("Failed to sending the response: %v", err));
		return status.Error(codes.Internal, fmt.Sprintf("Failed to sending the response: %v", err))
	}
	_, err = model_stream.CloseAndRecv()
	if err != nil {
		fmt.Println(fmt.Sprintf("Failed to close the model stream: %v", err));
		return status.Error(codes.Internal, fmt.Sprintf("Failed to close the model stream: %v", err))
	}
	fileName := fmt.Sprintf("./images/%s.png", imageId)

	if err := saveBufferToImageFile(&data, fileName); err != nil{
		fmt.Println(fmt.Sprintf("Failed to store file: %v", err));
		return status.Error(codes.Internal, fmt.Sprintf("Failed to store file: %v", err))
	}

	//TODO
	//Store the embedding to vector database




	return nil
}

func (s service) Query(req *pb.SearchRequest, query_stream pb.ImageService_SearchServer) error {
	_, err := s.model_client.TextEmbedding(context.Background(), &pb.TextEmbeddingRequest{Text: req.Prompt});
	if err != nil{
		return status.Error(codes.Internal, err.Error())
	}
	//TODO
	//Use text embedding to query the K-closest vector
	return status.Error(codes.Unimplemented, "Query method is not implemented")
}


func saveBufferToImageFile(buffer *bytes.Buffer, filePath string) error {
    file, err := os.Create(filePath)
    if err != nil {
        return fmt.Errorf("failed to create file: %w", err)
    }
    defer file.Close()

    _, err = file.Write(buffer.Bytes()) 
    if err != nil {
        return fmt.Errorf("failed to write to file: %w", err)
    }

    fmt.Println("File saved successfully:", filePath)
    return nil
}
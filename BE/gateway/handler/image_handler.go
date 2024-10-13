package handler

import (
	pb "common/api"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ImageHandler struct {
	batchSize int
	client    pb.ImageServiceClient
}

func NewImageHandler() *ImageHandler {
	conn, err := grpc.NewClient(viper.GetString("ImageServerAddr"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewImageServiceClient(conn)
	return &ImageHandler{client: client, batchSize: viper.GetInt("image_upload_batch_size")}
}

func (handler *ImageHandler) Upload(ctx *gin.Context) {
	file_header, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	file, err := file_header.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	stream, err := handler.client.Upload(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := stream.Send(&pb.UploadRequest{Data: &pb.UploadRequest_Info{Info: &pb.Info{FileName: file_header.Filename}}}); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	buf := make([]byte, handler.batchSize)
	for {
		num, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		chunk := buf[:num]

		if err := stream.Send(&pb.UploadRequest{Data: &pb.UploadRequest_Chunk{Chunk: chunk}}); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"imageId": resp.Id, "imageSize": resp.Size})

}

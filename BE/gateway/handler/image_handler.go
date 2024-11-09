package handler

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "gateway/proto"
)

type ImageHandler struct {
	batchSize int
	client    pb.ImageServiceClient
}

func NewImageHandler() *ImageHandler {
	conn, err := grpc.NewClient(viper.GetString("IMAGE_SERVER_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewImageServiceClient(conn)
	return &ImageHandler{client: client, batchSize: viper.GetInt("IMAGE_UPLOAD_BATCH_SIZE")}
}

type UploadRequestPayload struct {
	Image       string `json:"image"`
	FileName    string `json:"fileName"`
	ContentType string `json:"contentType"`
}

func (handler *ImageHandler) Upload(ctx *gin.Context) {
	// Parse the JSON payload
	var payload UploadRequestPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Failed to bind JSON data: %v", err)})
		return
	}

	imageData, err := base64.StdEncoding.DecodeString(payload.Image)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Failed to decoded image data: %v", err)})
		return
	}

	// Open the stream to the image server
	stream, err := handler.client.Upload(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to open the uploading stream to image server: %v", err)})
		return
	}

	// Send the file info
	if err := stream.Send(&pb.UploadRequest{Data: &pb.UploadRequest_Info{
		Info: &pb.Info{FileName: payload.FileName},
	}}); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to send file info to image server: %v", err)})
		return
	}

	// Send the image data in chunks
	chunkSize := handler.batchSize
	for i := 0; i < len(imageData); i += chunkSize {
		end := i + chunkSize
		if end > len(imageData) {
			end = len(imageData)
		}
		chunk := imageData[i:end]

		if err := stream.Send(&pb.UploadRequest{
			Data: &pb.UploadRequest_Chunk{Chunk: chunk},
		}); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to send buffer data to image server: %v", err)})
			return
		}
	}

	// Close the stream and receive the response
	resp, err := stream.CloseAndRecv()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to close image server stream: %v", err)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"imageId":   resp.Id,
		"imageSize": resp.Size,
	})
}

func (handler *ImageHandler) Search(ctx *gin.Context) {

	prompt := ctx.DefaultQuery("prompt", "")

	limit, err := strconv.ParseUint(ctx.DefaultQuery("limit", "10"), 10, 64)
	if err != nil || limit == 0 {
		limit = 10
	}

	offset, err := strconv.ParseUint(ctx.DefaultQuery("offset", "0"), 10, 64)
	if err != nil {
		offset = 0
	}

	res, err := handler.client.Search(ctx, &pb.SearchRequest{Prompt: prompt, Limit: limit, Offset: offset})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Failed to send search request to search engine: %v", err)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"images": res.Path,
		"total":  res.Total,
	})
}

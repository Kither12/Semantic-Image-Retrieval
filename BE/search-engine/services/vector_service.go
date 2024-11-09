package services

import (
	"context"
	"fmt"

	"github.com/qdrant/go-client/qdrant"
	"github.com/spf13/viper"
)

type VectorService struct {
	client *qdrant.Client
	nonce  uint64
}

func NewVectorService() (*VectorService, error) {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: viper.GetString("QDRANT_HOST"),
		Port: viper.GetInt("QDRANT_PORT"),
	})

	if err != nil {
		return nil, err
	}

	collection_existed, err := client.CollectionExists(context.Background(), "ImageEmbedding")
	if err != nil {
		return nil, err
	}

	if collection_existed {
		if err := client.DeleteCollection(context.Background(), "ImageEmbedding"); err != nil {
			return nil, err
		}
	}

	if err := client.CreateCollection(context.Background(), &qdrant.CreateCollection{
		CollectionName: "ImageEmbedding",
		VectorsConfig: qdrant.NewVectorsConfig(&qdrant.VectorParams{
			Size:     512,
			Distance: qdrant.Distance_Cosine,
		}),
	}); err != nil {
		return nil, err
	}

	return &VectorService{client: client}, nil
}

func (vs *VectorService) UpsertVector(ctx context.Context, vector []float32, image_path string) error {
	point := &qdrant.UpsertPoints{
		CollectionName: "ImageEmbedding",
		Points: []*qdrant.PointStruct{
			{
				Id:      qdrant.NewIDNum(vs.nonce),
				Vectors: qdrant.NewVectorsDense(vector),
				Payload: qdrant.NewValueMap(map[string]any{"path": image_path}),
			},
		}}
	_, err := vs.client.Upsert(ctx, point)
	if err != nil {
		return fmt.Errorf("failed to upsert vector: %w", err)
	}
	vs.nonce += 1
	return nil
}

func (vs *VectorService) Query(ctx context.Context, vector []float32, limit *uint64, offset *uint64) ([]string, error) {
	query := &qdrant.QueryPoints{
		CollectionName: "ImageEmbedding",
		Query:          qdrant.NewQueryDense(vector),
		Offset:         offset,
		Limit:          limit,
		WithPayload:    qdrant.NewWithPayload(true),
	}

	res, err := vs.client.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query nearest vectors: %w", err)
	}
	var result []string
	for _, point := range res {
		result = append(result, point.Payload["path"].GetStringValue())
	}

	return result, nil
}

func (vs *VectorService) GetNumberImages(ctx context.Context) (uint64, error) {
	return vs.client.Count(ctx, &qdrant.CountPoints{
		CollectionName: "ImageEmbedding",
	})
}

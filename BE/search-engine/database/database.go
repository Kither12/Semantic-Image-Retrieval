package database

import (
	"context"
	"log"

	"github.com/pinecone-io/go-pinecone/pinecone"
	"github.com/spf13/viper"
)
type Database struct{
	index *pinecone.Index
}
func CreateDatabase(ctx context.Context) Database{
    pc, err := pinecone.NewClient(pinecone.NewClientParams{
        ApiKey: viper.GetString("PINECONE_API_KEY"),
    })
	if err != nil{
		log.Fatal(err)
	}
	idx, err := pc.CreateServerlessIndex(ctx, &pinecone.CreateServerlessIndexRequest{
		Name:      "imgstore",
		Dimension: 512,
		Metric:    pinecone.Cosine,
		Cloud:     pinecone.Aws,
		Region:    "us-east-1",
	})
	if err != nil{
		log.Fatal(err)
	}
	return Database{idx}
}
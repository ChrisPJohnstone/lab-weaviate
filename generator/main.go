package main

import (
	"context"

	"github.com/weaviate/weaviate-go-client/v6"
	"github.com/weaviate/weaviate-go-client/v6/collections"
	"github.com/weaviate/weaviate-go-client/v6/collections/vectorindex"
	"github.com/weaviate/weaviate-go-client/v6/data"
	"github.com/weaviate/weaviate-go-client/v6/modules/model2vec"
)

func createCollection(ctx context.Context, client *weaviate.Client) error {
	exists, err := client.Collections.Exists(ctx, collectionName)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	// TODO: Move props to map
	_, err = client.Collections.Create(ctx, collections.Collection{
		Name: collectionName,
		Properties: []collections.Property{
			{Name: "question", DataType: collections.DataTypeText},
			{Name: "answer", DataType: collections.DataTypeText},
		},
		Vectors: map[string]collections.VectorConfig{
			"default": {
				Vectorizer: model2vec.Text2Vec{
					Properties: []string{"question"},
				},
				Index: vectorindex.HNSW{},
			},
		},
	})
	return err
}

func insertRecords(
	ctx context.Context,
	collection *collections.Handle,
	record ...*data.Object,
) error {
	_, err := collection.Data.Insert(ctx, record...)
	return err
}

func main() {
	ctx := context.Background()
	client, err := weaviate.NewLocal(ctx)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	err = createCollection(ctx, client)
	if err != nil {
		panic(err)
	}
	questions := client.Collections.Use(collectionName)
	err = insertRecords(ctx, questions, &data.Object{Properties: map[string]any{
		"question": "What is the meaning of life",
		"answer":   "7",
	}}, &data.Object{Properties: map[string]any{
		"question": "Who has 2 thumbs and can't come up with a second question",
		"answer":   "This guy",
	}})
	if err != nil {
		panic(err)
	}
}

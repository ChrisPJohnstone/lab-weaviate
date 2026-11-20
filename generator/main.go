package main

import (
	"context"

	"github.com/weaviate/weaviate-go-client/v6"
	"github.com/weaviate/weaviate-go-client/v6/collections"
	"github.com/weaviate/weaviate-go-client/v6/data"
)

func createColleciton(ctx context.Context, client *weaviate.Client) error {
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
	})
	return err
}

func main() {
	ctx := context.Background()
	client, err := weaviate.NewLocal(ctx)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	err = createColleciton(ctx, client)
	if err != nil {
		panic(err)
	}
	questions := client.Collections.Use(collectionName)
	_, err = questions.Data.Insert(ctx, &data.Object{Properties: map[string]any{
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

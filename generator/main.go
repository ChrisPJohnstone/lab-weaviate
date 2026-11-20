package main

import (
	"context"
	"fmt"

	"github.com/weaviate/weaviate-go-client/v6"
)

func main() {
	ctx := context.Background()
	client, err := weaviate.NewLocal(ctx)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	fmt.Printf("%s\n", client)
}

package main

import (
	"context"
	"fmt"
	"grpcTest/api"
	"os"

	"google.golang.org/grpc"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "127.0.0.1:4567", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := api.NewDialogClient(conn)

	response, err := client.Ask(ctx, &api.Request{
		Question: os.Args[1],
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(response.GetAnswer())
}

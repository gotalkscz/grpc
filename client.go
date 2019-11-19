package main

import (
	"context"
	"fmt"
	"grpcTest/api"
	"io"
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

	response, err := client.Monologue(ctx, &api.Request{
		Question: os.Args[1],
	})

	if err != nil {
		panic(err)
	}

	for {
		msg, err := response.Recv()
		if err != nil {
			if err == io.EOF {
				return
			}
			panic(err)
		}
		fmt.Println(msg.GetAnswer())
	}
}

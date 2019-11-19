package main

import (
	"context"
	"fmt"
	"grpcTest/api"
	"io"
	"os"
	"strconv"
	"sync"

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

	dialogue, err := client.Dialogue(ctx)

	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			msg, err := dialogue.Recv()
			if err != nil {
				if err == io.EOF {
					return
				}
				panic(err)
			}
			fmt.Println(msg.GetAnswer())
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			if err := dialogue.Send(&api.Request{
				Question: strconv.Itoa(i),
			}); err != nil {
				return
			}
		}
		dialogue.CloseSend()
	}()

	wg.Wait()
}

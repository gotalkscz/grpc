package main

import (
	"context"
	"fmt"
	"grpcTest/api"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		interupt := make(chan os.Signal, 1)
		signal.Notify(interupt, os.Interrupt)
		<-interupt
		cancel()
	}()

	server := grpc.NewServer()
	api.RegisterDialogServer(server, &DialogServer{})

	go func() {
		lis, err := net.Listen("tcp", ":4567")
		if err != nil {
			panic(err)
		}

		defer lis.Close()
		fmt.Println("started")
		if err := server.Serve(lis); err != nil {
			panic(err)
		}
	}()
	<-ctx.Done()
	fmt.Println("stopping")
	server.GracefulStop()
	fmt.Println("stopped")

}

type DialogServer struct {
}

func (DialogServer) Ask(ctx context.Context, in *api.Request) (*api.Response, error) {
	fmt.Println("Q: ", in.GetQuestion())
	return &api.Response{
		Answer: "42",
	}, nil
}

package main

import (
	"context"
	"fmt"
	"grpcTest/api"
	"io"
	"net"
	"os"
	"os/signal"
	"strconv"

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

func (d DialogServer) Dialogue(server api.Dialog_DialogueServer) error {

	for {
		msg, err := server.Recv()
		if err != nil {
			if err == io.EOF {
				fmt.Println("konec diskuze")
				return nil
			}
			fmt.Println("hulvat")
			return nil
		}
		if err := d.Monologue(msg, server); err != nil {
			return err
		}
	}

}

func (DialogServer) Monologue(in *api.Request, server api.Dialog_MonologueServer) error {

	fmt.Println("Q:", in.GetQuestion())

	for i := 0; i < 43; i++ {
		if err := server.Send(&api.Response{
			Answer: in.GetQuestion() + ": " + strconv.Itoa(i),
		}); err != nil {
			if err == io.EOF {
				fmt.Println("hulvat")
				return nil
			}
			return err
		}
	}
	return nil
}

func (DialogServer) Ask(ctx context.Context, in *api.Request) (*api.Response, error) {
	fmt.Println("Q: ", in.GetQuestion())
	return &api.Response{
		Answer: "42",
	}, nil
}

package main

import (
	"context"
	"os"
	"os/signal"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)
		<-interrupt
		cancel()
	}()

	<-ctx.Done()
}

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	fmt.Println("Running worker, waiting for interrupt...")
	<-ctx.Done()
	fmt.Println("Interrupt received, shutting down.")
}

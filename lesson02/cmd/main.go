package main

import (
	"context"
	"fmt"
	"log"
	// "os"
	"os/signal"
	"syscall"

	"github.com/ptsypyshev/gb-golang-backend-level02/lesson02/internal/app"
	// "go.uber.org/zap"
)

func main() {
	fmt.Println("It's ready! But storage and methods are not implemented now :-(")
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	a := app.App{}
	if err := a.Init(ctx); err != nil {
		log.Fatalf("cannot initialize web application: %s", err)
	}
	
	if err := a.Serve(); err != nil {
		log.Fatal(err)
	}

	<-ctx.Done()
	cancel()
}

package main

import (
	"fmt"
	"log"

	"github.com/alexgolang/package-task/internal/app"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	app, err := app.NewApp()
	if err != nil {
		return fmt.Errorf("failed to create app: %w", err)
	}

	return app.Run()
}
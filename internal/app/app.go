package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alexgolang/package-task/internal/app/config"
	service "github.com/alexgolang/package-task/internal/app/service"
	"github.com/alexgolang/package-task/internal/app/transport/httpserver"
	"github.com/alexgolang/package-task/internal/app/transport/httpserver/handlers"
)

type App struct {
	packageService *service.PackageService
	server         httpserver.Server
}

func NewApp() (*App, error) {

	appCfg, err := config.Read()
	if err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	logger := log.New(os.Stdout, "PACKAGE_APP: ", log.Ldate|log.Ltime|log.Lshortfile)

	packageService, err := service.NewPackageService(appCfg.PackageSizes, logger)
	if err != nil {
		return nil, fmt.Errorf("create package service: %w", err)
	}
	handler := handlers.NewPackageHandler(packageService, logger)
	server := httpserver.NewServer(handler, appCfg.HttpServerPort)


	return &App{
		packageService: packageService,
		server:         server,
	}, nil
}

func (a *App) Run() error {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    
    errChan := make(chan error, 1)
    
    shutdown := make(chan os.Signal, 1)
    signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

    go func() {
        fmt.Println("package app starting on port", a.server.Port())
        if err := a.server.Start(); err != nil && err != http.ErrServerClosed {
            errChan <- err
        }
    }()

    select {
    case err := <-errChan:
        return fmt.Errorf("server error: %w", err)
    case sig := <-shutdown:
        fmt.Printf("shutdown signal received: %v\n", sig)
    }

    shutdownCtx, shutdownCancel := context.WithTimeout(ctx, 30*time.Second)
    defer shutdownCancel()

    if err := a.Shutdown(shutdownCtx); err != nil {
        return fmt.Errorf("server shutdown: %w", err)
    }

    fmt.Println("package app gracefully stopped")
    return nil
}

func (a *App) Shutdown(ctx context.Context) error {
	return a.server.Shutdown(ctx)
}

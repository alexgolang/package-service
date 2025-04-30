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
	service "github.com/alexgolang/package-task/internal/app/services"
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

	packageService := service.NewPackageService(appCfg.PackageSizes, logger)
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
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := a.server.Start(); err != nil && err != http.ErrServerClosed {
			log.Printf("server error: %v", err)
			done <- syscall.SIGTERM
		}
	}()

	fmt.Println("package app started")

	<-done

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

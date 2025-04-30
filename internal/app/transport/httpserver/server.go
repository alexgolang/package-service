package httpserver

import (
	"context"
	"fmt"
	"net/http"

	handlers "github.com/alexgolang/package-task/internal/app/transport/httpserver/handlers"
)

type Server interface {
	Start() error
	Shutdown(ctx context.Context) error
}

type httpServer struct {
	srv *http.Server
}

var _ Server = (*httpServer)(nil)

func NewServer(handler *handlers.PackageHandler, port int) Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/package", handler.GetPackage)

	return &httpServer{
		srv: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: mux,
		},
	}
}

func (s *httpServer) Start() error {
	return s.srv.ListenAndServe()
}

func (s *httpServer) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

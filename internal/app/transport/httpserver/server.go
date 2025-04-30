package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	handlers "github.com/alexgolang/package-task/internal/app/transport/httpserver/handlers"
)

type Server interface {
	Start() error
	Shutdown(ctx context.Context) error
	Port() int
}

type httpServer struct {
	srv *http.Server
	port int
}

var _ Server = (*httpServer)(nil)

func NewServer(handler *handlers.PackageHandler, port int) Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/package", handler.GetPackage)

	return &httpServer{
		port: port,
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

func (s *httpServer) Port() int {
    // If port is already stored, return it
    if s.port != 0 {
        return s.port
    }
    
    // Extract port from server address
    addr := s.srv.Addr
    if addr == "" {
        return 8080 // default port
    }
    
    // Remove colon prefix if present
    addr = strings.TrimPrefix(addr, ":")
    
    // Convert to integer
    port, err := strconv.Atoi(addr)
    if err != nil {
        return 8080 // default port on error
    }
    
    s.port = port
    return port
}
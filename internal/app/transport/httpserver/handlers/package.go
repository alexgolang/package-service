package handlers

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/alexgolang/package-task/internal/app/common/server"
	"github.com/alexgolang/package-task/internal/app/service"
)

type Logger interface {
	Println(v ...interface{})
}

type PackageService interface {
	GetPackageSize(ctx context.Context, target int) (map[int]int, error)
}

type PackageHandler struct {
	packageService PackageService
	logger Logger
}

func NewPackageHandler(packageService PackageService, logger Logger) *PackageHandler {
	return &PackageHandler{
		packageService: packageService,
		logger: logger,
	}
}

func (h *PackageHandler) GetPackage(w http.ResponseWriter, r *http.Request) {

	h.logger.Println("received request from", r.RemoteAddr)

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	packageSize := r.URL.Query().Get("size")
	packageSizeInt, err := strconv.Atoi(packageSize)
	if err != nil {
		http.Error(w, "Invalid package size", http.StatusBadRequest)
		return
	}

	result, err := h.packageService.GetPackageSize(r.Context(), packageSizeInt)
	if err != nil {
		if errors.Is(err, service.ErrInvalidSize) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	server.RespondOK(result, w, r)
}

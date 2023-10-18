package http

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type HTTPServer struct {
	// Echo Server Instance
	server *echo.Echo
	// Server Context
	ctx context.Context
}

// NewHTTPServer method
// Create New Server
func NewHTTPServer() *HTTPServer {
	server := echo.New()
	return &HTTPServer{
		server: server,
	}
}

// RunServer HTTP server
func (h *HTTPServer) RunHttpServer(address string) {
	if err := h.server.Start(address); err != nil {
		log.Fatal().Err(err).Msg("")
	}
}

// RunServer HTTP server
func (h *HTTPServer) RunServerWithGraceFull(address string) {
	go h.RunHttpServer(address)

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(h.ctx, 5*time.Second)
	defer cancel()

	if err := h.server.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("")
	}
}

// Shutdown method
// Shutdown The Server by accept the context
func (h *HTTPServer) Shutdown() {
	h.server.Shutdown(h.ctx)
}

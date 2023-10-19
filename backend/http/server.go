package http

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
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
	server.HideBanner = true
	return &HTTPServer{
		server: server,
		ctx:    context.Background(),
	}
}

// RegisterHttp function
// Invoke Routes from http class
func RegisterHttp(services ...HttpRoute) {
	for _, service := range services {
		service.Routes()
	}
}

// RegisterRoute method
func (h *HTTPServer) RegisterRoute(routes *Routes) {
	s := strings.Repeat("=", 10)
	fmt.Printf("%s Register Routes %s \n", s, s)
	for _, route := range routes.Routes {
		r := h.server.Add(route.Method, route.Path, route.Handler, route.Middleware...)
		fmt.Printf("[path:%s] %s [method:%s]\n", r.Path, strings.Repeat("-", 20), r.Method)
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

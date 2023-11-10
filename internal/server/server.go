package server

import (
	"context"
	"net"
	"os"

	"worldofwisdom.com/m/internal/handlers"
	"worldofwisdom.com/m/internal/tcp"
)

// Logger ...
type Logger interface {
	ErrorContext(ctx context.Context, msg string, args ...any)
	WarnContext(ctx context.Context, msg string, args ...any)
	DebugContext(ctx context.Context, msg string, args ...any)
	InfoContext(ctx context.Context, msg string, args ...any)
}

// Handler ...
type Handler interface {
	Handle(ctx context.Context, conn handlers.Conn)
}

// Server ...
type Server struct {
	listener net.Listener
	log      Logger
	timeout  int
}

// NewServer - initialize new tcp server.
//
// 1 argument have to implement Logger, 2 arg = Params specify host, port and timeout. By default 127.0.0.1:8000 and 150 milliseconds
func NewServer(ctx context.Context, logger Logger, params *tcp.Params) *Server {
	listener, err := net.Listen("tcp", params.GetAddress())
	if err != nil {
		logger.ErrorContext(ctx, "error listening", err)
		os.Exit(1)
		return nil
	}

	return &Server{
		listener: listener,
		log:      logger,
		timeout:  params.GetTimeout(),
	}
}

//go:generate  mockgen -source=client.go -destination=mocks/mocks.go
package client

import (
	"context"
	"net"
	"os"

	"worldofwisdom.com/m/internal/services/proof_of_work"
	"worldofwisdom.com/m/internal/tcp"
	"worldofwisdom.com/m/internal/tcp/proto"
)

// Logger ...
type Logger interface {
	ErrorContext(ctx context.Context, msg string, args ...any)
	WarnContext(ctx context.Context, msg string, args ...any)
	DebugContext(ctx context.Context, msg string, args ...any)
	InfoContext(ctx context.Context, msg string, args ...any)
}

// Solver ...
type Solver interface {
	SolveChallenge(ctx context.Context, complexity byte, data []byte) (proof_of_work.Nonce, error)
}

// Conn ...
type Conn interface {
	Send(ctx context.Context, request *proto.SendRequestV1) error
	Read(ctx context.Context) (*proto.Message, error)
	Close(ctx context.Context)
}

// Client ...
type Client struct {
	log     Logger
	solver  Solver
	conn    Conn
	timeout int
}

// NewClient - initialize new tcp client.
//
// 1 argument have to implement Logger, 2 arg = specific host, 3 = specific port, 4 = timeout. By default 127.0.0.1:8000 and 150 milliseconds
func NewClient(ctx context.Context, solver Solver, logger Logger, params *tcp.Params) *Client {
	conn, err := net.Dial("tcp", params.GetAddress())
	if err != nil {
		logger.ErrorContext(ctx, "error listening", err)
		os.Exit(1)
	}

	return &Client{
		log:     logger,
		conn:    tcp.NewConnection(conn, logger),
		timeout: params.GetTimeout(),
		solver:  solver,
	}
}

// Close - method to close tcp connection with error logging
func (c *Client) Close(ctx context.Context) {
	c.conn.Close(ctx)
}

//go:generate  mockgen -source=quotes_handler.go -destination=mocks/mocks.go

package handlers

import (
	"context"

	"worldofwisdom.com/m/internal/tcp/proto"
)

// Logger ...
type Logger interface {
	ErrorContext(ctx context.Context, msg string, args ...any)
	WarnContext(ctx context.Context, msg string, args ...any)
	DebugContext(ctx context.Context, msg string, args ...any)
	InfoContext(ctx context.Context, msg string, args ...any)
}

// QuotesStorage ...
type QuotesStorage interface {
	GetRandomQuote() string
}

// Conn ...
type Conn interface {
	Send(ctx context.Context, request *proto.SendRequestV1) error
	Read(ctx context.Context) (*proto.Message, error)
	Close(ctx context.Context)
}

// POWGenerator ...
type POWGenerator interface {
	Generate(ctx context.Context) ([]byte, byte, error)
	CheckNonce(nonceBytes []byte, data []byte) bool
}

// QuotesHandler ...
type QuotesHandler struct {
	log       Logger
	quotes    QuotesStorage
	generator POWGenerator
}

// NewQuotesHandler ...
func NewQuotesHandler(log Logger, quotes QuotesStorage, generator POWGenerator) *QuotesHandler {
	return &QuotesHandler{log: log, quotes: quotes, generator: generator}
}

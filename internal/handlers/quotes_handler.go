package handlers

import "context"

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

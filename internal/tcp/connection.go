package tcp

import (
	"context"
	"time"
)

// Logger ...
type Logger interface {
	ErrorContext(ctx context.Context, msg string, args ...any)
}

// Conn ...
type Conn interface {
	Read(b []byte) (n int, err error)
	Write(b []byte) (n int, err error)
	Close() error
	SetDeadline(t time.Time) error
}

// Connection ...
type Connection struct {
	conn Conn
	log  Logger
}

// NewConnection ...
func NewConnection(conn Conn, log Logger) *Connection {
	return &Connection{conn: conn, log: log}
}

// Close ...
func (c *Connection) Close(ctx context.Context) {
	err := c.conn.Close()
	if err != nil {
		c.log.ErrorContext(ctx, "close connection error", err)
	}
}

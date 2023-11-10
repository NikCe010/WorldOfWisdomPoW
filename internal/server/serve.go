package server

import (
	"context"
	"time"

	"pow.com/m/cmd/pow/internal/tcp"
)

// Serve - method to serve incoming request using Handler
func (s *Server) Serve(handler Handler) error {
	for {
		ctx := context.Background()

		conn, err := s.listener.Accept()
		if err != nil {
			s.log.ErrorContext(ctx, "Listener accept error", err)
			return err
		}

		err = conn.SetDeadline(time.Now().Add(time.Duration(s.timeout) * time.Millisecond))
		if err != nil {
			s.log.ErrorContext(ctx, "Set deadline error", err)
			return err
		}

		tcpConn := tcp.NewConnection(conn, s.log)
		go handler.Handle(ctx, tcpConn)
	}
}

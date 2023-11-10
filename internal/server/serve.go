package server

import (
	"context"
	"time"

	"worldofwisdom.com/m/internal/tcp"
)

// Serve - method to serve incoming request using Handler
func (s *Server) Serve(handler Handler) error {
	s.wg.Add(1)
	defer s.wg.Done()

	for {
		ctx := context.Background()

		conn, err := s.listener.Accept()
		if err != nil {
			select {
			case <-s.quit:
				return nil
			default:
				s.log.ErrorContext(ctx, "Listener accept error", err)
				continue
			}
		}

		err = conn.SetDeadline(time.Now().Add(time.Duration(s.timeout) * time.Millisecond))
		if err != nil {
			s.log.ErrorContext(ctx, "Set deadline error", err)
		}

		tcpConn := tcp.NewConnection(conn, s.log)
		go handler.Handle(ctx, tcpConn)
	}
}

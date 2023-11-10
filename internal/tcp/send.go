package tcp

import (
	"context"
	"errors"
	"os"

	"worldofwisdom.com/m/internal/tcp/proto"
)

// Send - method for sending data throw protocol.
//
// Convert selected proto.SendRequestV1 to proto.Message. Pack proto.Message into slice of bytes. Send data from tcp.
func (c *Connection) Send(ctx context.Context, request *proto.SendRequestV1) error {
	if request == nil {
		c.log.ErrorContext(ctx, "request is nil")
		return NilRequestErr
	}
	if len(request.Content) > 255 {
		return ContentSizeTooLargeErr
	}

	err := c.send(ctx, request.
		ToMessage().
		ToBytes())
	if err != nil {
		return err
	}

	return nil
}

func (c *Connection) send(ctx context.Context, payload []byte) error {
	_, err := c.conn.Write(payload)
	if err != nil {
		if errors.Is(err, os.ErrDeadlineExceeded) {
			c.log.ErrorContext(ctx, "deadline exceeded")
			return nil
		}
		c.log.ErrorContext(ctx, "tcp client write error", err)
		return err
	}
	return nil
}

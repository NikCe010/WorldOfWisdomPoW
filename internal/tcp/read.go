package tcp

import (
	"context"
	"errors"
	"io"
	"os"

	"worldofwisdom.com/m/internal/tcp/proto"
)

// Read - method for reading data with Message.
//
// Work in two steps. First step - reading content length, Operation, and complexity.
// Second step - reading the exact number of bytes of content.
func (c *Connection) Read(ctx context.Context) (*proto.Message, error) {
	buff := make([]byte, 3)
	err := c.read(ctx, buff) //read params // if less than 3 -> EOF
	if err != nil {
		return nil, err
	}

	operation := proto.Operation(buff[0])
	if operation == proto.Initialize { // initialize does not have any payload
		return &proto.Message{
			Operation: operation,
		}, nil
	}

	length := buff[2]
	data := make([]byte, length)
	err = c.read(ctx, data) //read payload
	if err != nil {
		return nil, err
	}

	return &proto.Message{
		Operation:  operation,
		Complexity: buff[1],
		Length:     length,
		Content:    data,
	}, nil
}

func (c *Connection) read(ctx context.Context, buff []byte) error {
	_, err := io.ReadAtLeast(c.conn, buff, len(buff))
	if err != nil {
		if errors.Is(err, os.ErrDeadlineExceeded) {
			c.log.ErrorContext(ctx, "deadline exceeded")
			return err
		}
		if errors.Is(err, io.EOF) {
			c.log.ErrorContext(ctx, "number of read parameters is less than set")
			return err
		}
		c.log.ErrorContext(ctx, "read request error", err)
		return err
	}
	return nil
}

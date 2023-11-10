package client

import (
	"context"
	"time"

	"worldofwisdom.com/m/internal/tcp"
	"worldofwisdom.com/m/internal/tcp/proto"
)

// GetQuote - method to request random "World Of Wisdom" quote from server
//
// Has 5 stages: Initiate communication, read challenge, solve challenge, send nonce and read quote.
func (c *Client) GetQuote(ctx context.Context) ([]byte, error) {
	c.log.InfoContext(ctx, "Stage 1: Initiate communication")
	err := c.conn.Send(ctx, proto.NewInitiateRequest())
	if err != nil {
		c.log.ErrorContext(ctx, "Quote request error", err)
		return nil, err
	}

	c.log.InfoContext(ctx, "Stage 2: Read challenge")
	challengeResp, err := c.conn.Read(ctx)
	if err != nil {
		c.log.ErrorContext(ctx, "Read challenge error", err)
		return nil, err
	}
	if challengeResp.Operation != proto.SendChallenge {
		c.log.ErrorContext(ctx, "invalid protocol operation")
		return nil, tcp.InvalidProtocolErr
	}

	c.log.InfoContext(ctx, "Stage 3: Solve challenge")
	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Millisecond*time.Duration(c.timeout))
	defer cancel()
	nonce, err := c.solver.SolveChallenge(ctxWithTimeout, challengeResp.Complexity, challengeResp.Content)
	if err != nil {
		c.log.ErrorContext(ctx, "Solve challenge error", err)
		return nil, err
	}

	c.log.InfoContext(ctx, "Stage 4: Send solved challenge")
	err = c.conn.Send(ctx, proto.NewSolvedChallengeRequest(nonce.ToBytes()))
	if err != nil {
		c.log.ErrorContext(ctx, "Send solved challenge error", err)
		return nil, err
	}

	c.log.InfoContext(ctx, "Stage 5: Read quote")
	quoteResp, err := c.conn.Read(ctx)
	if err != nil {
		c.log.ErrorContext(ctx, "Read quote error", err)
		return nil, err
	}
	if quoteResp.Operation != proto.SendData {
		c.log.ErrorContext(ctx, "invalid protocol operation")
		return nil, tcp.InvalidProtocolErr
	}

	return quoteResp.Content, nil
}

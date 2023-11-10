package handlers

import (
	"context"

	"pow.com/m/cmd/pow/internal/server"
	"pow.com/m/cmd/pow/internal/tcp/proto"
)

// Handle ...
func (h *QuotesHandler) Handle(ctx context.Context, conn server.Conn) {
	defer conn.Close(ctx)

	h.log.InfoContext(ctx, "Stage 1: Read init request")
	initResp, err := conn.Read(ctx)
	if err != nil {
		h.log.ErrorContext(ctx, "read request error", err)
		return
	}
	if initResp.Operation != proto.Initialize {
		h.log.ErrorContext(ctx, "invalid protocol operation")
		return
	}

	h.log.InfoContext(ctx, "Stage 2: Generate challenge")
	challenge, complexity, err := h.generator.Generate(ctx)
	if err != nil {
		h.log.ErrorContext(ctx, "generate challenge error", err)
		return
	}

	h.log.InfoContext(ctx, "Stage 3: Send challenge")
	err = conn.Send(ctx, proto.NewSendChallengeRequest(complexity, challenge))
	if err != nil {
		h.log.ErrorContext(ctx, "write response error", err)
		return
	}

	h.log.InfoContext(ctx, "Stage 4: Read nonce")
	resp, err := conn.Read(ctx)
	if err != nil {
		h.log.ErrorContext(ctx, "read request error", err)
		return
	}
	if resp.Operation != proto.SendNonce {
		h.log.ErrorContext(ctx, "invalid protocol operation")
		return
	}

	h.log.InfoContext(ctx, "Stage 5: Check nonce")
	if h.generator.CheckNonce(resp.Content, challenge) {
		h.log.InfoContext(ctx, "client resolve challenge")
	} else {
		h.log.InfoContext(ctx, "client was unable to solve the challenge")
		return
	}

	h.log.InfoContext(ctx, "Stage 6: Get random quote")
	quote := h.quotes.GetRandomQuote()

	h.log.InfoContext(ctx, "Stage 7: Send quote")
	err = conn.Send(ctx, proto.NewSendDataRequest([]byte(quote)))
	if err != nil {
		h.log.ErrorContext(ctx, "write response error", err)
		return
	}
}

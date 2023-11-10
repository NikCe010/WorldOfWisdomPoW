package main

import (
	"context"
	"log/slog"
	"os"

	"pow.com/m/cmd/pow/internal/handlers"
	"pow.com/m/cmd/pow/internal/server"
	"pow.com/m/cmd/pow/internal/services/proof_of_work"
	"pow.com/m/cmd/pow/internal/services/quotes_service"
	"pow.com/m/cmd/pow/internal/tcp"
)

func main() {
	ctx := context.Background()
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Init services
	quotesService := quotes_service.New()
	generator := proof_of_work.NewGenerator(2)

	// Init server
	tcpServer := server.NewServer(ctx, log,
		tcp.NewParams().
			SetTimeout(200))

	// Init handlers
	quotesHandler := handlers.NewQuotesHandler(log, quotesService, generator)

	err := tcpServer.Serve(quotesHandler)
	if err != nil {
		return
	}
}

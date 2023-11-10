package main

import (
	"context"
	"log/slog"
	"os"

	"worldofwisdom.com/m/internal/handlers"
	"worldofwisdom.com/m/internal/server"
	"worldofwisdom.com/m/internal/services/proof_of_work"
	"worldofwisdom.com/m/internal/services/quotes_service"
	"worldofwisdom.com/m/internal/tcp"
)

func main() {
	ctx := context.Background()
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Init services
	quotesService := quotes_service.New()
	generator := proof_of_work.NewGenerator(2)

	params := tcp.NewParams(200)
	// Init server
	tcpServer := server.NewServer(ctx, log, params)

	// Init handlers
	quotesHandler := handlers.NewQuotesHandler(log, quotesService, generator)

	err := tcpServer.Serve(quotesHandler)
	if err != nil {
		return
	}
}

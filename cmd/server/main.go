package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"worldofwisdom.com/m/internal/handlers"
	"worldofwisdom.com/m/internal/server"
	"worldofwisdom.com/m/internal/services/proof_of_work"
	"worldofwisdom.com/m/internal/services/quotes_service"
	"worldofwisdom.com/m/internal/tcp"
)

func main() {
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Init services
	quotesService := quotes_service.New()
	generator := proof_of_work.NewGenerator(2)

	// Init server
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	params := tcp.NewParams(150)
	tcpServer := server.NewServer(ctx, log, params)
	defer tcpServer.Stop()

	// Init handlers
	quotesHandler := handlers.NewQuotesHandler(log, quotesService, generator)

	err := tcpServer.Serve(quotesHandler)
	if err != nil {
		return
	}
}

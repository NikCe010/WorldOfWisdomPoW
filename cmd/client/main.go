package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"worldofwisdom.com/m/internal/client"
	"worldofwisdom.com/m/internal/services/proof_of_work"
	"worldofwisdom.com/m/internal/tcp"
)

func main() {
	ctx := context.Background()
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))

	solver := proof_of_work.NewSolver()

	params := tcp.NewParams(150)
	client := client.NewClient(ctx, solver, log, params)

	defer client.Close(ctx)
	start := time.Now()

	quote, err := client.GetQuote(ctx)
	if err != nil {
		return
	}

	log.Info("Working time", time.Now().Sub(start).String())
	log.Info(string(quote))
}

package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"pow.com/m/cmd/pow/internal/client"
	"pow.com/m/cmd/pow/internal/services/proof_of_work"
)

func main() {
	ctx := context.Background()
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))

	solver := proof_of_work.NewSolver()
	client := client.NewClient(ctx, solver, log, nil)

	defer client.Close(ctx)
	start := time.Now()

	quote, err := client.GetQuote(ctx)
	if err != nil {
		return
	}

	log.Debug("Working time", time.Now().Sub(start).String())
	log.Info(string(quote))
}

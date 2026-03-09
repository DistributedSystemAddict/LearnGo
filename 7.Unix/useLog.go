package main

import (
	"fmt"
	"log/slog"
	"os"
)

func main14() {
	slog.Error("This is an ERROR message")
	slog.Debug("This is a DEBUG message")
	slog.Info("This ia an INFO message")
	slog.Warn("This is a WARNING message")

	logLevel := &slog.LevelVar{}
	fmt.Println("Log level:", logLevel)

	opts := &slog.HandlerOptions{
		Level: logLevel,
	}
	handler := slog.NewTextHandler(os.Stdout, opts)
	logger := slog.New(handler)
	logLevel.Set(slog.LevelInfo)
	logger.Debug("This is a DEBUG message")
	logger.Info("This ia an INFO message")

	logJSON := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logJSON.Error("ERROR message in JSON")
}

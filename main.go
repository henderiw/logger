package main

import (
	"errors"
	"log/slog"

	"github.com/henderiw/logger/log"
)

func main() {
	var l *slog.Logger
	l = slog.
		New(log.NewHandler(&log.HandlerOptions{
			Name:      "wirer-controller",
			AddSource: false,
		})).
		With(slog.Group(
			"gvknsn",
			"kind", "Node",
			"name", "high",
			"mood", "amazing",
		)).
		WithGroup("data")
	l.Error("Ops, something went wrong!",
		"error", errors.New("bad luck my friend"),
		"key1", "value1",
	)
	l.Info("This is pretty, right?",
		"count", 10,
		"pi", 3.14,
		"happy", true,
		"nothing", nil,
	)

	l = log.NewLogger(&log.HandlerOptions{Name: "wire-controller", AddSource: false})
	l = l.WithGroup("group").With("k11", "v11", "k12", "v12")
	l.Info("test",
		"k1", "v1")

	l.Info("")
}

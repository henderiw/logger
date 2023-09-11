package log

import (
	"context"
	"fmt"
	"log/slog"
)

type Handler struct {
	h slog.Handler
	//r func([]string, slog.Attr) slog.Attr
	//b *bytes.Buffer // capture the output from the "nested" handler
	//m *sync.Mutex   // guarantee thread safe access to *bytes.Buffer
}

func (h *Handler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.h.Enabled(ctx, level)
}

func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &Handler{h: h.h.WithAttrs(attrs)}
}

func (h *Handler) WithGroup(name string) slog.Handler {
	return &Handler{h: h.h.WithGroup(name)}
}

func (h *Handler) Handle(ctx context.Context, r slog.Record) error {
	/*
		if r.Level >= slog.LevelError {
			r.AddAttrs(slog.String(attrErrorTypeKey, attrErrorTypeVal))
		}
	*/
	err := h.h.Handle(ctx, r)
	if err != nil {
		return fmt.Errorf("error when calling nested handler's Handle: %w", err)
	}
	return nil
}

package log

import (
	"os"

	"log/slog"
)

type HandlerOptions struct {
	Name      string
	MinLevel  slog.Leveler
	AddSource bool
}

func logAttrs(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.MessageKey {
		a.Key = "message"
		return a
	}
	if a.Key == slog.SourceKey {
		a.Key = "src"
		return a
	}
	if err, ok := a.Value.Any().(error); ok {
		return slog.Group("error",
			slog.String("message", err.Error()),
			slog.Any("stack", CaptureStack().Slice()),
		)
	}
	return a
}

func NewHandler(opts *HandlerOptions) *Handler {
	handlerOpts := &slog.HandlerOptions{
		Level:       opts.MinLevel,
		AddSource:   opts.AddSource,
		ReplaceAttr: logAttrs,
	}
	handler := slog.
		NewJSONHandler(os.Stdout, handlerOpts).
		WithAttrs([]slog.Attr{
			slog.String("logger",opts.Name),	
		})
	return &Handler{h: handler}
}

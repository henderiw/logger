package log

import "log/slog"


func NewLogger(opts *HandlerOptions) *slog.Logger {
	return slog.
		New(NewHandler(opts)).
		WithGroup("data")
}

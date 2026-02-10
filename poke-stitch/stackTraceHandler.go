package main

import (
	"context"
	"log/slog"
	"runtime"
)

type StackTraceHandler struct {
	slog.Handler
}

func (h *StackTraceHandler) Handle(ctx context.Context, r slog.Record) error {
	// Capture a stack trace for the current goroutine.
	buf := make([]byte, 64<<10)
	n := runtime.Stack(buf, false)
	r.Add("stack", string(buf[:n]))
	return h.Handler.Handle(ctx, r)
}

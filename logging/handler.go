package logging

import (
	"context"
	"go-playground/utility"
	"log/slog"
	"os"

	"go.opentelemetry.io/otel/trace"
)

const (
	threadIdKey       string = "threadId"
	traceIdKey        string = "traceId"
	logSpecVersionKey string = "logSpecVersion"
)

type UnifiedLogHandler struct {
	options *handlerOptions
	handler slog.Handler
}

func (h *UnifiedLogHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

func (h *UnifiedLogHandler) Handle(ctx context.Context, record slog.Record) error {
	appendNonBuiltIns(ctx, &record)
	return h.handler.Handle(ctx, record)
}

func appendNonBuiltIns(ctx context.Context, record *slog.Record) {
	record.AddAttrs(slog.String(threadIdKey, utility.CurrentRoutineName()))

	traceId := ""
	span := trace.SpanFromContext(ctx)
	if span != nil {
		traceId = span.SpanContext().TraceID().String()
	}
	record.AddAttrs(slog.String(traceIdKey, traceId))
}

func (h *UnifiedLogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &UnifiedLogHandler{
		options: h.options,
		handler: h.handler.WithAttrs(attrs),
	}
}

func (h *UnifiedLogHandler) WithGroup(name string) slog.Handler {
	return &UnifiedLogHandler{
		options: h.options,
		handler: h.handler.WithGroup(name),
	}
}

func NewUnifiedLogHandler(options ...HandlerOption) *UnifiedLogHandler {
	opts := &handlerOptions{}
	for _, opt := range options {
		opt(opts)
	}
	if opts.Level == nil {
		opts.Level = new(slog.LevelVar)
	}
	if opts.Writer == nil {
		opts.Writer = new(atomicWriter)
		opts.Writer.Set(newWriterHolder(os.Stdout))
	}

	attrs := []slog.Attr{slog.String(logSpecVersionKey, opts.LogSpecVersion)}
	handler := slog.NewJSONHandler(opts.Writer, &slog.HandlerOptions{
		AddSource:   true,
		Level:       opts.Level,
		ReplaceAttr: newCompositeAttrReplacer(opts).replace,
	}).WithAttrs(attrs)

	return &UnifiedLogHandler{
		options: opts,
		handler: handler,
	}
}

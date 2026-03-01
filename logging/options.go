package logging

import (
	"io"
	"log/slog"
)

type HandlerOption func(options *handlerOptions)

type handlerOptions struct {
	TimeKey           string
	TimeFormat        string
	LevelKey          string
	SourceKey         string
	MessageKey        string
	ThreadIdKey       string
	TraceIdKey        string
	LogSpecVersionKey string
	LogSpecVersion    string
	Level             *slog.LevelVar
	Writer            *atomicWriter
}

func WithTimeKey(key string) HandlerOption {
	return func(options *handlerOptions) {
		options.TimeKey = key
	}
}

func WithTimeFormat(format string) HandlerOption {
	return func(options *handlerOptions) {
		options.TimeFormat = format
	}
}

func WithLevelKey(key string) HandlerOption {
	return func(options *handlerOptions) {
		options.LevelKey = key
	}
}

func WithSourceKey(key string) HandlerOption {
	return func(options *handlerOptions) {
		options.SourceKey = key
	}
}

func WithMessageKey(key string) HandlerOption {
	return func(options *handlerOptions) {
		options.MessageKey = key
	}
}

func WithThreadIdKey(key string) HandlerOption {
	return func(options *handlerOptions) {
		options.ThreadIdKey = key
	}
}

func WithTraceIdKey(key string) HandlerOption {
	return func(options *handlerOptions) {
		options.TraceIdKey = key
	}
}

func WithLogSpecVersionKey(key string) HandlerOption {
	return func(options *handlerOptions) {
		options.LogSpecVersionKey = key
	}
}

func WithLogSpecVersion(version string) HandlerOption {
	return func(options *handlerOptions) {
		options.LogSpecVersion = version
	}
}

func WithLevel(level slog.Level) HandlerOption {
	return func(options *handlerOptions) {
		if options.Level == nil {
			options.Level = new(slog.LevelVar)
		}
		options.Level.Set(level)
	}
}

func WithWriter(writer io.Writer) HandlerOption {
	return func(options *handlerOptions) {
		if options.Writer == nil {
			options.Writer = new(atomicWriter)
		}
		options.Writer.Set(newWriterHolder(writer))
	}
}

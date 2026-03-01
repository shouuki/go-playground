package logging

import (
	"log/slog"
	"sync"
)

var (
	mutex         = sync.Mutex{}
	loggers       = make(map[string]*slog.Logger)
	options       = make(map[string]*handlerOptions)
	defaultConfig = &Config{
		Options: make([]HandlerOption, 0),
		Level:   make(map[string]slog.Level, 0),
	}
)

type Config struct {
	Options []HandlerOption
	Level   map[string]slog.Level
}

func GetLogger(name string) *slog.Logger {
	mutex.Lock()
	defer mutex.Unlock()
	if logger, ok := loggers[name]; ok {
		return logger
	}
	handler := NewUnifiedLogHandler(defaultConfig.Options...)
	if level, ok := defaultConfig.Level[name]; ok {
		handler.options.Level.Set(level)
	}
	logger := slog.New(handler)
	loggers[name] = logger
	options[name] = handler.options
	return logger
}

func UpdateConfig(config *Config) {
	if config == nil {
		return
	}
	mutex.Lock()
	defer mutex.Unlock()
	defaultConfig = config
	refreshLoggers()
}

func refreshLoggers() {
	for _, handlerOptions := range options {
		for _, opt := range defaultConfig.Options {
			opt(handlerOptions)
		}
	}
	for name, level := range defaultConfig.Level {
		if handlerOptions, ok := options[name]; ok {
			handlerOptions.Level.Set(level)
		}
	}
}

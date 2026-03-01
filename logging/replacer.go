package logging

import (
	"fmt"
	"log/slog"
	"time"
)

type attrReplacer interface {
	replace(groups []string, attr slog.Attr) slog.Attr
}

type timeAttrReplacer struct {
	options *handlerOptions
}

func (r *timeAttrReplacer) replace(groups []string, attr slog.Attr) slog.Attr {
	_ = groups
	if attr.Key == slog.TimeKey && attr.Value.Kind() == slog.KindTime {
		key := r.options.TimeKey
		if key == "" {
			key = slog.TimeKey
		}
		format := r.options.TimeFormat
		if format == "" {
			format = time.RFC3339Nano
		}
		return slog.String(key, attr.Value.Time().Format(format))
	}
	return attr
}

type levelAttrReplacer struct {
	options *handlerOptions
}

func (r *levelAttrReplacer) replace(groups []string, attr slog.Attr) slog.Attr {
	_ = groups
	if attr.Key == slog.LevelKey && r.options.LevelKey != "" {
		return slog.Any(r.options.LevelKey, attr.Value)
	}
	return attr
}

type sourceAttrReplacer struct {
	options *handlerOptions
}

func (r *sourceAttrReplacer) replace(groups []string, attr slog.Attr) slog.Attr {
	_ = groups
	if attr.Key == slog.SourceKey {
		if src, ok := attr.Value.Any().(*slog.Source); ok {
			key := r.options.SourceKey
			if key == "" {
				key = slog.SourceKey
			}
			return slog.String(key, fmt.Sprintf("%v [%v:%v]", src.Function, src.File, src.Line))
		}
	}
	return attr
}

type messageAttrReplacer struct {
	options *handlerOptions
}

func (r *messageAttrReplacer) replace(groups []string, attr slog.Attr) slog.Attr {
	_ = groups
	if attr.Key == slog.MessageKey && r.options.MessageKey != "" {
		return slog.Any(r.options.MessageKey, attr.Value)
	}
	return attr
}

type threadIdAttrReplacer struct {
	options *handlerOptions
}

func (r *threadIdAttrReplacer) replace(groups []string, attr slog.Attr) slog.Attr {
	_ = groups
	if attr.Key == threadIdKey && r.options.ThreadIdKey != "" {
		return slog.Any(r.options.ThreadIdKey, attr.Value)
	}
	return attr
}

type traceIdAttrReplacer struct {
	options *handlerOptions
}

func (r *traceIdAttrReplacer) replace(groups []string, attr slog.Attr) slog.Attr {
	_ = groups
	if attr.Key == traceIdKey && r.options.TraceIdKey != "" {
		return slog.Any(r.options.TraceIdKey, attr.Value)
	}
	return attr
}

type logSpecVersionAttrReplacer struct {
	options *handlerOptions
}

func (r *logSpecVersionAttrReplacer) replace(groups []string, attr slog.Attr) slog.Attr {
	_ = groups
	if attr.Key == logSpecVersionKey && r.options.LogSpecVersionKey != "" {
		return slog.Any(r.options.LogSpecVersionKey, attr.Value)
	}
	return attr
}

type compositeAttrReplace struct {
	replacers []attrReplacer
}

func (c *compositeAttrReplace) replace(groups []string, attr slog.Attr) slog.Attr {
	if len(c.replacers) == 0 {
		return attr
	}
	for _, r := range c.replacers {
		attr = r.replace(groups, attr)
	}
	return attr
}

func newCompositeAttrReplacer(options *handlerOptions) attrReplacer {
	return &compositeAttrReplace{
		replacers: []attrReplacer{
			&timeAttrReplacer{options: options},
			&levelAttrReplacer{options: options},
			&sourceAttrReplacer{options: options},
			&messageAttrReplacer{options: options},
			&threadIdAttrReplacer{options: options},
			&traceIdAttrReplacer{options: options},
			&logSpecVersionAttrReplacer{options: options},
		},
	}
}

package utility

import (
	"encoding/json"
	"net/http"
)

type jsonOptions struct {
	UseNumber             bool
	DisallowUnknownFields bool
}

type JsonOption func(*jsonOptions)

func WithUseNumber() JsonOption {
	return func(opts *jsonOptions) {
		opts.UseNumber = true
	}
}

func WithDisallowUnknownFields() JsonOption {
	return func(opts *jsonOptions) {
		opts.DisallowUnknownFields = true
	}
}

// BindJson binds the JSON body of the request to the given value pointed by v.
func BindJson(r *http.Request, v any, opts ...JsonOption) error {
	options := &jsonOptions{}
	for _, opt := range opts {
		opt(options)
	}
	decoder := json.NewDecoder(r.Body)
	if options.UseNumber {
		decoder.UseNumber()
	}
	if options.DisallowUnknownFields {
		decoder.DisallowUnknownFields()
	}
	return decoder.Decode(v)
}

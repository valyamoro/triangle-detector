package detect

import "github.com/gopherchan2006/go-triangle-detector/internal/detect/spec"

type Option func(*opts)

type opts struct {
	params  spec.Params
	counter spec.RejectCounter
}

func newOpts(options []Option) opts {
	o := opts{
		params:  spec.DefaultParams(),
		counter: NoopCounter{},
	}
	for _, opt := range options {
		opt(&o)
	}
	return o
}

func WithTrace(_ bool) Option {
	return func(o *opts) {}
}

func WithParams(p Params) Option {
	return func(o *opts) { o.params = p }
}

func WithCounter(c RejectCounter) Option {
	return func(o *opts) { o.counter = c }
}

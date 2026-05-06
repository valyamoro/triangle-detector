package detect

import (
	"github.com/gopherchan2006/go-triangle-detector/internal/detect/engine"
	"github.com/gopherchan2006/go-triangle-detector/internal/domain"
)

func DetectAscendingTriangle(candles []domain.Candle, options ...Option) Result {
	o := newOpts(options)
	return engine.Detect(candles, engine.RunOpts{
		Params:  o.params,
		Counter: o.counter,
	})
}

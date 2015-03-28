package dptest

import (
	"sync"

	"github.com/signalfx/metricproxy/datapoint"
	"golang.org/x/net/context"
)

// BasicSink is a pure testing sink that blocks forwarded points onto a channel
type BasicSink struct {
	RetErr     error
	PointsChan chan []*datapoint.Datapoint

	mu sync.Mutex
}

// Next returns a single datapoint from the top of PointsChan and panics if the top doesn't contain
// only one point
func (f *BasicSink) Next() *datapoint.Datapoint {
	r := <-f.PointsChan
	if len(r) != 1 {
		panic("Expect a single point")
	}
	return r[0]
}

// AddDatapoints buffers the point on an internal chan or returns errors if RetErr is set
func (f *BasicSink) AddDatapoints(ctx context.Context, points []*datapoint.Datapoint) error {
	f.mu.Lock()
	if f.RetErr != nil {
		defer f.mu.Unlock()
		return f.RetErr
	}
	f.mu.Unlock()
	select {
	case f.PointsChan <- points:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RetError sets an error that is returned on AddDatapoints calls
func (f *BasicSink) RetError(err error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.RetErr = err
}

// NewBasicSink creates a BasicSink with an unbuffered chan.  Note, calls to AddDatapoints will then
// block until you drain the PointsChan.
func NewBasicSink() *BasicSink {
	return &BasicSink{
		PointsChan: make(chan []*datapoint.Datapoint),
	}
}

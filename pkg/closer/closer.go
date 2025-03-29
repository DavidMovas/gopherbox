package closer

import (
	"context"
	"errors"
	"io"
)

type ICloser interface {
	Close(ctx context.Context) error
}

var _ ICloser = (*Closer)(nil)

type closeFn func() error
type closeCtxFn func(ctx context.Context) error

// Closer helper structure that store entities that can be closed at once by invoke Close function
type Closer struct {
	closers    []closeFn
	closersCtx []closeCtxFn
}

// NewCloser returns new Closer
func NewCloser(closers ...closeFn) *Closer {
	return &Closer{
		closers: closers,
	}
}

// Push lets add to store close function
func (c *Closer) Push(closer closeFn) {
	c.closers = append(c.closers, closer)
}

// PushIO lets add to store entity that implements io.Closer
func (c *Closer) PushIO(closer io.Closer) {
	c.closers = append(c.closers, closer.Close)
}

// PushCtx lets add to store close function that demand context to be invoked
func (c *Closer) PushCtx(closer closeCtxFn) {
	c.closersCtx = append(c.closersCtx, closer)
}

// PushNE lets add to store close function that not returns error after invoking
func (c *Closer) PushNE(closer func()) {
	c.closers = append(c.closers, func() error {
		closer()
		return nil
	})
}

// Close function implementation of ICloser that close all stored closers in reverse order,
// demand context for close function that needed context to be invoked
func (c *Closer) Close(ctx context.Context) error {
	var err error

	for i := len(c.closers) - 1; i >= 0; i-- {
		err = errors.Join(err, c.closers[i]())
	}

	for i := len(c.closersCtx) - 1; i >= 0; i-- {
		err = errors.Join(err, c.closersCtx[i](ctx))
	}

	return err
}

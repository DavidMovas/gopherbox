package ctcx

import (
	"context"
	"time"
)

const shortContextTimeout = time.Second * 5

// ShortCtx returns context and cancel it after short time
//
// Timeout: 5s
func ShortCtx() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), shortContextTimeout)

	time.AfterFunc(shortContextTimeout, func() {
		cancel()
	})

	return ctx
}

// ShortCtxParent returns context delivered from parent context and cancel it after short time
//
// Timeout: 5s
func ShortCtxParent(ctx context.Context) context.Context {
	ctx, cancel := context.WithTimeout(ctx, shortContextTimeout)

	time.AfterFunc(shortContextTimeout, func() {
		cancel()
	})

	return ctx
}

// TimeoutCtx returns context with target timeout and cancel it after time
func TimeoutCtx(timeout time.Duration) context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	time.AfterFunc(timeout, func() {
		cancel()
	})

	return ctx
}

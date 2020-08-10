package identity

import (
	"github.com/google/uuid"
	"github.com/knowlet/atreugo-bootstrapper/bootstrap"
	"github.com/savsgio/atreugo/v11"
)

// New returns the middleware to adds an identifier to the request
// Header name: X-Request-ID
// Header value: uuid4()
//
// It's recomemded to add this middleware at first and before the view execution.
func New() atreugo.Middleware {
	return func(ctx *atreugo.RequestCtx) error {
		if ctx.Request.Header.Peek(atreugo.XRequestIDHeader) == nil {
			ctx.Request.Header.Set(atreugo.XRequestIDHeader, uuid.New().String())
		}

		return ctx.Next()
	}
}

// Configure creates a new identity middleware and registers that to the app.
func Configure(b *bootstrap.Bootstrapper) {
	h := New()
	b.UseBefore(h)
}

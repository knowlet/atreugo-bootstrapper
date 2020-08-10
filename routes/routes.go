package routes

import (
	"github.com/knowlet/atreugo-bootstrapper/bootstrap"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {
	b.GET("/", GetIndexHandler)

	b.GET("/echo/{path:*}", GetEchoHandler)

	v1 := b.NewGroupPath("/api")
	v1.GET("/", GetApiIndexHandler)
}

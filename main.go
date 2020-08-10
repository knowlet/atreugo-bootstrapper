package main

import (
	"github.com/knowlet/atreugo-bootstrapper/bootstrap"
	"github.com/knowlet/atreugo-bootstrapper/middleware/identity"
	"github.com/knowlet/atreugo-bootstrapper/routes"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("Awesome App")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen("localhost:9527")
}

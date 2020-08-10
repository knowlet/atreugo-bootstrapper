package bootstrap

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/savsgio/atreugo/v11"
)

type Configurator func(*Bootstrapper)

type Bootstrapper struct {
	*atreugo.Atreugo
	AppSpawnDate time.Time
}

// New returns a new Bootstrapper.
func New(appName string, cfgs ...Configurator) *Bootstrapper {
	b := &Bootstrapper{
		AppSpawnDate: time.Now(),
		Atreugo: atreugo.New(atreugo.Config{
			Name: appName,
			NotFoundView: func(ctx *atreugo.RequestCtx) error {
				return ctx.JSONResponse(atreugo.JSON{
					"status":  http.StatusNotFound,
					"message": "",
				}, http.StatusNotFound)
			},
			MethodNotAllowedView: func(ctx *atreugo.RequestCtx) error {
				return ctx.JSONResponse(atreugo.JSON{
					"status":  http.StatusMethodNotAllowed,
					"message": "",
				}, http.StatusMethodNotAllowed)
			},
			ErrorView: func(ctx *atreugo.RequestCtx, err error, statusCode int) {
				if err := ctx.JSONResponse(atreugo.JSON{
					"status":  statusCode,
					"message": err,
				}); err != nil {
					panic(err)
				}
			},
			PanicView: func(ctx *atreugo.RequestCtx, msg interface{}) {
				if err := ctx.JSONResponse(atreugo.JSON{
					"status":  http.StatusInternalServerError,
					"message": "",
				}); err != nil {
					panic(err)
				}
			},
		}),
	}

	for _, cfg := range cfgs {
		cfg(b)
	}

	return b
}

const (
	// StaticAssets is the root directory for public assets like images, css, js.
	StaticAssets = "./public"
	// Favicon is the relative 9to the "StaticAssets") favicon path for our app.
	Favicon = "/favicon.ico"
)

// Configure accepts configurations and runs them inside the Bootstraper's context.
func (b *Bootstrapper) Configure(cfgs ...Configurator) {
	for _, cfg := range cfgs {
		cfg(b)
	}
}

// Bootstrap prepares our application.
//
// Returns itself.
func (b *Bootstrapper) Bootstrap() *Bootstrapper {

	// b.SetupErrorHandlers()

	// static files
	b.Static(Favicon, StaticAssets+Favicon)
	b.StaticCustom("/", &atreugo.StaticFS{
		Root:               StaticAssets,
		GenerateIndexPages: false,
		AcceptByteRange:    true,
		Compress:           true,
		PathNotFound: func(ctx *atreugo.RequestCtx) error {
			return ctx.JSONResponse(atreugo.JSON{
				"status":  http.StatusNotFound,
				"message": "",
			})
		},
	})

	// middleware, after static files

	return b
}

// Listen starts the http server with the specified "addr".
func (b *Bootstrapper) Listen(addr string) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	if err := b.Serve(ln); err != nil {
		panic(err)
	}
}

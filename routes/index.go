package routes

import "github.com/savsgio/atreugo/v11"

// GetIndexHandler handles the GET: /
func GetIndexHandler(ctx *atreugo.RequestCtx) error {
	return ctx.TextResponse("Hello World")
}

// GetEchoHandler handles the GET: /echo
func GetEchoHandler(ctx *atreugo.RequestCtx) error {
	return ctx.TextResponse("Echo message: " + ctx.UserValue("path").(string))
}

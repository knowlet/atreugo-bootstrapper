package routes

import "github.com/savsgio/atreugo/v11"

// GetApiIndexHandler handles the GET: /
func GetApiIndexHandler(ctx *atreugo.RequestCtx) error {
	return ctx.TextResponse("Hello API Group")
}

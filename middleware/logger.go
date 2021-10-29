package middleware

import (
	"github.com/codelax/paper/context"
	"log"
)

func Logger(handler context.HandlerFunc) context.HandlerFunc {
	return func(ctx *context.Context) {
		log.Println(ctx.Request.Method, ":", ctx.Request.RequestURI)
		handler(ctx)
	}
}

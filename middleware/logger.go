package middleware

import (
	"github.com/codelax/paper/context"
	"log"
)

func Logger(handler context.Handler) context.Handler {
	return func(ctx context.Context) {
		log.Println(ctx.Request.Method, ":", ctx.Request.RequestURI)
		handler(ctx)
	}
}
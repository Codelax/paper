package router

import (
	"github.com/codelax/paper/context"
	"net/http"
)

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var validRoute *Route
	var ctx = &context.Context{
		Response: w,
		Request:  r,
		Params:   map[string]string{},
	}
	for _, route := range router.routes {
		params := route.parser(r.RequestURI)
		if len(params) == 0 || route.method != r.Method {
			continue
		}
		params = params[1:]
		if len(route.params) != 0 {
			for k := range params {
				ctx.Params[route.params[k]] = params[k]
			}
		}
		validRoute = &route
		break
	}
	if validRoute != nil {
		handler := validRoute.handler
		for _, middleware := range router.middlewares {
			handler = middleware(handler)
		}
		handler(ctx)
	} else {
		ctx.Respond(http.StatusNotFound)
	}
}

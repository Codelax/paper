package router

import (
	"fmt"
	"github.com/codelax/paper/context"
	"net/http"
	"regexp"
)

type Router struct {
	routes          []Route
	middlewares     []Middleware
	pathParamsRegex *regexp.Regexp
}

func NewRouter() *Router {
	return &Router{}
}

func (router *Router) AddMiddleware(middleware Middleware) {
	router.middlewares = append(router.middlewares, middleware)
}

func (router *Router) Add(method string, path string, handler context.Handler) {
	var route = Route{
		method:  method,
		handler: handler,
		parser:  router.paramParser(path),
		params:  router.listParams(path),
	}
	router.routes = append(router.routes, route)
}

func (router *Router) GET(path string, handler context.Handler) {
	router.Add(http.MethodGet, path, handler)
}

func (router *Router) POST(path string, handler context.Handler) {
	router.Add(http.MethodPost, path, handler)
}

func (router *Router) PATCH(path string, handler context.Handler) {
	router.Add(http.MethodPatch, path, handler)
}

func (router *Router) PUT(path string, handler context.Handler) {
	router.Add(http.MethodPut, path, handler)
}

func (router *Router) DELETE(path string, handler context.Handler) {
	router.Add(http.MethodDelete, path, handler)
}

func (router *Router) Serve(address string) error {
	fmt.Println(paperSplash)
	fmt.Printf("Listening on %v\n", address)
	return http.ListenAndServe(address, router)
}

const paperSplash = "______\n" +
	"| ___ \\\n" +
	"| |_/ /_ _ _ __   ___ _ __ \n" +
	"|  __/ _` | '_ \\ / _ \\ '__|\n" +
	"| | | (_| | |_) |  __/ |\n" +
	"\\_|  \\__,_| .__/ \\___|_|\n" +
	"          | |\n" +
	"          |_|"

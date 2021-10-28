package main

import (
	"github.com/codelax/paper"
	"github.com/codelax/paper/middleware"
	"log"
	"net/http"
)

func testHandler(ctx paper.Context) {
	ctx.Respond(http.StatusOK)
}

func main() {
	r := paper.NewRouter()
	r.AddMiddleware(middleware.Logger)
	r.GET("/", testHandler)
	log.Fatal(r.Serve(":1337"))
}

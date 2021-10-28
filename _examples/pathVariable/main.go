package main

import (
	"fmt"
	"github.com/codelax/paper"
	"github.com/codelax/paper/middleware"
	"log"
	"net/http"
)

func testHandler(ctx paper.Context) {
	out := fmt.Sprintf("id : %s", ctx.Params["id"])
	ctx.String(http.StatusOK, out)
}

func main() {
	r := paper.NewRouter()
	r.AddMiddleware(middleware.Logger)
	r.Add("/:id", testHandler)
	log.Fatal(r.Serve(":1337"))
}

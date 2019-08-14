// +build go1.9

package paper

import (
	"github.com/codelax/paper/context"
	"github.com/codelax/paper/router"
)

type (
	Context = context.Context
	Router = router.Router
	Middleware = router.Middleware
)

func NewRouter() *router.Router {
	return router.NewRouter()
}
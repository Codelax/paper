package router

import "github.com/codelax/paper/context"

type Middleware func (handler context.Handler) context.Handler
package router

import "github.com/codelax/paper/context"

type Middleware func (handler context.HandlerFunc) context.HandlerFunc

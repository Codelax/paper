package context

import (
	"github.com/codelax/paper/contentTypes"
	"net/http"
)

type Context struct {
	Response http.ResponseWriter
	Request  *http.Request
	Params map[string]string
}

func (c *Context) Respond(statusCode int) {
	c.Response.WriteHeader(statusCode)
}

func (c *Context) SetContentType(contentType string) {
	c.Response.Header().Add("Content-Type", contentType)
}

func (c *Context) String(statusCode int, str string) {
	c.Response.Write([]byte(str))
	c.SetContentType(contentTypes.Bytes)
}
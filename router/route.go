package router

import (
	"github.com/codelax/paper/context"
	"log"
	"regexp"
)

type Route struct {
	method string
	handler context.Handler
	parser pathParser
	params []string
}

type pathParser func (path string) []string

func (router *Router) setupParamParser() {
	var err error
	if router.pathParamsRegex == nil {
		router.pathParamsRegex, err = regexp.Compile(":([[:alnum:]]+)")
		if err != nil {
			log.Println(err)
		}
	}
}

// We turn route into a regex that will parse path arguments
// path pattern (:id) are replaced by simple string regex
func (router *Router) paramParser(pathDef string) pathParser {
	router.setupParamParser()
	pathDef = string(router.pathParamsRegex.ReplaceAll([]byte(pathDef), []byte("([[:alnum:]]+)")))
	pathDef = "^" + pathDef + "$"
	pathRegex, err := regexp.Compile(pathDef)
	if err != nil {
		log.Println(err)
	}
	return func(path string) []string {
		return pathRegex.FindStringSubmatch(path)
	}
}

func (router *Router) listParams(path string) []string {
	router.setupParamParser()
	matches := router.pathParamsRegex.FindAllStringSubmatch(path, -1)
	var submatches = []string{}
	for _, match := range matches {
		if len(match) > 1 {
			submatches = append(submatches, match[1])
		}
	}
	return submatches
}

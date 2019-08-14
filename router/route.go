package router

import (
	"github.com/codelax/paper/context"
	"log"
	"regexp"
)

type Route struct {
	handler context.Handler
	parser pathParser
	params []string
}

type pathParser func (path string) []string

func (router *Router) setupParamParser() {
	var err error
	if router.pathParamsRegex == nil {
		router.pathParamsRegex, err = regexp.Compile("{(.+?)}")
		if err != nil {
			log.Println(err)
		}
	}
}

func (router *Router) paramParser(pathDef string) pathParser {
	router.setupParamParser()
	pathDef = string(router.pathParamsRegex.ReplaceAll([]byte(pathDef), []byte("([\\S]+)")))
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
package http

import "github.com/julienschmidt/httprouter"

type Router interface {
	Register(router *httprouter.Router)
}

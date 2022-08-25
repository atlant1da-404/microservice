package http

import "github.com/julienschmidt/httprouter"

// Router is a controllers interface to register api´s
type Router interface {
	// Register register all Api´s endpoints
	Register(router *httprouter.Router)
}

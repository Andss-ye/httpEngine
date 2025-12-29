package engine

import "net/http"

type Context struct {
	Writer http.ResponseWriter
	Request *http.Request

	// se irán agregando:
	// Params
	// Data
	// index del middleware
}

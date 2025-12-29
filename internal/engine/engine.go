package engine

import "net/http"

type Engine struct {
	middlewares []MiddlewareFunc
	handler     HandlerFunc
}

func New() *Engine {
	return &Engine{
		middlewares: make([]MiddlewareFunc, 0),
	}
}

func (e *Engine) Use(m MiddlewareFunc) {
	e.middlewares = append(e.middlewares, m)
}

func (e *Engine) Handle(h HandlerFunc) {
	e.handler = h
}

func (e *Engine) buildChain() HandlerFunc {
	h := e.handler

	for i := len(e.middlewares) - 1; i >= 0; i-- {
		m := e.middlewares[i]
		next := h

		h = func(ctx *Context) {
			m(ctx, next)
		}
	}

	return h
}

// Implementa http.Handler
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e.handler == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not handler implemented"))
		return
	}

	ctx := &Context{
		Writer: w,
		Request: r,
	}
	
	chain := e.buildChain()
	chain(ctx)
}

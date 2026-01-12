package engine

import "net/http"

type Engine struct {
	middlewares []MiddlewareFunc
	handler     HandlerFunc
	
	routes map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{
		middlewares: make([]MiddlewareFunc, 0),
		routes:      make(map[string]HandlerFunc),
	}
}

func (e *Engine) Use(m MiddlewareFunc) {
	e.middlewares = append(e.middlewares, m)
}

func (e *Engine) Handle(h HandlerFunc) {
	e.handler = h
}

func (e *Engine) buildChain(final HandlerFunc) HandlerFunc {
	h := final

	for i := len(e.middlewares) - 1; i >= 0; i-- {
		m := e.middlewares[i]
		next := h

		h = func(ctx *Context) {
			m(ctx, next)
		}
	}

	return h
}

func (e *Engine) HandleRoute(method, path string, h HandlerFunc) {
	key := method + ":" + path
	e.routes[key] = h
}

func (e *Engine) GET(path string, h HandlerFunc) {
	e.HandleRoute("GET", path, h)
}

func (e *Engine) POST(path string, h HandlerFunc) {
	e.HandleRoute("POST", path, h)
}

func (e *Engine) PUT(path string, h HandlerFunc) {
	e.HandleRoute("PUT", path, h)
}

func (e *Engine) DELETE(path string, h HandlerFunc) {
	e.HandleRoute("DELETE", path, h)
}

// Implementa http.Handler
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := &Context{
		Writer: w,
		Request: r,
	}
	
	key := r.Method + ":" + r.URL.Path
	handler, ok := e.routes[key]
	
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 not found"))
		return
	}
	
	chain := e.buildChain(handler)
	chain(ctx)
}

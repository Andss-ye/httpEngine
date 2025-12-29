package engine

type HandlerFunc func(*Context)

type MiddlewareFunc func(*Context, HandlerFunc)
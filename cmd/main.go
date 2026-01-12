package main

import (
	"log"
	"net/http"

	"github.com/andrew/http-engine/internal/engine"
)

func main() {
	app := engine.New()

	// Middleware 1: Logger
	app.Use(func(ctx *engine.Context, next engine.HandlerFunc) {
		log.Println("→ Incoming request:", ctx.Request.Method, ctx.Request.URL.Path)
		next(ctx)
		log.Println("← Request finished")
	})

	// Middleware 2: Auth fake
	app.Use(func(ctx *engine.Context, next engine.HandlerFunc) {
		token := ctx.Request.Header.Get("Authorization")

		if token == "a" {
			ctx.Writer.WriteHeader(401)
			ctx.Writer.Write([]byte("Unauthorized"))
			return 
		}

		next(ctx)
	})

	app.HandleRoute("GET", "/", func(ctx *engine.Context) {
		ctx.Writer.Write([]byte("Home"))
	})

	app.HandleRoute("GET", "/health", func(ctx *engine.Context) {
		ctx.Writer.Write([]byte("OK"))
	})	
	
	// Handler final
	app.Handle(func(ctx *engine.Context) {
		ctx.Writer.Write([]byte("Welcome, authenticated user"))
	})

	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", app)
}

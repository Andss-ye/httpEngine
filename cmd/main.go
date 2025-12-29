package main

import (
	"log"
	"net/http"

	"github.com/andrew/http-engine/internal/engine"
)

func main() {
	app := engine.New()

	app.Use(func(ctx *engine.Context, next engine.HandlerFunc) {
		log.Println("Before handler")
		next(ctx)
		log.Println("After handler")
	})

	app.Handle(func(ctx *engine.Context) {
		ctx.Writer.Write([]byte("Hello from engine"))
	})

	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", app)
}

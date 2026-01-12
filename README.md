# Mini HTTP Engine (Go)

Este proyecto es un **HTTP engine minimalista escrito en Go**, creado principalmente con fines **educativos**.

La idea no es crear “otro framework”, sino **entender cómo funcionan por dentro** las herramientas que usamos todos los días:  
cómo `net/http` maneja requests, cómo se conectan los handlers, cómo funcionan los middlewares y cómo se puede construir un router desde cero.

Todo está basado en la **libreria estándar de Go**, sin dependencias externas.

---

## 🎯 ¿Por qué existe este proyecto?

Porque muchas veces usamos frameworks sin saber realmente qué pasa detrás.

Este proyecto busca:
- entender cómo `net/http` usa interfaces (`http.Handler`)
- construir un engine propio, paso a paso
- aprender cómo funciona una middleware chain
- implementar un router simple y explícito
- mantener el código fácil de leer y razonar

No está pensado para producción, sino para **aprender**.

---

## 🧱 Qué incluye

- ✅ Engine que implementa `http.Handler`
- ✅ Middlewares globales con control de flujo (`next`)
- ✅ Router mínimo basado en `method + path`
- ✅ Shortcuts (`GET`, `POST`, etc.)
- ✅ Context propio para manejar request y response

---

## 🚫 Qué no incluye (a propósito)

Hay muchas cosas que **no están**, y es totalmente intencional:

- ❌ Params (`/users/:id`)
- ❌ Wildcards
- ❌ Middlewares por ruta
- ❌ Helpers tipo `ctx.JSON()` o `ctx.Text()`
- ❌ Manejo avanzado de errores

---

## Ejemplo de uso uwu

```go
func main() {
	app := engine.New()

	app.Use(func(ctx *engine.Context, next engine.HandlerFunc) {
		log.Println(ctx.Request.Method, ctx.Request.URL.Path)
		next(ctx)
	})

	app.GET("/", func(ctx *engine.Context) {
		ctx.Writer.Write([]byte("Home"))
	})

	app.GET("/health", func(ctx *engine.Context) {
		ctx.Writer.Write([]byte("OK"))
	})

	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", app)
}
```


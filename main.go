package main

import (
	//"net/http"

	"github.com/danielgtaylor/huma"
	"github.com/danielgtaylor/huma/cli"
	"github.com/danielgtaylor/huma/responses"
)

func main() {
	// Create a new router & CLI with default middleware.
	app := cli.NewRouter("Minimal Example", "1.0.0")

	// Declare the root resource and a GET operation on it.
	app.Resource("/").Get("get-root", "Get a short text message",
		// The only response is HTTP 200 with text/plain
		responses.OK().ContentType("text/plain"),
	).Run(func(ctx huma.Context) {
		// This is he handler function for the operation. Write the response.
		ctx.Header().Set("Content-Type", "text/plain")
		ctx.Write([]byte("Hello, world"))
	})

	// Run the CLI. When passed no arguments, it starts the server.
	app.Run()
}

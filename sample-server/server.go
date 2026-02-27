package main

// @title Sample Server API
// @version 1.0
// @description Sample server with Swagger docs.
// @BasePath /

import (
	"fmt"
	"net/http"

	_ "github.com/alfonsodelr/sample-server/docs"
	"github.com/alfonsodelr/sample-server/handlers"

	httpSwagger "github.com/swaggo/http-swagger"
)

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/version", handlers.HandlerVersionRequest)

	// Swagger UI at http://localhost:8080/swagger/index.html
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	return mux
}

func main() {
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", newMux())
}

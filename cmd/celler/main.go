package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-chai/examples/cmd/celler/docs" // This is required to be able to serve the stored swagger spec in prod
	"github.com/go-chai/examples/pkg/router"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	r := router.GetRoutes()

	// This should be used in prod to serve the swagger spec
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	fmt.Println("Listening on port 8080")
	fmt.Println("Find the swagger spec at http://localhost:8080/swagger/")

	http.ListenAndServe(":8080", r)
}

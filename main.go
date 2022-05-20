package main

import (
	"net/http"

	"github.com/byamada/ALURA/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}

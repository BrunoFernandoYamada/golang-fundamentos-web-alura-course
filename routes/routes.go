package routes

import (
	"net/http"

	"github.com/byamada/ALURA/controller"
)

func LoadRoutes() {
	http.HandleFunc("/", controller.FindAllProducts)
}

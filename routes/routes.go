package routes

import (
	"net/http"

	"github.com/byamada/ALURA/controller"
)

func LoadRoutes() {
	http.HandleFunc("/", controller.FindAllProducts)
	http.HandleFunc("/new", controller.CreateProduct)
	http.HandleFunc("/insert", controller.SaveProduct)
	http.HandleFunc("/remove", controller.DeleteProduct)
	http.HandleFunc("/edit", controller.Edit)
	http.HandleFunc("/update", controller.Update)
}

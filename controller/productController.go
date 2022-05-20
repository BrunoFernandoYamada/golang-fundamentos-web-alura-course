package controller

import (
	"html/template"
	"net/http"

	"github.com/byamada/ALURA/model"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func FindAllProducts(w http.ResponseWriter, r *http.Request) {
	produtos := model.FindAll()
	temp.ExecuteTemplate(w, "Index", produtos)
}

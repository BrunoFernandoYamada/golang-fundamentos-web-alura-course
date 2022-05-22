package controller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/byamada/ALURA/model"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func FindAllProducts(w http.ResponseWriter, r *http.Request) {
	produtos := model.FindAll()
	temp.ExecuteTemplate(w, "Index", produtos)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func SaveProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		precoConvertido, err := strconv.ParseFloat(r.FormValue("preco"), 64)
		if err != nil {
			log.Fatal(err)
		}

		quantidadeConvertida, err := strconv.Atoi(r.FormValue("quantidade"))
		if err != nil {
			log.Fatal(err)
		}

		p := model.Produto{}
		p.Nome = r.FormValue("nome")
		p.Descricao = r.FormValue("descricao")
		p.Preco = precoConvertido
		p.Quantidade = quantidadeConvertida

		model.Save(p)

		http.Redirect(w, r, "/", 301)
	}

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	productId := r.URL.Query().Get("id")
	model.Delete(productId)

	http.Redirect(w, r, "/", 301)

}

func Edit(w http.ResponseWriter, r *http.Request) {

	productId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err.Error())
	}

	product := model.FindById(productId)

	temp.ExecuteTemplate(w, "Edit", product)

}

func Update(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		precoConvertido, err := strconv.ParseFloat(r.FormValue("preco"), 64)
		if err != nil {
			panic(err.Error())
		}

		quantidadeConvertida, err := strconv.Atoi(r.FormValue("quantidade"))
		if err != nil {
			panic(err.Error())
		}

		idConvertido, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			panic(err.Error())
		}

		p := model.Produto{}
		p.Id = idConvertido
		p.Nome = r.FormValue("nome")
		p.Descricao = r.FormValue("descricao")
		p.Preco = precoConvertido
		p.Quantidade = quantidadeConvertida

		model.Update(p)

		http.Redirect(w, r, "/", 301)

	}

}

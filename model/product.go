package model

import "github.com/byamada/ALURA/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func FindAll() []Produto {

	db := db.ConnectDB()
	defer db.Close()

	selectProducts, err := db.Query("Select * from product")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectProducts.Next() {

		var nome, descricao string
		var id, quantidade int
		var preco float64

		err := selectProducts.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)

	}

	return produtos

}

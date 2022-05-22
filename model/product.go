package model

import (
	"github.com/byamada/ALURA/db"
)

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

	selectProducts, err := db.Query("Select * from product order by id asc")
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

func Save(p Produto) {

	db := db.ConnectDB()

	query := "INSERT INTO Product (nome, descricao, preco, quantidade) values ($1, $2, $3, $4)"

	PreparedStatement, err := db.Prepare(query)
	if err != nil {
		panic(err.Error())
	}

	PreparedStatement.Exec(p.Nome, p.Descricao, p.Preco, p.Quantidade)

	defer db.Close()

}

func Delete(id string) {

	db := db.ConnectDB()

	query := "DELETE FROM Product WHERE id = $1"

	preparedStatement, err := db.Prepare(query)
	if err != nil {
		panic(err.Error())
	}

	preparedStatement.Exec(id)

	defer db.Close()

}

func FindById(id int) Produto {

	db := db.ConnectDB()

	producFromDb, err := db.Query("SELECT * FROM Product WHERE id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}

	for producFromDb.Next() {

		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := producFromDb.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

	}

	return p

}

func Update(p Produto) {

	produtoFormDb := FindById(p.Id)
	if produtoFormDb.Id == 0 {
		panic("Produto não encontrado!")
	}

	db := db.ConnectDB()

	query := "UPDATE Product SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id = $5"

	preparedStatement, err := db.Prepare(query)
	if err != nil {
		panic(err.Error())
	}

	preparedStatement.Exec(p.Nome, p.Descricao, p.Preco, p.Quantidade, p.Id)

	db.Close()
}

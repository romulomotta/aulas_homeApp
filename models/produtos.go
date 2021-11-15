package models

import (
	"romulo/db"

	_ "github.com/lib/pq"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func SelectAllProducts() []Produto {
	db := db.ConnectDataBase()

	selectProducts, err := db.Query("select * from produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectProducts.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProducts.Scan(&id, &nome,
			&descricao, &preco, &quantidade)
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
	defer db.Close()

	return produtos
}

func CreateNewProduct(name, desc string, price float64, qtd int) {
	db := db.ConnectDataBase()

	insertIntoDataBase, err := db.Prepare(
		"INSERT INTO produtos (" +
			"nome, " +
			"descricao, " +
			"preco, " +
			"quantidade" +
			") values ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insertIntoDataBase.Exec(name, desc, price, qtd)

	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConnectDataBase()

	toDelete, err := db.Prepare("DELETE FROM produtos WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}

	toDelete.Exec(id)
	defer db.Close()
}

func EditProduct(id string) Produto {
	db := db.ConnectDataBase()

	dataProduct, err := db.Query(
		"SELECT * FROM produtos WHERE id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	productToUpdate := Produto{}

	for dataProduct.Next() {
		var id, qtd int
		var name, desc string
		var price float64

		err = dataProduct.Scan(&id, &name, &desc, &price, &qtd)
		if err != nil {
			panic(err.Error())
		}

		productToUpdate.Id = id
		productToUpdate.Nome = name
		productToUpdate.Descricao = desc
		productToUpdate.Preco = price
		productToUpdate.Quantidade = qtd
	}
	defer db.Close()
	return productToUpdate
}

func UpdateProduct(id int, nome, descricao string, preco float64, qtd int) {
	db := db.ConnectDataBase()

	ToUpdateProduct, err := db.Prepare(
		"UPDATE produtos SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5")

	if err != nil {
		panic(err.Error())
	}
	ToUpdateProduct.Exec(nome, descricao, preco, qtd, id)
	db.Query("select * from produtos order by id asc")
	defer db.Close()
}

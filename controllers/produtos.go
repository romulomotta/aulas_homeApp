package controllers

import (
	"fmt"
	"log"
	"net/http"
	"romulo/models"
	"strconv"
	"text/template"
)

var temp = template.Must(
	template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.SelectAllProducts()

	temp.ExecuteTemplate(w, "Index", allProducts)
}

func NewProduct(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "NewProduct", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("nome")
		desc := r.FormValue("descricao")
		price := r.FormValue("preco")
		qtd := r.FormValue("quantidade")

		newPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro na conversão de preço:", err)
		}

		newQtd, err := strconv.Atoi(qtd)
		if err != nil {
			log.Println("Erro na conversão de quantidade:", err)
		}

		models.CreateNewProduct(name, desc, newPrice, newQtd)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {

	productId := r.URL.Query().Get("id")
	models.DeleteProduct(productId)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	fmt.Println(productId)
	produto := models.EditProduct(productId)
	fmt.Println(produto)
	println(w)
	// temp.ExecuteTemplate(w, "Edit", produto)
	temp.ExecuteTemplate(w, "/", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		convertId, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão de id:", err)
		}

		convertQtd, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão de quantidade:", err)
		}

		convertPreco, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão de preço:", err)
		}

		models.UpdateProduct(convertId, nome, descricao, convertPreco, convertQtd)
	}
	http.Redirect(w, r, "/", 301)
}

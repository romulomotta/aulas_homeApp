package main

import (
	"net/http"
	"romulo/routes"
)

func main() {
	routes.LoadRoutes()
	println("Sevidor rodando na porta 8000.")
	http.ListenAndServe(":8000", nil)
}

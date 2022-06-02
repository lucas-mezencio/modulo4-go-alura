package main

import (
	"fmt"
	"go-rest/models"
	"go-rest/routes"
)

func main() {

	models.Personalities = []models.Personality{
		{Nome: "nome 1", Historia: "historia 1"},
		{Nome: "nome 2", Historia: "historia 2"},
	}

	fmt.Println("iniciando servidor")
	routes.HandleRequest()
}

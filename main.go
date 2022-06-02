package main

import (
	"fmt"
	"go-rest/routes"
)

func main() {
	fmt.Println("iniciando servidor")
	routes.HandleRequest()
}

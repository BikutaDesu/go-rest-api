package main

import (
	"fmt"

	"github.com/bikutadesu/go-rest-api/database"
	"github.com/bikutadesu/go-rest-api/models"
	"github.com/bikutadesu/go-rest-api/routes"
)

func main() {
	models.People = []models.Person{
		{Id: 1, Name: "Gleiso", History: "Main caitlyn"},
		{Id: 2, Name: "Gui", History: "Cara do front, odeia anime, mas é asiático"},
	}

	database.Connect()

	fmt.Println("Iniciando servidor REST...")
	routes.HandleRequest()
}

package main

import (
	"github.com/gubisss/api-go-gin/database"
	"github.com/gubisss/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}

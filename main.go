package main

import (
	"github.com/IgorFVicente/api-go-gin/database"
	"github.com/IgorFVicente/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}

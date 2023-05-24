package main

import (
	"simple-gin-server/config"
	"simple-gin-server/database"
	"simple-gin-server/routes"
)

func main() {
	config.Init()
	database.Init()
	routes.Init()
}

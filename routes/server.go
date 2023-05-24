package routes

import (
	"simple-gin-server/config"
	"simple-gin-server/database"

	"github.com/gin-gonic/gin"
)

func setUpRouter() *gin.Engine {
	database := database.GetDatabaseSession()
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	healthCheckRouter := HealthCheckRouter(database)
	healthCheckRouter(router.Group("/"))

	itemRouter := ItemRouter(database)
	itemRouter(router.Group("/"))

	return router
}

func Init() {
	conf := config.GetConfig()
	router := setUpRouter()
	router.Run(conf.Application.Address)
}

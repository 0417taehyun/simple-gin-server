package routes

import (
	"database/sql"
	"simple-gin-server/controllers"

	"github.com/gin-gonic/gin"
)

func ItemRouter(database *sql.DB) func(*gin.RouterGroup) {
	item := controllers.NewItemController(database)

	return func(router *gin.RouterGroup) {
		router.GET("/item/:id", item.GetItem)
	}
}

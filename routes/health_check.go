package routes

import (
	"database/sql"
	"simple-gin-server/controllers"

	"github.com/gin-gonic/gin"
)

const (
	LIVENESS  = "livez"
	READINESS = "readyz"
)

func HealthCheckRouter(database *sql.DB) func(*gin.RouterGroup) {
	healthCheck := controllers.NewHealthCheckController(database)

	return func(router *gin.RouterGroup) {
		router.GET(LIVENESS, healthCheck.ProbeLiveness)
		router.GET(READINESS, healthCheck.ProbeReadiness)
	}
}

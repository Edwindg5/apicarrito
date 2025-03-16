package routes

import (
	"mqtt/src/core"
	"mqtt/src/procesamiento/application"
	"mqtt/src/procesamiento/infraestructure/routes"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, messageUsecase *usecase.MessageUsecase) {	
	r.Use(core.CORSMiddleware())

	apiRoutes := r.Group("/api")
	{
		routes.RegisterRoutes(apiRoutes, messageUsecase)
	}
}

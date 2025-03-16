package routes

import (
	"mqtt/src/procesamiento/application"
	"mqtt/src/procesamiento/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes registra las rutas de la API
func RegisterRoutes(r *gin.RouterGroup, messageUsecase *usecase.MessageUsecase) {
	r.GET("/messages", func(c *gin.Context) {
		controllers.GetMessages(c, messageUsecase)
	})
}

package main

import (
	"log"
	"mqtt/src/core/routes"
	"mqtt/src/procesamiento/application"
	"mqtt/src/procesamiento/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting API...")

	// Iniciar el caso de uso de manejo de mensajes
	messageUsecase := usecase.NewMessageUsecase()

	// Iniciar MQTT Listener
	controllers.StartMQTTListener(messageUsecase)

	// Configurar router
	r := gin.Default()
	routes.SetupRouter(r, messageUsecase)

	// Iniciar servidor en el puerto 8080
	r.Run(":8080")
}

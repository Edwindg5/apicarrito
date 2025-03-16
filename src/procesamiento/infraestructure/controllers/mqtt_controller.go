package controllers

import (
	"log"
	"net/http"
	"sync"

	"mqtt/src/procesamiento/application"
	"mqtt/src/procesamiento/domain/entities"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

const (
	mqttBroker  = "tcp://52.7.35.94:1883"
	rabbitMQURL = "amqp://admin:admin@52.7.35.94:5672/"
	topic       = "/MedicHealth/sensores/temperatura"
	queueName   = "objeto_basura"
)

var mutex sync.Mutex

// StartMQTTListener inicia la conexión MQTT y envía mensajes a RabbitMQ
func StartMQTTListener(messageUsecase *usecase.MessageUsecase) {
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Fatalf("Error connecting to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	opts := mqtt.NewClientOptions().AddBroker(mqttBroker)
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("MQTT Connection Error: %v", token.Error())
	}

	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		mutex.Lock()
		defer mutex.Unlock()

		message := entities.Message{Content: string(msg.Payload())}

		// Guardar en la capa de aplicación
		messageUsecase.SaveMessage(message)

		err := ch.Publish(
			"",
			queueName,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        msg.Payload(),
			},
		)
		if err != nil {
			log.Printf("Failed to publish message to RabbitMQ: %v", err)
		}
	})
}

// GetMessages obtiene los mensajes almacenados
func GetMessages(c *gin.Context, messageUsecase *usecase.MessageUsecase) {
	messages := messageUsecase.GetAllMessages()
	c.JSON(http.StatusOK, gin.H{"messages": messages})
}

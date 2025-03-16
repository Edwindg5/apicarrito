package interfaces

import "mqtt/src/procesamiento/domain/entities"

// MessageService define los métodos para manejar mensajes
type MessageService interface {
	SaveMessage(message entities.Message)
	GetAllMessages() []entities.Message
}

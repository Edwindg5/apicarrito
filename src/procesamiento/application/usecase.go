package usecase

import (
	"sync"

	"mqtt/src/procesamiento/domain/entities"
)

// MessageUsecase maneja la lógica de los mensajes
type MessageUsecase struct {
	messages []entities.Message
	mutex    sync.Mutex
}

// NewMessageUsecase crea una nueva instancia de MessageUsecase
func NewMessageUsecase() *MessageUsecase {
	return &MessageUsecase{}
}

// SaveMessage guarda un mensaje recibido
func (u *MessageUsecase) SaveMessage(message entities.Message) {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	u.messages = append(u.messages, message)
}

// GetAllMessages devuelve todos los mensajes almacenados
func (u *MessageUsecase) GetAllMessages() []entities.Message {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	return u.messages
}

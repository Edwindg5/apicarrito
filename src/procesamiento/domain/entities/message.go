// mqtt-sp32/src/procesamiento/domain/entities/message.go
package entities

// Message representa un mensaje recibido desde MQTT
type Message struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

package models

import (
	"encoding/json"
	"fmt"
)


type Message struct {
	SenderName string
	Message    string
	ChatGroupName string
}


// desserializa uma linha JSON recebida pela conexão TCP.
func FromJsonString(data string) (Message, error) {

	var msg Message

	if len(data) <= 0 {
		return Message{}, fmt.Errorf("invalid data")
	}

	err := json.Unmarshal([]byte(data), &msg)
	if err != nil {
		return Message{}, fmt.Errorf("invalid data")
	}

	return msg, nil

}

// serializa a mensagem em JSON com \n no final para o Scanner.
func (m Message) ToJsonString() string {
	data, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s\n", string(data))
}
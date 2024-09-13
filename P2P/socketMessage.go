package main

import "encoding/json"

type SocketMessage struct {
	TypeMsg string `json:"type"`
	Data    string `json:"data"`
}

func NewSocketMessage(typeMsg string, data string) *SocketMessage {
	return &SocketMessage{
		TypeMsg: typeMsg,
		Data:    data,
	}
}

func (s *SocketMessage) GetType() string {
	return s.TypeMsg
}

func (s *SocketMessage) GetData() string {
	return s.Data
}

func ParseSocketMessage(data []byte) (*SocketMessage, error) {
	var sockMsg SocketMessage

	err := json.Unmarshal(data, &sockMsg)

	if err != nil {
		return nil, err
	}

	return &sockMsg, nil

}

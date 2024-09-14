package main

import (
	"bufio"
	"encoding/json"
	"net"
)

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

func (s *SocketMessage) ParseToJson() (string, error) {
	str, err := json.Marshal(s)

	if err != nil {
		return "", err
	}

	return string(str), nil
}

func GetParsedDataFromMessage[T any](msg *SocketMessage, parser func(string) (T, error)) (T, error) {
	return parser(msg.GetData())
}

// Encapsulates the data and type of the message in a socket message
func ParseDataAndEncapsulateSocketMessage[T any](typeMsg string, data T, parser func(t T) ([]byte, error)) (string, error) {
	parsed, err := parser(data)
	if err != nil {
		return "", err
	}

	sockData, err := encapsulateSocketMessage(typeMsg, parsed)
	if err != nil {
		return "", err
	}
	return sockData, nil
}

func encapsulateSocketMessage(typeMsg string, data []byte) (string, error) {
	sockMsg := NewSocketMessage(typeMsg, string(data))
	sockData, err := sockMsg.ParseToJson()
	if err != nil {
		return "", err
	}
	return sockData + "\n", nil
}

func ReadMessage(conn net.Conn) (*SocketMessage, error) {
	reader := bufio.NewReader(conn)
	response, err := reader.ReadString('\n')

	if err != nil {
		return nil, err
	}

	return DecodeSocketMessage([]byte(response))
}

func DecodeSocketMessage(data []byte) (*SocketMessage, error) {
	var sockMsg SocketMessage

	err := json.Unmarshal(data, &sockMsg)

	if err != nil {
		return nil, err
	}

	return &sockMsg, nil

}

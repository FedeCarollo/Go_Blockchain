package p2p

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

func DecodeSocketMessage(data []byte) (*SocketMessage, error) {
	var sockMsg SocketMessage

	err := json.Unmarshal(data, &sockMsg)

	if err != nil {
		return nil, err
	}

	return &sockMsg, nil

}

func (s *SocketMessage) ParseToJson() (string, error) {
	data, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func encapsulateSocketMessage(typeMsg string, data []byte) (string, error) {
	sockMsg := NewSocketMessage(typeMsg, string(data))
	sockData, err := sockMsg.ParseToJson()
	if err != nil {
		return "", err
	}
	return sockData + "\n", nil
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

// Decodes data field of the socket message and parse it
func DecodeMessage[T any](sockMsg *SocketMessage, parser func(data []byte) (T, error)) (T, error) {
	return parser([]byte(sockMsg.GetData()))
}

// Get data sent by the socket message and parse it
func ParseMessageToData[T any](data []byte, parser func(data []byte) T) T {
	sockMsg, _ := DecodeSocketMessage(data)
	return parser([]byte(sockMsg.GetData()))
}

// Sends a message to the connection
func SendMessage[T any](conn net.Conn, typeMsg string, data T, parser func(t T) ([]byte, error)) error {
	msg, err := ParseDataAndEncapsulateSocketMessage(typeMsg, data, parser)
	if err != nil {
		return err
	}
	_, err = conn.Write([]byte(msg))
	return err
}

// Reads incoming message and returns the decoded socket message
func ReadMessage(conn net.Conn) (*SocketMessage, error) {
	reader := bufio.NewReader(conn)
	response, err := reader.ReadString('\n')

	if err != nil {
		return nil, err
	}

	return DecodeSocketMessage([]byte(response))
}

// Reads incoming message, decodes it and parses it
func ReadMessageAndParse[T any](conn net.Conn, parser func(data []byte) (T, error)) (T, error) {
	sockMsg, err := ReadMessage(conn)
	if err != nil {
		var t T
		return t, err
	}
	return DecodeMessage(sockMsg, parser)
}

// Send message with no content
func EncapsulateMessageNoContent(conn net.Conn, typeMsg string) *SocketMessage {
	sockMsg := NewSocketMessage(typeMsg, "")
	return sockMsg
}

// Send message with no content
func SendMessageNoContent(conn net.Conn, typeMsg string) error {
	msg, err := encapsulateSocketMessage(typeMsg, []byte(""))
	if err != nil {
		return err
	}
	_, err = conn.Write([]byte(msg))

	return err
}

package main

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

func ParseSocketMessage(data string) *SocketMessage {
	return &SocketMessage{
		TypeMsg: "type",
		Data:    data,
	}
}

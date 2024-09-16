package main

import (
	"errors"
	"log"
	"net"

	"github.com/sirupsen/logrus"
)

type Server struct {
	tracker *Tracker
}

func startServer(tracker *Tracker) {
	// Create a new server
	server := NewServer(tracker)

	// Start the server
	server.Start()
}

func NewServer(tracker *Tracker) *Server {
	return &Server{
		tracker: tracker,
	}
}

func (s *Server) getAddress() string {
	return s.tracker.GetInfo().GetAddress()
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", s.getAddress())

	if err != nil {
		logrus.Fatalf("Cannot start server: %v", err)
	}

	logrus.Infof("Server started on %s", s.getAddress())
	log.Default().Printf("Server started on %s", s.getAddress())

	for {
		conn, err := listener.Accept()

		if err != nil {
			logrus.Errorf("Error accepting connection: %v", err)
			continue
		}

		go s.handleConnection(conn)
	}
}

// First level of handling the connection message
func (s *Server) handleConnection(conn net.Conn) {
	//TODO: Implement
	defer conn.Close()

	logrus.Infof("Connection from %s", conn.RemoteAddr())
	msg, err := ReadMessage(conn)

	if err != nil {
		logrus.Errorf("Error reading message: %v", err)
		return
	}

	logrus.Infof("Received message of type: %s", msg.GetType())
	log.Printf("Received message of type: %s", msg.GetType())

	switch msg.GetType() {
	case "announce":
		handleAnnounce(s.tracker, msg, conn)
	case "ping":
		handlePing(s.tracker, msg, conn)
	default:
		logrus.Errorf("Unknown message type: %s", msg.GetType())
	}

}

// Sends a message to the connection
func SendMessage[T any](conn net.Conn, typeMsg string, data T, parser func(t T) ([]byte, error)) error {
	msg, err := ParseDataAndEncapsulateSocketMessage(typeMsg, data, parser)
	if err != nil {
		return err
	}
	if conn != nil {
		_, err = conn.Write([]byte(msg))
		return err
	} else {
		return errors.New("connection is nil")
	}
}

package main

import (
	"bufio"
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

/*
Read data from the socket connection
and returns it as a string
*/
func (s *Server) readData(conn net.Conn) (string, error) {
	reader := bufio.NewReader(conn)
	data, err := reader.ReadString('\n')

	if err != nil {
		logrus.Errorf("Error reading data: %v", err)
		return "", err
	}
	return data, nil
}

// First level of handling the connection message
func (s *Server) handleConnection(conn net.Conn) {
	//TODO: Implement
	defer conn.Close()

	logrus.Infof("Connection from %s", conn.RemoteAddr())
	data, err := s.readData(conn)
	if err != nil {
		logrus.Errorf("Error reading data: %v", err)
		return
	}

	msg := ParseSocketMessage(data)

	logrus.Infof("Received message of type: %s", msg.GetType())

	switch msg.GetType() {
	case "announce":
		handleAnnounce(s.tracker, msg)
	default:
		logrus.Errorf("Unknown message type: %s", msg.GetType())
	}

}

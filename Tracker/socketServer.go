package main

import (
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

	for {
		conn, err := listener.Accept()

		if err != nil {
			logrus.Errorf("Error accepting connection: %v", err)
			continue
		}

		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	//TODO: Implement
}

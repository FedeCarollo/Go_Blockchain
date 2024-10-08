package p2p

import (
	"bufio"
	"log"
	"net"

	"github.com/sirupsen/logrus"
)

func StartServer(node *Node) error {
	listener, err := net.Listen("tcp", node.GetInfo().GetAddress())

	if err != nil {
		logrus.Errorf("Error starting server: %v", err)
		return err
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			logrus.Errorf("Error accepting connection: %v", err)
			continue
		}

		go handleConnection(conn, node)
	}
}

func handleConnection(conn net.Conn, node *Node) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	data, err := reader.ReadString('\n')

	if err != nil {
		logrus.Errorf("Error reading data: %v", err)
		return
	}

	// Parse the data
	sockMsg, err := DecodeSocketMessage([]byte(data))

	if err != nil {
		logrus.Errorf("Error parsing socket message: %v", err)
		return
	}

	logrus.Infof("Received message: %s", sockMsg)

	// Handle the message
	switch sockMsg.GetType() {
	case "announce":

	case "ping":
		handlePing(conn, node, sockMsg)
	default:
		log.Println(sockMsg.GetType())
		logrus.Errorf("Unknown message type: %s", sockMsg.GetType())
	}
}

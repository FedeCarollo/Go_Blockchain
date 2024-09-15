package p2p

import (
	"bufio"
	"net"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	pingTimeout = 10 * time.Second
)

func PingPeers(node *Node) {
	peers := node.GetPeers()
	for _, peer := range peers {
		go PingPeer(node, peer)
	}
}

func PingPeer(node *Node, peer Peer) {
	conn, err := net.DialTimeout("tcp", peer.GetAddress(), pingTimeout)
	if err != nil {
		logrus.Errorf("Error connecting to peer: %v", err)
		return
	}

	defer conn.Close()

	msg, err := ParseDataAndEncapsulateSocketMessage("ping", node.GetInfo(), ParsePeerToJson)

	if err != nil {
		logrus.Errorf("Error parsing data to send to peer: %v", err)
		return
	}

	_, err = conn.Write([]byte(msg))

	if err != nil {
		logrus.Errorf("Error sending message to peer: %v", err)
		node.peers.Remove(peer)
		return
	}

	reader := bufio.NewReader(conn)

	response, err := reader.ReadString('\n')

	if err != nil {
		logrus.Errorf("Error reading response from peer: %v", err)
		node.peers.Remove(peer)
		return
	}

	logrus.Infof("Response from peer: %v", response)

}

func handlePing(conn net.Conn, node *Node, sockMsg *SocketMessage) {
	peer, err := DecodeMessage(sockMsg, DecodeJsonToPeer)

	if err != nil {
		logrus.Errorf("Error decoding message: %v", err)
		return
	}

	logrus.Infof("Received ping from peer: %v", peer.GetAddress())

	// Send a response
	err = SendMessage(conn, "pong", node.GetInfo(), ParsePeerToJson)

	if err != nil {
		logrus.Errorf("Error sending response to peer: %v", err)
	}
}

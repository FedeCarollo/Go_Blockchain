package main

import (
	"bufio"
	"net"

	"github.com/sirupsen/logrus"
)

func PingPeers(node *Node) {
	peers := node.GetPeers()
	for _, peer := range peers {
		go PingPeer(node, peer)
	}
}

func PingPeer(node *Node, peer Peer) {
	conn, err := net.Dial("tcp", peer.GetAddress())
	if err != nil {
		logrus.Errorf("Error connecting to peer: %v", err)
		return
	}

	defer conn.Close()

	msg, err := ParseDataAndEncapsulateSocketMessage("node", node.GetInfo(), ParsePeerToJson)

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

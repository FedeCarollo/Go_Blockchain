package main

import (
	"net"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	PING_INTERVAL = 5 * time.Second
)

func handlePing(t *Tracker, msg *SocketMessage, conn net.Conn) {
	pingingPeer, err := GetParsedDataFromMessage(msg, ParseJsonToPeer)

	if err != nil {
		logrus.Errorf("Error parsing ping message: %v", err)
		return
	}

	t.AddPeer(pingingPeer)

	SendMessage(conn, "pong", t.GetInfo(), ParsePeerToJson)

}

func pingPeers(t *Tracker) {
	ticker := time.NewTicker(PING_INTERVAL)

	for range ticker.C {
		peers := t.GetPeers()

		for _, peer := range peers {
			go pingPeer(peer, t)
		}
	}
}

func pingPeer(peer *Peer, t *Tracker) {
	conn, err := net.Dial("tcp", peer.GetAddress())
	if err != nil {
		logrus.Errorf("Error connecting to peer %v: %v", peer.Id, err)
	}

	err = SendMessage(conn, "ping", t.GetInfo(), ParsePeerToJson)

	if err != nil {
		logrus.Errorf("Error sending ping message to peer %v: %v", peer.Id, err)
		t.RemovePeer(peer)
		return
	}

	defer conn.Close()

	msg, err := ReadMessage(conn)

	if err != nil {
		logrus.Errorf("Error reading message from peer %v: %v", peer.Id, err)
		t.RemovePeer(peer)
		return
	}

	if msg.GetType() != "pong" {
		logrus.Errorf("Expected pong message, got %v", msg.GetType())
		return
	}

}

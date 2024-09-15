package main

import (
	"net"

	"github.com/sirupsen/logrus"
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

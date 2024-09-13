package main

import (
	"encoding/json"
	"net"

	"github.com/sirupsen/logrus"
)

type AnnounceData = Peer // As of now, if new data is added AnnounceData will be redefined

func handleAnnounce(tracker *Tracker, sockMsg *SocketMessage, conn net.Conn) {
	// Parse the data
	newPeer, err := parseAnnounceData(sockMsg.GetData())

	if err != nil {
		logrus.Errorf("Error parsing announce data: %v", err)
		return
	}

	// Add the peer to the tracker
	tracker.AddPeer(newPeer)

	jsonPeers, err := JsonPeers(tracker.GetPeers())

	if err != nil {
		logrus.Errorf("Error converting peers to json: %v", err)
		return
	}

	msg := NewSocketMessage("peerList", jsonPeers)

	tosend, err := msg.ParseToJson()

	if err != nil {
		logrus.Errorf("Error converting message to json: %v", err)
		return
	}

	conn.Write([]byte(tosend + "\n"))

	//TODO: send to the peers the new peer
}

func parseAnnounceData(data string) (*AnnounceData, error) {
	//data is json
	var announceInfo AnnounceData

	err := json.Unmarshal([]byte(data), &announceInfo)

	if err != nil {
		return nil, err
	}

	return &announceInfo, nil
}

func writeToPeer(peer *Peer, data string) error {
	tcpAddr := peer.GetAddress()

	conn, err := net.Dial("tcp", tcpAddr)

	if err != nil {
		logrus.Errorf("Error connecting to peer: %v", err)

		return err
	}

	defer conn.Close()

	_, err = conn.Write([]byte(data))

	if err != nil {
		logrus.Errorf("Error writing to peer: %v", err)

		return err
	}

	return nil
}

func SendUpdatedPeers(tracker *Tracker) {
	jsonPeers, err := JsonPeers(tracker.GetPeers())

	if err != nil {
		logrus.Errorf("Error converting peers to json: %v", err)
		return
	}

	msg := NewSocketMessage("peerList", jsonPeers)

	tosend, err := msg.ParseToJson()

	if err != nil {
		logrus.Errorf("Error converting message to json: %v", err)
		return
	}

	for _, peer := range tracker.GetPeers() {
		writeToPeer(peer, tosend)
	}
}

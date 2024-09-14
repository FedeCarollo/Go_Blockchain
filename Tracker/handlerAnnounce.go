package main

import (
	"encoding/json"
	"net"

	"github.com/sirupsen/logrus"
)

type AnnounceData = Peer // As of now, if new data is added AnnounceData will be redefined

func handleAnnounce(tracker *Tracker, sockMsg *SocketMessage, conn net.Conn) {
	// Parse the data
	newPeer, err := GetParsedDataFromMessage(sockMsg, parseAnnounceData)

	if err != nil {
		logrus.Errorf("Error parsing announce data: %v", err)
		return
	}

	// Add the peer to the tracker
	tracker.AddPeer(newPeer)

	SendMessage(conn, "peerList", tracker.GetPeers(), ParsePeersToJson)

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

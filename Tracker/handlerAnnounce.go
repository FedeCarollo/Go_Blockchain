package main

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

type AnnounceData = Peer // As of now, if new data is added AnnounceData will be redefined

func handleAnnounce(tracker *Tracker, sockMsg *SocketMessage) {
	// Parse the data
	newPeer, err := parseAnnounceData(sockMsg.GetData())

	if err != nil {
		logrus.Errorf("Error parsing announce data: %v", err)
		return
	}

	// Add the peer to the tracker
	tracker.AddPeer(newPeer)

	//TODO: send the peer list to the peer

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

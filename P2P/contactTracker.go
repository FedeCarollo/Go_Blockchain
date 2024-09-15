package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/sirupsen/logrus"
)

func GetTrackersFromFile(path string) ([]Peer, error) {
	trackers := make([]Peer, 0)

	// Read the trackers from the file
	file, err := os.Open(path)

	if err != nil {
		logrus.Errorf("Error opening tracker file: %v", err)
		return nil, err
	}

	err = json.NewDecoder(file).Decode(&trackers)

	if err != nil {
		logrus.Errorf("Error decoding tracker file: %v", err)
		return nil, err
	}

	fmt.Println("Trackers:", trackers)

	return trackers, nil

}

// Node info refer to executor's info
func ContactTrackers(trackers []Peer, node *Node) bool {
	contacted := false //If all trackers are down, return false
	for _, tracker := range trackers {
		if err := ContactTracker(tracker, node); err == nil {
			contacted = true
		}
	}
	return contacted
}

func ContactTracker(tracker Peer, node *Node) error {
	conn, err := net.Dial("tcp", tracker.GetAddress())

	if err != nil {
		logrus.Errorf("Error connecting to tracker: %v", err)
		return err
	}

	defer conn.Close()

	err = SendMessage(conn, "announce", node.GetInfo(), ParsePeerToJson)

	if err != nil {
		logrus.Errorf("Error sending message to tracker: %v", err)
		return err
	}

	peers, err := ReadMessageAndParse(conn, DecodeJsonToPeers)

	if err != nil {
		logrus.Errorf("Error reading peers from tracker: %v", err)
		return err
	}

	node.peers.Append(peers...)
	node.peers.Remove(*node.GetInfo())

	return nil
}

package main

import (
	"net"

	"github.com/sirupsen/logrus"
)

// Node info refer to executor's info
func ContactTrackers(trackers []Peer, node *Node) {
	for _, tracker := range trackers {
		ContactTracker(tracker, node)
	}

}

func ContactTracker(tracker Peer, node *Node) {
	conn, err := net.Dial("tcp", tracker.GetAddress())

	if err != nil {
		logrus.Errorf("Error connecting to tracker: %v", err)
		return
	}

	defer conn.Close()

	// Send the node info to the tracker

	data, err := node.GetInfo().ParseToJson()

	if err != nil {
		logrus.Errorf("Error parsing node info to json: %v", err)
		return
	}

	_, err = conn.Write(data)

	if err != nil {
		logrus.Errorf("Error sending node info to tracker: %v", err)
		return
	}

}

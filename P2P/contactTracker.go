package main

import (
	"bufio"
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

	// Send the node info to the tracker

	data, err := node.GetInfo().ParseToJson()

	if err != nil {
		logrus.Errorf("Error parsing node info to json: %v", err)
		return err
	}
	sockMsg := NewSocketMessage("announce", string(data))

	sockData, err := sockMsg.ParseToJson()

	if err != nil {
		logrus.Errorf("Error parsing node info to json: %v", err)
		return err
	}

	_, err = conn.Write([]byte(sockData + "\n"))

	if err != nil {
		logrus.Errorf("Error sending node info to tracker: %v", err)
		return err
	}

	// Read the response from the tracker
	reader := bufio.NewReader(conn)
	response, err := reader.ReadString('\n')

	if err != nil {
		logrus.Errorf("Error reading response from tracker: %v", err)
		return err
	}

	logrus.Infof("Response from tracker: %v", response)

	var sockRes SocketMessage
	var peers []Peer

	_ = json.Unmarshal([]byte(response), &sockRes)

	//TODO: Check if the response is correct
	err = json.Unmarshal([]byte(sockRes.Data), &peers)
	if err != nil {
		logrus.Errorf("Error decoding peers from tracker: %v", err)
		return err
	}

	node.peers.Append(peers...)

	return nil
}

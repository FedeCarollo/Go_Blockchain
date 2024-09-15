package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/sirupsen/logrus"
)

func main() {
	port := readArgs()
	myInfo := NewPeer("::1", port, IPv6)
	node := NewNode(myInfo)
	createLogger(true)

	trackers, err := GetTrackersFromFile("trackers.json")

	if err != nil {
		log.Fatalf("Error reading trackers: %v", err)
	}

	if !ContactTrackers(trackers, node) {
		log.Fatalf("Could not contact any trackers")
	}
	PingPeers(node)

	// Start the server
	wg := sync.WaitGroup{}
	wg.Add(1)
	go StartServer(node)
	wg.Wait()
}

func createLogger(console bool) {
	if console {
		logrus.SetOutput(os.Stdout)
		return
	}
	//Create logger with logrus
	logFile, err := os.Create("logs/log.txt")

	if err != nil {
		log.Fatal("Cannot create log file", err)
	}

	logrus.SetOutput(logFile)
}

func readArgs() int {
	// Read the arguments
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Usage: ./p2p.exe <port>")
		os.Exit(1)
	}

	port := args[1]

	// Check if the port is valid
	if pInt, err := strconv.Atoi(port); err != nil {
		fmt.Println("Invalid port")
		os.Exit(1)
	} else {
		return pInt
	}
	return -1
}

package main

import (
	"fmt"
	"log"
	"os"
	"p2p_network/client"
	"p2p_network/p2p"
	"p2p_network/p2p/btc_blockchain/blockchain"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/sirupsen/logrus"
)

func main() {
	port := readArgs()
	myInfo := p2p.NewPeer("::1", port, p2p.IPv6)
	node := p2p.NewNode(myInfo)
	createLogger(true)
	trackers, err := p2p.GetTrackersFromFile(filepath.Join(client.GetWd(), "p2p", "trackers.json"))

	if err != nil {
		log.Fatalf("Error reading trackers: %v", err)
	}

	if !p2p.ContactTrackers(trackers, node) {
		log.Fatalf("Could not contact any trackers")
	}
	p2p.PingPeers(node)

	// Start the server
	wg := sync.WaitGroup{}
	wg.Add(2)
	go StartServer(node)
	StartClient(node)
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

func StartServer(node *p2p.Node) {
	err := p2p.StartServer(node)

	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func StartClient(node *p2p.Node) {
	var usr *blockchain.User = nil
	var bchain *blockchain.Blockchain = nil

	for {
		var cmd string
		fmt.Print("Enter command: ")
		fmt.Scanln(&cmd)
		switch cmd {
		case "createUser":
			usr = client.CreateUser()
		case "getBlockchain":
			bchain = client.GetBlockchain(node)
		case "printMySelf":
			if usr != nil {
				client.PrintUser(usr)
			} else {
				fmt.Println("User not created")
			}
		case "peers":
			client.PrintPeers(node)
		case "wallet":
			if usr != nil {
				fmt.Println("Wallet:", usr.GetWallet(bchain))
			} else {
				fmt.Println("User not created")
			}
		case "blocks":
			// blockchain.PrintBlocks()
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Invalid command")
		}
	}
}

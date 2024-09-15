package client

import (
	"fmt"
	"os"
	"p2p_network/p2p"
	"p2p_network/p2p/btc_blockchain/blockchain"
	"path/filepath"
)

func GetWd() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return wd
}

func CreateUser() *blockchain.User {
	usr := blockchain.CreateUserWithParams(true, filepath.Join(GetWd(), "private.key"))
	return &usr
}

func PrintPeers(node *p2p.Node) {
	peers := node.GetPeers()
	for _, peer := range peers {
		fmt.Println(peer.String())
	}
}

func PrintUser(usr *blockchain.User) {
	usr.PrintKeys()
}

func GetBlockchain(node *p2p.Node) *blockchain.Blockchain {

	return nil
}

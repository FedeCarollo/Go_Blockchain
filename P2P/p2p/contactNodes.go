package p2p

import (
	"net"
	"p2p_network/p2p/btc_blockchain/blockchain"
	"sync"
)

func GetBlockchain(node *Node) *blockchain.Blockchain {
	peers := node.GetPeers()
	if len(peers) == 0 {
		//TODO: Ask trackers for stored blockchain
		return nil
	}
	var blockchain *blockchain.Blockchain
	wg := sync.WaitGroup{}
	mut := sync.Mutex{}
	for _, peer := range peers {
		wg.Add(1)
		go GetBlockchainFromPeer(&peer, node, &wg, blockchain, &mut)
	}
	wg.Wait()

	return blockchain
}

func GetBlockchainFromPeer(peer *Peer, node *Node, wg *sync.WaitGroup, blockchain *blockchain.Blockchain, mut *sync.Mutex) {
	defer wg.Done()

	conn, err := net.Dial("tcp", peer.GetAddress())
	if err != nil {
		node.RemovePeer(*peer)
		return
	}

	//Send message to get blockchain
	err = SendMessageNoContent(conn, "get_blockchain")

	if err != nil {
		node.RemovePeer(*peer)
		return
	}

	//Receive blockchain

	bc, err := ReadMessageAndParse(conn, ParseJsonToBlockchain)

	if err != nil {
		bc = nil
		return
	}

	if bc != nil {
		mut.Lock()
		if bc.GetLength() > blockchain.GetLength() {
			*blockchain = *bc
		}
		mut.Unlock()
	}

}

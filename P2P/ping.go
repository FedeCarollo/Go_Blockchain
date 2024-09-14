package main

func PingPeers(node *Node) {
	peers := node.GetPeers()
	for _, peer := range peers {
		go PingPeer(peer)
	}
}

func PingPeer(peer Peer) {

}

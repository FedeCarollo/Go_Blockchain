package p2p

import (
	mapset "github.com/deckarep/golang-set/v2"
)

type Node struct {
	info  *Peer
	peers mapset.Set[Peer]
}

func NewNode(info *Peer) *Node {
	return &Node{
		info:  info,
		peers: mapset.NewSet[Peer](),
	}
}

func (n *Node) AddPeer(peer Peer) {
	n.peers.Add(peer)
}

func (n *Node) RemovePeer(peer Peer) {
	n.peers.Remove(peer)
}

func (n *Node) GetPeers() []Peer {
	return n.peers.ToSlice()
}

func (n *Node) GetInfo() *Peer {
	return n.info
}

func (n *Node) SetInfo(info *Peer) {
	n.info = info
}

func (n *Node) GetPeerCount() int {
	return n.peers.Cardinality()
}

func (n *Node) String() string {
	str := "Node Info:\n"
	str += n.info.String()

	str += "Peers:\n"
	for _, peer := range n.GetPeers() {
		str += peer.String()
	}
	return str
}

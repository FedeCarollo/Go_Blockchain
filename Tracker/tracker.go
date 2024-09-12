package main

import "github.com/hashicorp/go-set/v3"

//region - TrackerInfo
type IpVersion int

const (
	IPv4 IpVersion = 1
	IPv6 IpVersion = 2
)

//endregion

type TrackerInfo = Peer

//#endregion

type Tracker struct {
	info  *Peer
	peers *set.Set[*Peer]
}

func NewTracker(info *TrackerInfo) *Tracker {
	return &Tracker{
		info:  info,
		peers: set.New[*Peer](0),
	}
}

func (t *Tracker) AddPeer(peer *Peer) {
	t.peers.Insert(peer)
}

func (t *Tracker) RemovePeer(peer *Peer) {
	t.peers.Remove(peer)
}

func (t *Tracker) GetPeers() []*Peer {
	return t.peers.Slice()
}

func (t *Tracker) GetInfo() *Peer {
	return t.info
}

func (t *Tracker) SetInfo(info *Peer) {
	t.info = info
}

func (t *Tracker) GetPeerCount() int {
	return t.peers.Size()
}

package main

import (
	mapset "github.com/deckarep/golang-set/v2"
)

//region - TrackerInfo

//endregion

type TrackerInfo = Peer

//#endregion

type Tracker struct {
	info  *Peer
	peers mapset.Set[*Peer]
}

func NewTracker(info *TrackerInfo) *Tracker {
	return &Tracker{
		info:  info,
		peers: mapset.NewSet[*Peer](),
	}
}

func (t *Tracker) AddPeer(peer *Peer) {
	t.peers.Add(peer)
}

func (t *Tracker) RemovePeer(peer *Peer) {
	t.peers.Remove(peer)
}

func (t *Tracker) GetPeers() []*Peer {
	return t.peers.ToSlice()
}

func (t *Tracker) GetInfo() *Peer {
	return t.info
}

func (t *Tracker) SetInfo(info *Peer) {
	t.info = info
}

func (t *Tracker) GetPeerCount() int {
	return t.peers.Cardinality()
}

func (t *Tracker) String() string {
	str := "Tracker Info:\n"
	str += t.info.String()

	str += "Peers:\n"
	for _, peer := range t.GetPeers() {
		str += peer.String()
	}
	return str
}

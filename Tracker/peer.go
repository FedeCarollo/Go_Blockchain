package main

import (
	"encoding/json"
	"strconv"

	"github.com/google/uuid"
)

type IpVersion int

const (
	IPv4 IpVersion = 4
	IPv6 IpVersion = 6
)

type Peer struct {
	Ip        string
	Port      int
	Ipversion IpVersion
	Id        string
}

func NewPeer(ip string, port int, ipversion IpVersion) *Peer {
	return &Peer{
		Ip:        ip,
		Port:      port,
		Ipversion: ipversion,
		Id:        uuid.New().String(),
	}
}

func (p *Peer) GetIp() string {
	return p.Ip
}

func (p *Peer) GetPort() int {
	return p.Port
}

func (p *Peer) GetIpVersion() IpVersion {
	return p.Ipversion
}

func (p *Peer) String() string {
	str := "Peer Info:\n"
	str += "IP: " + p.Ip + "\n"
	str += "Port: " + strconv.Itoa(p.Port) + "\n"
	str += "IP Version: " + p.getIpVersionString() + "\n"
	str += "ID: " + p.Id + "\n"
	return str
}

func (p *Peer) Equals(other *Peer) bool {
	return p.Ip == other.Ip && p.Port == other.Port && p.Ipversion == other.Ipversion
}

func (p *Peer) getIpVersionString() string {
	if p.Ipversion == IPv4 {
		return "IPv4"
	} else {
		return "IPv6"
	}
}

func (p *Peer) GetAddress() (addr string) {
	if p.Ipversion == IPv4 {
		addr = p.Ip + ":" + strconv.Itoa(p.Port)
	} else {
		addr = "[" + p.Ip + "]:" + strconv.Itoa(p.Port)
	}
	return addr
}

func ParsePeersToJson(peers []*Peer) ([]byte, error) {
	data, err := json.Marshal(peers)

	if err != nil {
		return nil, err
	}

	return data, nil

}

func ParsePeerToJson(peer *Peer) ([]byte, error) {
	data, err := json.Marshal(peer)

	if err != nil {
		return nil, err
	}

	return data, nil

}

func ParseJsonToPeer(data string) (*Peer, error) {
	var peer Peer

	err := json.Unmarshal([]byte(data), &peer)

	if err != nil {
		return nil, err
	}

	return &peer, nil
}

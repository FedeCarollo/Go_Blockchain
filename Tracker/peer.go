package main

import "strconv"

type IpVersion int

const (
	IPv4 IpVersion = 1
	IPv6 IpVersion = 2
)

type Peer struct {
	Ip        string
	Port      int
	Ipversion IpVersion
}

func NewPeer(ip string, port int, ipversion IpVersion) *Peer {
	return &Peer{
		Ip:        ip,
		Port:      port,
		Ipversion: ipversion,
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

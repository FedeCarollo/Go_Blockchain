package main

import "strconv"

type Peer struct {
	ip        string
	port      int
	ipversion IpVersion
}

func NewPeer(ip string, port int, ipversion IpVersion) *Peer {
	return &Peer{
		ip:        ip,
		port:      port,
		ipversion: ipversion,
	}
}

func (p *Peer) GetIp() string {
	return p.ip
}

func (p *Peer) GetPort() int {
	return p.port
}

func (p *Peer) GetIpVersion() IpVersion {
	return p.ipversion
}

func (p *Peer) String() string {
	str := "Peer Info:\n"
	str += "IP: " + p.ip + "\n"
	str += "Port: " + strconv.Itoa(p.port) + "\n"
	str += "IP Version: " + p.getIpVersionString() + "\n"

	return str
}

func (p *Peer) Equals(other *Peer) bool {
	return p.ip == other.ip && p.port == other.port && p.ipversion == other.ipversion
}

func (p *Peer) getIpVersionString() string {
	if p.ipversion == IPv4 {
		return "IPv4"
	} else {
		return "IPv6"
	}
}

func (p *Peer) GetAddress() (addr string) {
	if p.ipversion == IPv4 {
		addr = p.ip + ":" + strconv.Itoa(p.port)
	} else {
		addr = "[" + p.ip + "]:" + strconv.Itoa(p.port)
	}
	return addr
}

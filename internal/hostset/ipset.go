package hostset

import (
	"fmt"
	"net"
)

type IpSet interface {
	Add(ip string) error
	Contains(ip net.IP) bool
}

type IP4 [4]byte
type IP6 [16]byte

type InMemoryIpSet struct {
	ip4Addresses map[IP4]struct{}
	ip6Addresses map[IP6]struct{}
	ip4Nets      []*net.IPNet
	ip6Nets      []*net.IPNet
}

func NewInMemoryIpSet() *InMemoryIpSet {
	return &InMemoryIpSet{
		ip4Addresses: make(map[IP4]struct{}),
		ip6Addresses: make(map[IP6]struct{}),
		ip4Nets:      make([]*net.IPNet, 0),
		ip6Nets:      make([]*net.IPNet, 0),
	}
}

func (s *InMemoryIpSet) Add(ip string) error {
	_, ipNet, err := net.ParseCIDR(ip)
	if err == nil {
		s.AddSubNet(ipNet)
		return nil
	}
	addr := net.ParseIP(ip)
	if addr != nil {
		s.AddAddress(addr)
		return nil
	}
	return fmt.Errorf("invalid ip or network: %s", ip)
}

func (s *InMemoryIpSet) AddAddress(ip net.IP) {
	if addr := ip.To4(); addr != nil {
		s.ip4Addresses[IP4(addr)] = struct{}{}
		return
	}
	addr := ip.To16()
	s.ip6Addresses[IP6(addr)] = struct{}{}
}

func (s *InMemoryIpSet) AddSubNet(ipNet *net.IPNet) {
	if ipNet.IP.To4() != nil {
		s.ip4Nets = append(s.ip4Nets, ipNet)
		return
	}
	s.ip6Nets = append(s.ip6Nets, ipNet)
}

func (s *InMemoryIpSet) Contains(ip net.IP) bool {
	if addr := ip.To4(); addr != nil {
		return s.contain4(addr)
	}
	return s.contain6(ip)
}

func (s *InMemoryIpSet) contain4(ip net.IP) bool {
	_, ok := s.ip4Addresses[IP4(ip)]
	if ok {
		return true
	}
	for _, ipNet := range s.ip4Nets {
		if ipNet.Contains(ip) {
			return true
		}
	}
	return false
}

func (s *InMemoryIpSet) contain6(ip net.IP) bool {
	_, ok := s.ip6Addresses[IP6(ip)]
	if ok {
		return true
	}
	for _, ipNet := range s.ip6Nets {
		if ipNet.Contains(ip) {
			return true
		}
	}
	return false
}

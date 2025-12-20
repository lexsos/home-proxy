package hostset

import (
	"fmt"
	"net"
)

type IpSet struct {
	ip4Nets      []*net.IPNet
	ip4Addresses []net.IP
	ip6Nets      []*net.IPNet
	ip6Addresses []net.IP
}

func NewIpSet() *IpSet {
	return &IpSet{}
}

func (s *IpSet) Add(ip string) error {
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

func (s *IpSet) AddAddress(ip net.IP) {
	if addr := ip.To4(); addr != nil {
		s.ip4Addresses = append(s.ip4Addresses, addr)
		return
	}
	s.ip6Addresses = append(s.ip6Addresses, ip)
}

func (s *IpSet) AddSubNet(ipNet *net.IPNet) {
	if ipNet.IP.To4() != nil {
		s.ip4Nets = append(s.ip4Nets, ipNet)
		return
	}
	s.ip6Nets = append(s.ip6Nets, ipNet)
}

func (s *IpSet) Contains(ip net.IP) bool {
	if addr := ip.To4(); addr != nil {
		return s.contain4(addr)
	}
	return s.contain6(ip)
}

func (s *IpSet) contain4(ip net.IP) bool {
	for _, addr := range s.ip4Addresses {
		if addr.Equal(ip) {
			return true
		}
	}
	for _, ipNet := range s.ip4Nets {
		if ipNet.Contains(ip) {
			return true
		}
	}
	return false
}

func (s *IpSet) contain6(ip net.IP) bool {
	for _, addr := range s.ip6Addresses {
		if addr.Equal(ip) {
			return true
		}
	}
	for _, ipNet := range s.ip6Nets {
		if ipNet.Contains(ip) {
			return true
		}
	}
	return false
}

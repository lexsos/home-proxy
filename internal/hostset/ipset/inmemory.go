package ipset

import (
	"fmt"
	"net"
)

type InMemoryIpSet struct {
	ip4Nets map[IP4]struct{}
	ip6Nets map[IP6]struct{}
	lens4   map[int]struct{}
	lens6   map[int]struct{}
}

func NewInMemoryIpSet() *InMemoryIpSet {
	return &InMemoryIpSet{
		ip4Nets: make(map[IP4]struct{}),
		ip6Nets: make(map[IP6]struct{}),
		lens4:   make(map[int]struct{}),
		lens6:   make(map[int]struct{}),
	}
}

func (s *InMemoryIpSet) Add(ip string) error {
	_, ipNet, err := net.ParseCIDR(ip)
	if err == nil {
		return s.addSubNet(ipNet)
	}
	addr := net.ParseIP(ip)
	if addr != nil {
		return s.addAddress(addr)
	}
	return fmt.Errorf("invalid ip or network: %s", ip)
}

func (s *InMemoryIpSet) addAddress(ip net.IP) error {
	if addr4 := toIP4(ip); addr4 != nil {
		s.ip4Nets[*addr4] = struct{}{}
		s.lens4[32] = struct{}{}
		return nil
	}
	if addr6 := toIP6(ip); addr6 != nil {
		s.ip6Nets[*addr6] = struct{}{}
		s.lens6[128] = struct{}{}
		return nil
	}
	return fmt.Errorf("invalid ip: %s", ip)
}

func (s *InMemoryIpSet) addSubNet(ipNet *net.IPNet) error {
	maskLen, bits := ipNet.Mask.Size()
	if bits == 32 {
		masked := MaskIp4(ipNet.IP, maskLen)
		s.ip4Nets[*masked] = struct{}{}
		s.lens4[maskLen] = struct{}{}
		return nil
	}
	if bits == 128 {
		masked := MaskIp6(ipNet.IP, maskLen)
		s.ip6Nets[*masked] = struct{}{}
		s.lens6[maskLen] = struct{}{}
		return nil
	}
	return fmt.Errorf("invalid ip net: %s", ipNet)
}

func (s *InMemoryIpSet) Contains(ip net.IP) (bool, error) {
	sig, err := NewIpSignature(ip)
	if err != nil {
		return false, fmt.Errorf("invalid ip: %s", ip)
	}
	return s.ContainsSig(sig)
}

func (s *InMemoryIpSet) ContainsSig(sig *IpSignature) (bool, error) {
	if sig.Is4() {
		for maskLen := range s.lens4 {
			masked, err := sig.GetForMask4(maskLen)
			if err != nil {
				return false, fmt.Errorf("invalid mask len for ip4: %d", maskLen)
			}
			_, ok := s.ip4Nets[masked]
			if ok {
				return true, nil
			}
		}
		return false, nil
	}
	if sig.Is6() {
		for maskLen := range s.lens6 {
			masked, err := sig.GetForMask6(maskLen)
			if err != nil {
				return false, fmt.Errorf("invalid mask len for ip6: %d", maskLen)
			}
			_, ok := s.ip6Nets[masked]
			if ok {
				return true, nil
			}
		}
	}
	return false, fmt.Errorf("invalid ip: %s", sig.Src())
}

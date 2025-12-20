package ipset

import (
	"fmt"
	"net"
)

type IpVersin int

const (
	IPv4 IpVersin = 4
	IPv6 IpVersin = 6
)

type IpSignature struct {
	src     net.IP
	version IpVersin
	ip4     IP4
	ip6     IP6
	nets4   map[int]IP4
	nets6   map[int]IP6
}

func NewIpSignature(ip net.IP) (*IpSignature, error) {
	if addr4 := toIP4(ip); addr4 != nil {
		return &IpSignature{
			src:     ip,
			version: IPv4,
			ip4:     *addr4,
			nets4:   make(map[int]IP4),
			nets6:   make(map[int]IP6),
		}, nil
	}
	if addr6 := toIP6(ip); addr6 != nil {
		return &IpSignature{
			src:     ip,
			version: IPv6,
			ip6:     *addr6,
			nets6:   make(map[int]IP6),
			nets4:   make(map[int]IP4),
		}, nil
	}
	return nil, fmt.Errorf("can't create ip signature: invalid ip: %s", ip)
}

func (s *IpSignature) GetForMask4(maskLen int) (IP4, error) {
	masked, ok := s.nets4[maskLen]
	if ok {
		return masked, nil
	}
	if s.version != IPv4 {
		return IP4{}, fmt.Errorf("can't get ip for mask4: invalid ip version: %d", s.version)
	}
	if maskLen > 32 || maskLen < 0 {
		return IP4{}, fmt.Errorf("can't get ip for mask4: invalid mask len: %d", maskLen)
	}
	mask := net.CIDRMask(maskLen, 32)
	masked = *toIP4(s.src.Mask(mask))
	s.nets4[maskLen] = masked
	return masked, nil
}

func (s *IpSignature) GetForMask6(maskLen int) (IP6, error) {
	masked, ok := s.nets6[maskLen]
	if ok {
		return masked, nil
	}
	if s.version != IPv6 {
		return IP6{}, fmt.Errorf("can't get ip for mask6: invalid ip version: %d", s.version)
	}
	if maskLen > 128 || maskLen < 0 {
		return IP6{}, fmt.Errorf("can't get ip for mask6: invalid mask len: %d", maskLen)
	}
	mask := net.CIDRMask(maskLen, 128)
	masked = *toIP6(s.src.Mask(mask))
	s.nets6[maskLen] = masked
	return masked, nil
}

func (s *IpSignature) Is4() bool {
	return s.version == IPv4
}

func (s *IpSignature) Is6() bool {
	return s.version == IPv6
}

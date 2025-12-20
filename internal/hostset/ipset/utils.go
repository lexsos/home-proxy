package ipset

import (
	"net"
)

func toIP4(ip net.IP) *IP4 {
	addr4 := ip.To4()
	if addr4 == nil {
		return nil
	}
	return (*IP4)(addr4)
}

func toIP6(ip net.IP) *IP6 {
	addr6 := ip.To16()
	if addr6 == nil {
		return nil
	}
	return (*IP6)(addr6)
}

package ipset

import (
	"net"
)

type IpSet interface {
	Contains(ip net.IP) (bool, error)
	ContainsSig(sig *IpSignature) (bool, error)
}

type IP4 [4]byte
type IP6 [16]byte

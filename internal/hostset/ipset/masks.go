package ipset

import (
	"net"
)

var (
	masks4 = map[int]net.IPMask{}
	masks6 = map[int]net.IPMask{}
)

func init() {
	for i := 0; i <= 32; i++ {
		masks4[i] = net.CIDRMask(i, 32)
	}
	for i := 0; i <= 128; i++ {
		masks6[i] = net.CIDRMask(i, 128)
	}
}

func MaskIp4(ip net.IP, maskLen int) *IP4 {
	if maskLen > 32 || maskLen < 0 {
		return nil
	}
	maked := ip.Mask(masks4[maskLen])
	return toIP4(maked)
}

func MaskIp6(ip net.IP, maskLen int) *IP6 {
	if maskLen > 128 || maskLen < 0 {
		return nil
	}
	maked := ip.Mask(masks6[maskLen])
	return toIP6(maked)
}

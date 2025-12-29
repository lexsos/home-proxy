package hostset

import (
	"github.com/lexsos/home-proxy/internal/hostset/domainset"
	"github.com/lexsos/home-proxy/internal/hostset/ipset"
)

type HostSet struct {
	domains domainset.DomainSet
	ips     ipset.IpSet
}

func NewHostSet(domains domainset.DomainSet, ips ipset.IpSet) *HostSet {
	return &HostSet{
		domains: domains,
		ips:     ips,
	}
}

func (h *HostSet) ContainsIpSig(ipSig *ipset.IpSignature) (bool, error) {
	return h.ips.ContainsSig(ipSig)
}

func (h *HostSet) ContainsDomainSig(domainSig *domainset.DomainSig) (bool, error) {
	return h.domains.ContainsSig(domainSig)
}

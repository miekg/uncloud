package docker

import (
	"fmt"
	"net"
	"net/netip"
)

// addrOfPrefix checks the interface and returns the address of each interface that container the prefix.
func addrOfPrefix(prefix netip.Prefix) ([]string, error) {
	ifis, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var addrs []string
	for _, ifi := range ifis {
		ifaddrs, _ := ifi.Addrs()
		for _, addr := range ifaddrs {
			ip, err := netip.ParseAddr(addr.String())
			if err != nil {
				continue
			}
			if prefix.Contains(ip) {
				addrs = append(addrs, ip.String())
			}
		}
	}
	if len(addrs) == 0 {
		return nil, fmt.Errorf("no addresses are contained in prefix '%s'", prefix)
	}

	return addrs, nil
}

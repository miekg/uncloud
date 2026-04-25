package docker

import (
	"fmt"
	"net"
)

func addrsFromInterface(iface net.Interface) ([]string, error) {
	ifis, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var addrs []net.Addr
	ok := false
	for _, ifi := range ifis {
		if ifi.Name == iface.Name {
			ok = true
			addrs, _ = ifi.Addrs()
		}
	}
	if !ok {
		return nil, fmt.Errorf("interface name '%s' not found", iface.Name)
	}
	if len(addrs) == 0 {
		return nil, fmt.Errorf("interface '%s' has no addresses", iface.Name)
	}

	ips := []string{}
	for _, addr := range addrs {
		ipnet, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipnet.IP.IsLoopback() || ipnet.IP.IsPrivate() {
			continue
		}
		if ipnet.IP.IsLinkLocalUnicast() || ipnet.IP.IsUnspecified() {
			continue
		}
		ips = append(ips, ipnet.IP.String())
	}
	return ips, nil
}

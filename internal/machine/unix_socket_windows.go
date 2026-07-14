package machine

import "net"

// listenUnixSocket is not supported on Windows.
func listenUnixSocket(path string) (net.Listener, error) {
	return nil, nil
}

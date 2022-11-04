//go:build darwin

package tproxy

import "net"

// These methods are here just for ease of development on darwin platforms

func ListenTCP(network string, laddr *net.TCPAddr) (net.Listener, error) {
	return nil, nil
}

func ListenUDP(network string, laddr *net.UDPAddr) (*net.UDPConn, error) {
	return nil, nil
}

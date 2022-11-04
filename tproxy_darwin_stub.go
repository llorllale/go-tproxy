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

func ReadFromUDP(conn *net.UDPConn, b []byte) (int, *net.UDPAddr, *net.UDPAddr, error) {
	return 0, nil, nil, nil
}

func DialUDP(network string, laddr *net.UDPAddr, raddr *net.UDPAddr) (*net.UDPConn, error) {
	return nil, nil
}

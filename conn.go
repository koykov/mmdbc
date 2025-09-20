package mmdbc

import "net/netip"

type Connection interface {
	Get(ip netip.Addr) (*Tuple, error)
	Gets(ip string) (*Tuple, error)
	PGet(dst *Tuple, ip netip.Addr) error
	PGets(dst *Tuple, ip string) error
}

func Connect(filePath string) (Connection, error) {
	// todo implement me
	return nil, nil
}

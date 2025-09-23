package mmdbc

import (
	"io"
	"net/netip"
	"os"
)

type Connection interface {
	Get(ip netip.Addr) (*Tuple, error)
	Gets(ip string) (*Tuple, error)
	PGet(dst *Tuple, ip netip.Addr) error
	PGets(dst *Tuple, ip string) error
	io.Closer
}

func Connect(filePath string) (_ Connection, err error) {
	var c conn
	var f *os.File
	if f, err = os.Open(filePath); err != nil {
		return
	}
	fi, err := f.Stat()
	if err != nil {
		return
	}
	c.buf = make([]byte, fi.Size())
	if _, err = io.ReadFull(f, c.buf); err != nil {
		return
	}
	// todo read meta
	return &c, nil
}

type conn struct {
	buf  []byte
	meta Meta
}

func (c *conn) Get(ip netip.Addr) (*Tuple, error) {
	// todo implement me
	return nil, nil
}

func (c *conn) Gets(ip string) (*Tuple, error) {
	// todo implement me
	return nil, nil
}

func (c *conn) PGet(dst *Tuple, ip netip.Addr) error {
	// todo implement me
	return nil
}

func (c *conn) PGets(dst *Tuple, ip string) error {
	// todo implement me
	return nil
}

func (c *conn) Close() error {
	c.meta.reset()
	return nil
}

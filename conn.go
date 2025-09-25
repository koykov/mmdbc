package mmdbc

import (
	"bytes"
	"io"
	"net/netip"
	"os"
)

const metaPrefix = "\xAB\xCD\xEFMaxMind.com"

type Connection interface {
	Get(ip netip.Addr) (*Tuple, error)
	Gets(ip string) (*Tuple, error)
	PGet(dst *Tuple, ip netip.Addr) error
	PGets(dst *Tuple, ip string) error
	io.Closer
}

func Connect(filePath string) (c Connection, err error) {
	var f *os.File
	if f, err = os.Open(filePath); err != nil {
		return
	}
	fi, err := f.Stat()
	if err != nil {
		return
	}
	cn := &conn{buf: make([]byte, fi.Size())}
	if _, err = io.ReadFull(f, cn.buf); err != nil {
		return
	}
	i := bytes.LastIndex(cn.buf, []byte(metaPrefix))
	if i == -1 {
		err = ErrMetaNotFound
		return
	}
	cn.bufm = cn.buf[i+len(metaPrefix):]
	if err = cn.decodeMeta(); err != nil {
		return
	}
	c = cn
	return
}

type conn struct {
	buf  []byte
	bufm []byte
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

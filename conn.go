package mmdbcli

import (
	"bytes"
	"context"
	"io"
	"net/netip"
	"os"
)

const metaPrefix = "\xAB\xCD\xEFMaxMind.com"

type Connection interface {
	Meta() *Meta
	Get(ctx context.Context, ip netip.Addr) (*Tuple, error)
	Gets(ctx context.Context, ip string) (*Tuple, error)
	PGet(ctx context.Context, dst *Tuple, ip netip.Addr) error
	PGets(ctx context.Context, dst *Tuple, ip string) error
	EachNetwork(ctx context.Context, fn func(*Tuple) error) error
	EachNetworkWithOptions(ctx context.Context, fn func(*Tuple) error, options NetworkOption) error
	io.Closer
}

func Connect(filePath string) (c Connection, err error) {
	return connect(filePath)
}

func connect(filePath string) (c *conn, err error) {
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

	cn.nodeoff = cn.meta.recSize / 4
	switch cn.meta.ipVer {
	case 4:
		cn.ipv4bits = 96
	case 6:
		var i, node uint64
		for i = 0; i < 96 && node < cn.meta.nodec; i++ {
			node, err = cn.getNode(node*cn.nodeoff, 0)
			if err != nil {
				return nil, err
			}
		}
		cn.ipv4off, cn.ipv4bits = node, i
	default:
		return nil, ErrMetaIpVersion
	}

	c = cn
	return
}

type conn struct {
	buf      []byte
	bufm     []byte
	meta     Meta
	nodeoff  uint64
	ipv4off  uint64
	ipv4bits uint64

	trvrsNextFn func(ctx context.Context, c *conn, ip *netip.Addr, node, bit uint64, stopbit int) (uint64, uint64, error)
}

func (c *conn) Meta() *Meta {
	return &c.meta
}

func (c *conn) Get(ctx context.Context, ip netip.Addr) (*Tuple, error) {
	_, _ = ctx, ip
	// todo implement me
	return nil, nil
}

func (c *conn) Gets(ctx context.Context, ip string) (*Tuple, error) {
	ip_, err := netip.ParseAddr(ip)
	if err != nil {
		return nil, err
	}
	return c.Get(ctx, ip_)
}

func (c *conn) PGet(ctx context.Context, dst *Tuple, ip netip.Addr) error {
	_, _, _ = ctx, dst, ip
	// todo implement me
	return nil
}

func (c *conn) PGets(ctx context.Context, dst *Tuple, ip string) error {
	ip_, err := netip.ParseAddr(ip)
	if err != nil {
		return err
	}
	return c.PGet(ctx, dst, ip_)
}

func (c *conn) Close() error {
	c.meta.reset()
	return nil
}

func (c *conn) getNode(off, bit uint64) (uint64, error) {
	switch c.meta.recSize {
	case 24:
		off += bit * 3
		return (uint64(c.buf[off]) << 16) | (uint64(c.buf[off+1]) << 8) | uint64(c.buf[off+2]), nil
	case 28:
		if bit == 0 {
			return ((uint64(c.buf[off+3]) & 0xF0) << 20) | (uint64(c.buf[off]) << 16) | (uint64(c.buf[off+1]) << 8) | uint64(c.buf[off+2]), nil
		}
		return ((uint64(c.buf[off+3]) & 0x0F) << 24) | (uint64(c.buf[off+4]) << 16) | (uint64(c.buf[off+5]) << 8) | uint64(c.buf[off+6]), nil
	case 32:
		off += bit * 4
		return (uint64(c.buf[off]) << 24) | (uint64(c.buf[off+1]) << 16) | (uint64(c.buf[off+2]) << 8) | uint64(c.buf[off+3]), nil
	default:
		return 0, ErrBadRecordSize
	}
}

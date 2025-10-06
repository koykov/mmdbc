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
	c = cn
	return
}

type conn struct {
	buf  []byte
	bufm []byte
	meta Meta
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
	_, _ = ctx, ip
	// todo implement me
	return nil, nil
}

func (c *conn) PGet(ctx context.Context, dst *Tuple, ip netip.Addr) error {
	_, _, _ = ctx, dst, ip
	// todo implement me
	return nil
}

func (c *conn) PGets(ctx context.Context, dst *Tuple, ip string) error {
	_, _, _ = ctx, dst, ip
	// todo implement me
	return nil
}

func (c *conn) Close() error {
	c.meta.reset()
	return nil
}

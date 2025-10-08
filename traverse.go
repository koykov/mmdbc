package mmdbcli

import (
	"context"
	"io"
	"net/netip"
)

func (c *conn) traverse(ctx context.Context, ip *netip.Addr, node uint64, bits int) (uint64, uint64, error) {
	var off uint64
	if ip.Is4() {
		off, node = c.ipv4bits, c.ipv4off
	}
	var err error
	raw := ip.As16()
	for ; off < uint64(bits) && node < c.meta.nodec; off++ {
		select {
		case <-ctx.Done():
			return 0, 0, ctx.Err()
		default:
			idx := off >> 3
			pos := 7 - (off & 7)
			bit := (uint64(raw[idx]) >> pos) & 1
			if node, err = c.trvrsNextFn(c, node, bit); err != nil {
				return 0, 0, err
			}
		}
	}
	return node, off, nil
}

func traverse24(c *conn, node, bit uint64) (uint64, error) {
	i := node * 6
	off := i + bit*3
	if off > uint64(len(c.buf))-3 {
		return 0, io.ErrUnexpectedEOF
	}
	node = (uint64(c.buf[off]) << 16) | (uint64(c.buf[off+1]) << 8) | uint64(c.buf[off+2])
	return node, nil
}

func traverse28(c *conn, node, bit uint64) (uint64, error) {
	i := node * 7
	off := i + bit*4
	if off > uint64(len(c.buf))-3 {
		return 0, io.ErrUnexpectedEOF
	}
	mask := uint64(0xf0 >> (bit * 4))
	sh := bit*4 + 20
	node = ((uint64(c.buf[off+3]) & mask) << sh) | (uint64(c.buf[off]) << 16) | (uint64(c.buf[off+1]) << 8) | uint64(c.buf[off+2])
	return node, nil
}

func traverse32(c *conn, node, bit uint64) (uint64, error) {
	i := node * 8
	off := i + bit*4
	if off > uint64(len(c.buf))-4 {
		return 0, io.ErrUnexpectedEOF
	}
	node = (uint64(c.buf[off]) << 24) | (uint64(c.buf[off+1]) << 16) | (uint64(c.buf[off+2]) << 8) | uint64(c.buf[off+3])
	return node, nil
}

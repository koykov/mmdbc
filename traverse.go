package mmdbcli

import (
	"context"
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
		i := off >> 3
		pos := 7 - (i & 7)
		bit := (uint64(raw[i]) >> pos) & 1
		if node, _, err = c.trvrsNextFn(ctx, c, ip, node, bit, bits); err != nil {
			return 0, 0, err
		}
	}
	return node, off, nil
}

func traverse24(ctx context.Context, c *conn, ip *netip.Addr, node, bit uint64, stopbit int) (uint64, uint64, error) {
	// todo implement me
	return 0, 0, nil
}

func traverse28(ctx context.Context, c *conn, ip *netip.Addr, node, bit uint64, stopbit int) (uint64, uint64, error) {
	// todo implement me
	return 0, 0, nil
}

func traverse32(ctx context.Context, c *conn, ip *netip.Addr, node, bit uint64, stopbit int) (uint64, uint64, error) {
	// todo implement me
	return 0, 0, nil
}

package mmdbcli

import (
	"context"
	"net/netip"
)

type NetworkOption uint64

const (
	NetworkOptionIncludeAliased NetworkOption = 1 << iota
	NetworkOptionIncludeEmptyNetwork
	NetworkOptionSkipEmptyTuple
)

const (
	NetworkOptionNoOptions NetworkOption = 0
	NetworkOptionAll                     = NetworkOptionIncludeAliased | NetworkOptionIncludeEmptyNetwork | NetworkOptionSkipEmptyTuple
)

func (c *conn) EachNetwork(ctx context.Context, fn func(*Record) error) error {
	return c.EachNetworkWithOptions(ctx, fn, NetworkOptionNoOptions)
}

func (c *conn) EachNetworkWithOptions(ctx context.Context, fn func(*Record) error, options NetworkOption) error {
	pfx := allpfx[c.meta.ipVer]
	if !pfx.IsValid() {
		return ErrInvalidPrefix
	}
	if c.meta.ipVer == 4 && pfx.Addr().Is6() {
		return ErrOverflowPrefix
	}

	addrRaw, bits := pfx.Addr(), pfx.Bits()
	addr := addrRaw
	if addrRaw.Is4() {
		raw := addrRaw.As4()
		var raw6 [16]byte
		copy(raw6[12:], raw[:])
		addrRaw = netip.AddrFrom16(raw6)
		bits += 96
	}
	if bits > 128 {
		return ErrOverflowIPv6
	}

	root, bit, err := c.traverse(ctx, &addrRaw, 0, bits)
	if err != nil {
		return err
	}
	if pfx, err = addr.Prefix(int(bit)); err != nil {
		return err
	}

	pfxAddr := pfx.Addr()
	if err = c.netwalk(ctx, root, &pfxAddr, fn, options, 0); err != nil {
		return err
	}

	return nil
}

func (c *conn) netwalk(ctx context.Context, root uint64, addr *netip.Addr, fn func(*Record) error, options NetworkOption, depth int) error {
	for {
		if root == c.meta.nodec {
			if options&NetworkOptionIncludeEmptyNetwork != 0 {
				_ = fn(&Record{})
			}
			break
		}
		if root == c.ipv4off && options&NetworkOptionIncludeAliased == 0 && !addr.Is4() {
			break
		}
		if root > c.meta.nodec {
			// todo complete me
		}
	}
	return nil
}

func (c *conn) eachNetwork(ctx context.Context, addr *netip.Addr, bits int, fn func(*Record) error, options NetworkOption) error {
	// todo implement me
	return nil
}

var allpfx = [7]netip.Prefix{
	{},
	{},
	{},
	{},
	netip.MustParsePrefix("0.0.0.0/0"),
	{},
	netip.MustParsePrefix("::/0"),
}

var _ = NetworkOptionAll

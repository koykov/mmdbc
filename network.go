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

func (c *conn) EachNetwork(ctx context.Context, fn func(*Tuple) error) error {
	return c.EachNetworkWithOptions(ctx, fn, NetworkOptionNoOptions)
}

func (c *conn) EachNetworkWithOptions(ctx context.Context, fn func(*Tuple) error, options NetworkOption) error {
	pfx := allpfx[c.meta.ipVer]
	if !pfx.IsValid() {
		return ErrInvalidPrefix
	}
	if c.meta.ipVer == 4 && pfx.Addr().Is6() {
		return ErrOverflowPrefix
	}

	addr, bits := pfx.Addr(), pfx.Bits()
	if addr.Is4() {
		raw := addr.As4()
		var raw6 [16]byte
		copy(raw6[12:], raw[:])
		addr = netip.AddrFrom16(raw6)
		bits += 96
	}
	if bits > 128 {
		return ErrOverflowIPv6
	}

	return c.eachNetwork(ctx, &addr, pfx.Bits(), fn, options)
}

func (c *conn) eachNetwork(ctx context.Context, addr *netip.Addr, bits int, fn func(*Tuple) error, options NetworkOption) error {
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

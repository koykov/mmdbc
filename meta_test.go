package mmdbc

import (
	"fmt"
	"testing"
)

func TestMeta(t *testing.T) {
	type meta struct {
		fname string
	}
	tcs := []meta{
		{fname: "GeoIP2-Anonymous-IP-Test"},
		{fname: "GeoIP2-City-Shield-Test"},
		{fname: "GeoIP2-City-Test-Broken-Double-Format"},
		{fname: "GeoIP2-City-Test-Invalid-Node-Count"},
		{fname: "GeoIP2-City-Test"},
		{fname: "GeoIP2-Connection-Type-Test"},
		{fname: "GeoIP2-Country-Shield-Test"},
		{fname: "GeoIP2-Country-Test"},
		{fname: "GeoIP2-DensityIncome-Test"},
		{fname: "GeoIP2-Domain-Test"},
		{fname: "GeoIP2-Enterprise-Shield-Test"},
		{fname: "GeoIP2-Enterprise-Test"},
		{fname: "GeoIP2-IP-Risk-Test"},
		{fname: "GeoIP2-ISP-Test"},
		{fname: "GeoIP2-Precision-Enterprise-Shield-Test"},
		{fname: "GeoIP2-Precision-Enterprise-Test"},
		{fname: "GeoIP2-Static-IP-Score-Test"},
		{fname: "GeoIP2-User-Count-Test"},
		{fname: "GeoIP-Anonymous-Plus"},
		{fname: "GeoIP-Anonymous-Plus-Test"},
		{fname: "GeoLite2-ASN-Test"},
		{fname: "GeoLite2-City-Test"},
		{fname: "GeoLite2-Country-Test"},
		{fname: "MaxMind-DB-no-ipv4-search-tree"},
		{fname: "MaxMind-DB-string-value-entries"},
		{fname: "MaxMind-DB-test-broken-pointers-24"},
		{fname: "MaxMind-DB-test-broken-search-tree-24"},
		{fname: "MaxMind-DB-test-decoder"},
		{fname: "MaxMind-DB-test-ipv4-24"},
		{fname: "MaxMind-DB-test-ipv4-28"},
		{fname: "MaxMind-DB-test-ipv4-32"},
		{fname: "MaxMind-DB-test-ipv6-24"},
		{fname: "MaxMind-DB-test-ipv6-28"},
		{fname: "MaxMind-DB-test-ipv6-32"},
		{fname: "MaxMind-DB-test-metadata-pointers"},
		{fname: "MaxMind-DB-test-mixed-24"},
		{fname: "MaxMind-DB-test-mixed-28"},
		{fname: "MaxMind-DB-test-mixed-32"},
		{fname: "MaxMind-DB-test-nested"},
		{fname: "MaxMind-DB-test-pointer-decoder"},
	}
	for _, tc := range tcs {
		t.Run(tc.fname, func(t *testing.T) {
			fpath := fmt.Sprintf("/tmp/mmdb-test/%s.mmdb", tc.fname)
			cn, err := Connect(fpath)
			if err != nil {
				t.Fatal(err)
			}
			defer func() { _ = cn.Close() }()
			_ = cn
			// todo check result
		})
	}
}

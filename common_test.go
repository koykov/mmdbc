package mmdbc

import (
	"bufio"
	"errors"
	"io"
	"net/http"
	"os"
	"testing"
)

func init() {
	dbs := map[string]string{
		"GeoIP2-Anonymous-IP-Test.mmdb":                "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/GeoIP2-Anonymous-IP-Test.mmdb",
		"GeoIP2-City-Shield-Test.mmdb":                 "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/GeoIP2-City-Shield-Test.mmdb",
		"GeoIP2-City-Test-Broken-Double-Format.mmdb":   "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/GeoIP2-City-Test-Broken-Double-Format.mmdb",
		"GeoIP2-City-Test-Invalid-Node-Count.mmdb":     "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/GeoIP2-City-Test-Invalid-Node-Count.mmdb",
		"GeoIP2-City-Test.mmdb":                        "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/GeoIP2-City-Test.mmdb",
		"GeoIP2-Connection-Type-Test.mmdb":             "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/GeoIP2-Connection-Type-Test.mmdb",
		"GeoIP2-Country-Shield-Test.mmdb":              "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/GeoIP2-Country-Shield-Test.mmdb",
		"GeoIP2-Country-Test.mmdb":                     "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/GeoIP2-Country-Test.mmdb",
		"GeoIP2-DensityIncome-Test.mmdb":               "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/GeoIP2-DensityIncome-Test.mmdb",
		"GeoIP2-Domain-Test.mmdb":                      "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/GeoIP2-Domain-Test.mmdb",
		"GeoIP2-Enterprise-Shield-Test.mmdb":           "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/GeoIP2-Enterprise-Shield-Test.mmdb",
		"GeoIP2-Enterprise-Test.mmdb":                  "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/GeoIP2-Enterprise-Test.mmdb",
		"GeoIP2-IP-Risk-Test.mmdb":                     "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/GeoIP2-IP-Risk-Test.mmdb",
		"GeoIP2-ISP-Test.mmdb":                         "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/GeoIP2-ISP-Test.mmdb",
		"GeoIP2-Precision-Enterprise-Shield-Test.mmdb": "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/GeoIP2-Precision-Enterprise-Shield-Test.mmdb",
		"GeoIP2-Precision-Enterprise-Test.mmdb":        "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/GeoIP2-Precision-Enterprise-Test.mmdb",
		"GeoIP2-Static-IP-Score-Test.mmdb":             "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/GeoIP2-Static-IP-Score-Test.mmdb",
		"GeoIP2-User-Count-Test.mmdb":                  "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/GeoIP2-User-Count-Test.mmdb",
		"GeoIP-Anonymous-Plus.mmdb":                    "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/GeoIP-Anonymous-Plus.mmdb",
		"GeoIP-Anonymous-Plus-Test.mmdb":               "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/GeoIP-Anonymous-Plus-Test.mmdb",
		"GeoLite2-ASN-Test.mmdb":                       "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/GeoLite2-ASN-Test.mmdb",
		"GeoLite2-City-Test.mmdb":                      "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/GeoLite2-City-Test.mmdb",
		"GeoLite2-Country-Test.mmdb":                   "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/GeoLite2-Country-Test.mmdb",
		"MaxMind-DB-no-ipv4-search-tree.mmdb":          "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/MaxMind-DB-no-ipv4-search-tree.mmdb",
		"MaxMind-DB-string-value-entries.mmdb":         "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/MaxMind-DB-string-value-entries.mmdb",
		"MaxMind-DB-test-broken-pointers-24.mmdb":      "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/MaxMind-DB-test-broken-pointers-24.mmdb",
		"MaxMind-DB-test-broken-search-tree-24.mmdb":   "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/MaxMind-DB-test-broken-search-tree-24.mmdb",
		"MaxMind-DB-test-decoder.mmdb":                 "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/MaxMind-DB-test-decoder.mmdb",
		"MaxMind-DB-test-ipv4-24.mmdb":                 "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/MaxMind-DB-test-ipv4-24.mmdb",
		"MaxMind-DB-test-ipv4-28.mmdb":                 "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/MaxMind-DB-test-ipv4-28.mmdb",
		"MaxMind-DB-test-ipv4-32.mmdb":                 "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/MaxMind-DB-test-ipv4-32.mmdb",
		"MaxMind-DB-test-ipv6-24.mmdb":                 "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/MaxMind-DB-test-ipv6-24.mmdb",
		"MaxMind-DB-test-ipv6-28.mmdb":                 "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/MaxMind-DB-test-ipv6-28.mmdb",
		"MaxMind-DB-test-ipv6-32.mmdb":                 "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/MaxMind-DB-test-ipv6-32.mmdb",
		"MaxMind-DB-test-metadata-pointers.mmdb":       "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/MaxMind-DB-test-metadata-pointers.mmdb",
		"MaxMind-DB-test-mixed-24.mmdb":                "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/MaxMind-DB-test-mixed-24.mmdb",
		"MaxMind-DB-test-mixed-28.mmdb":                "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/MaxMind-DB-test-mixed-28.mmdb",
		"MaxMind-DB-test-mixed-32.mmdb":                "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/MaxMind-DB-test-mixed-32.mmdb",
		"MaxMind-DB-test-nested.mmdb":                  "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/MaxMind-DB-test-nested.mmdb",
		"MaxMind-DB-test-pointer-decoder.mmdb":         "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/MaxMind-DB-test-pointer-decoder.mmdb",
	}
	for fname, fpath := range dbs {
		absfn := "/tmp/" + fname
		err := downloadTestDB(fpath, absfn)
		if err != nil {
			println(err.Error())
		}
	}
}

func downloadTestDB(path, dst string) error {
	if _, err := os.Stat(dst); errors.Is(err, os.ErrNotExist) {
		dfh, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return err
		}
		defer func() { _ = dfh.Close() }()

		sfh, err := http.Get(path)
		if err != nil {
			return err
		}
		defer func() { _ = sfh.Body.Close() }()
		rdr := bufio.NewReader(sfh.Body)

		buf := make([]byte, 1024)
		for {
			n, err := rdr.Read(buf)
			if err != nil && err != io.EOF {
				return err
			}
			if n == 0 {
				break
			}
			if _, err = dfh.Write(buf); err != nil {
				return err
			}
		}
	}
	return nil
}

func TestMMDBC(t *testing.T) {}

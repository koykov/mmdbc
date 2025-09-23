package mmdbc

import (
	"bufio"
	"context"
	"errors"
	"io"
	"net/http"
	"os"
	"sync"
	"testing"
)

func init() {
	dbs := []string{
		"GeoIP2-Anonymous-IP-Test",
		"GeoIP2-City-Shield-Test",
		"GeoIP2-City-Test-Broken-Double-Format",
		"GeoIP2-City-Test-Invalid-Node-Count",
		"GeoIP2-City-Test",
		"GeoIP2-Connection-Type-Test",
		"GeoIP2-Country-Shield-Test",
		"GeoIP2-Country-Test",
		"GeoIP2-DensityIncome-Test",
		"GeoIP2-Domain-Test",
		"GeoIP2-Enterprise-Shield-Test",
		"GeoIP2-Enterprise-Test",
		"GeoIP2-IP-Risk-Test",
		"GeoIP2-ISP-Test",
		"GeoIP2-Precision-Enterprise-Shield-Test",
		"GeoIP2-Precision-Enterprise-Test",
		"GeoIP2-Static-IP-Score-Test",
		"GeoIP2-User-Count-Test",
		"GeoIP-Anonymous-Plus",
		"GeoIP-Anonymous-Plus-Test",
		"GeoLite2-ASN-Test",
		"GeoLite2-City-Test",
		"GeoLite2-Country-Test",
		"MaxMind-DB-no-ipv4-search-tree",
		"MaxMind-DB-string-value-entries",
		"MaxMind-DB-test-broken-pointers-24",
		"MaxMind-DB-test-broken-search-tree-24",
		"MaxMind-DB-test-decoder",
		"MaxMind-DB-test-ipv4-24",
		"MaxMind-DB-test-ipv4-28",
		"MaxMind-DB-test-ipv4-32",
		"MaxMind-DB-test-ipv6-24",
		"MaxMind-DB-test-ipv6-28",
		"MaxMind-DB-test-ipv6-32",
		"MaxMind-DB-test-metadata-pointers",
		"MaxMind-DB-test-mixed-24",
		"MaxMind-DB-test-mixed-28",
		"MaxMind-DB-test-mixed-32",
		"MaxMind-DB-test-nested",
		"MaxMind-DB-test-pointer-decoder",
	}
	const (
		wc  = 8
		tmp = "/tmp/mmdb-test"
	)
	if _, err := os.Stat(tmp); errors.Is(err, os.ErrNotExist) {
		if err = os.Mkdir(tmp, 0777); err != nil {
			panic(err)
		}
	}
	ch := make(chan string, wc)
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	for i := 0; i < wc; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				default:
					fname := <-ch
					absfn := tmp + "/" + fname + ".mmdb"
					fpath := "https://github.com/maxmind/MaxMind-DB/raw/refs/heads/main/test-data/" + fname + ".mmdb"
					err := downloadTestDB(fpath, absfn)
					if err != nil {
						println(err.Error())
					}
				}
			}
		}()
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < len(dbs); i++ {
			ch <- dbs[i]
		}
		cancel()
	}()
	wg.Wait()
	close(ch)
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

package mmdbcli

import (
	"context"
	"testing"
)

func TestConn(t *testing.T) {
	t.Run("get", func(t *testing.T) {
		c, err := Connect("testdata/GeoIP2-ISP-Test.mmdb")
		if err != nil {
			t.Fatal(err)
		}
		r, err := c.Gets(context.Background(), "1.128.0.0")
		if err != nil {
			t.Fatal(err)
		}
		t.Log(r) // todo parse me
	})
}

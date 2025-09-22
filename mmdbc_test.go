package mmdbc

import (
	"bufio"
	"io"
	"net/http"
	"os"
)

func downloadTestDB(path, dst string) error {
	if _, err := os.Stat(dst); err == os.ErrNotExist {
		dfh, err := os.Open(path)
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

package mmdbcli

import "io"

func decodeUint16(buf []byte, offset, size uint64) (uint16, uint64, error) {
	if offset+size > uint64(len(buf)) {
		return 0, 0, io.ErrUnexpectedEOF
	}
	b := buf[offset : offset+size]
	var r uint16
	for i := 0; i < len(b); i++ {
		r = (r << 8) | uint16(b[i])
	}
	return r, offset + size, nil
}

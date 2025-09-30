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

func decodeUint64(buf []byte, offset, size uint64) (uint64, uint64, error) {
	if offset+size > uint64(len(buf)) {
		return 0, 0, io.ErrUnexpectedEOF
	}
	b := buf[offset : offset+size]
	var r uint64
	for i := 0; i < len(b); i++ {
		r = (r << 8) | uint64(b[i])
	}
	return r, offset + size, nil
}

func decodePtr(buf []byte, offset, size uint64) (uint64, uint64, error) {
	ptrsz := ((size >> 3) & 0x3) + 1
	off1 := offset + ptrsz
	if off1 > uint64(len(buf)) {
		return 0, 0, io.ErrUnexpectedEOF
	}
	var pfx uint64
	if ptrsz != 4 {
		pfx = size & 0x7
	}
	off2 := encodeBytes(buf[offset:offset+ptrsz], pfx)
	if ptrsz > 4 {
		return 0, 0, ErrBadPointerSize
	}
	off3 := ptrsz2off[ptrsz-1]
	return off2 + off3, off1, nil
}

func encodeBytes(buf []byte, pfx uint64) (r uint64) {
	r = pfx
	for i := 0; i < len(buf); i++ {
		r = (r << 8) | uint64(buf[i])
	}
	return
}

var ptrsz2off = [4]uint64{0, 2048, 526336, 0}

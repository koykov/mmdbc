package mmdbcli

import (
	"io"

	"github.com/koykov/byteconv"
)

func decode(buf []byte, offset, size, depth uint64, dst any) (uint64, error) {
	ctrlb := buf[offset]
	offset++
	etype := entryType(ctrlb >> 5)
	if etype == entryExtended {
		if offset > uint64(len(buf)) {
			return offset, io.ErrUnexpectedEOF
		}
		etype = entryType(buf[offset] + 7)
		offset++
	}

	size = uint64(ctrlb & 0x1f)
	if size >= 29 {
		offset1 := offset + size - 28
		if offset1 >= uint64(len(buf)) {
			return offset1, io.ErrUnexpectedEOF
		}
		if size == 29 {
			size = 29 + uint64(buf[offset])
			offset = offset1
		} else {
			b := buf[offset:offset1]
			if size == 30 {
				size = b2u(b, 0) + 285
			} else {
				size = b2u(b, 0) + 65821
			}
		}
	}
	if offset+size > uint64(len(buf)) {
		return offset, io.ErrUnexpectedEOF
	}

	if etype == entryPointer {
		addr, offset, err := decodePtr(buf, offset, size)
		if err != nil {
			return addr, err
		}
		if offset >= uint64(len(buf)) {
			return addr, io.ErrUnexpectedEOF
		}
		ctrlb = buf[addr]
		if (ctrlb >> 5) == 1 {
			return addr, ErrBadPointer
		}
		if _, err = decode(buf, addr, size, depth+1, dst); err != nil {
			return offset, err
		}
		return offset, nil
	}

	src := buf[offset : offset+size]
	switch x := dst.(type) {
	case *uint16:
		*x = uint16(b2u(src, 0))
	case *uint32:
		*x = uint32(b2u(src, 0))
	case *uint64:
		*x = b2u(src, 0)
	case *string:
		*x = byteconv.B2S(src)
	default:
		return offset, ErrUnknownType
	}
	offset += size
	return offset, nil
}

func b2u(buf []byte, pfx uint64) (r uint64) {
	r = pfx
	for i := 0; i < len(buf); i++ {
		r = (r << 8) | uint64(buf[i])
	}
	return
}

func decodePtr(buf []byte, offset, size uint64) (uint64, uint64, error) {
	ptrsz := ((size >> 3) & 0x3) + 1
	nextoff := offset + ptrsz
	if nextoff > uint64(len(buf)) {
		return 0, 0, io.ErrUnexpectedEOF
	}
	var pfx uint64
	if ptrsz != 4 {
		pfx = size & 0x7
	}
	unpack := b2u(buf[offset:offset+ptrsz], pfx)
	if ptrsz > 4 {
		return 0, 0, ErrBadPointerSize
	}
	ptroff := ptrsz2off[ptrsz-1]
	return unpack + ptroff, nextoff, nil
}

var ptrsz2off = [4]uint64{0, 2048, 526336, 0}

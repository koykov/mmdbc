package mmdbcli

import (
	"fmt"
	"io"

	"github.com/koykov/byteconv"
)

type Meta struct {
	desc    map[string]string
	dbType  string
	lang    []string
	bfmaj   uint64
	bfmin   uint64
	epoch   uint64
	ipVer   uint64
	nodec   uint64
	recSize uint64
}

func (m *Meta) Description(langCode string) string {
	return m.desc[langCode]
}

func (m *Meta) EachDescription(f func(langCode, desc string)) {
	for k, v := range m.desc {
		f(k, v)
	}
}

func (m *Meta) DatabaseType() string {
	return m.dbType
}

func (m *Meta) Languages() []string {
	return m.lang
}

func (m *Meta) BinaryFormatMajorVersion() uint64 {
	return m.bfmaj
}

func (m *Meta) BinaryFormatMinorVersion() uint64 {
	return m.bfmin
}

func (m *Meta) BuildEpoch() uint64 {
	return m.epoch
}

func (m *Meta) IPVersion() uint64 {
	return m.ipVer
}

func (m *Meta) NodeCount() uint64 {
	return m.nodec
}

func (m *Meta) RecordSize() uint64 {
	return m.recSize
}

func (m *Meta) reset() {
	for k := range m.desc {
		delete(m.desc, k)
	}
	m.dbType = ""
	m.lang = m.lang[:0]
	m.bfmaj = 0
	m.bfmin = 0
	m.epoch = 0
	m.ipVer = 0
	m.nodec = 0
	m.recSize = 0
}

func (c *conn) decodeMeta() error {
	if c.meta.desc == nil {
		c.meta.desc = make(map[string]string)
	}

	var off int
	ctrlb := c.bufm[off]
	et := entryType(ctrlb >> 5)
	if et != entryMap {
		return ErrMetaRootMustBeMap
	}
	size := ctrlb & 0x1f
	if size == 0 {
		return ErrMetaEmpty
	}
	off++

	for i := 0; i < int(size); i++ {
		ctrlb = c.bufm[off]
		off++
		et1 := entryType(ctrlb >> 5)
		if et1 != entryString {
			return ErrMetaKeyMustBeString
		}
		size1 := ctrlb & 0x1f
		key := string(c.bufm[off : off+int(size1)])
		off += int(size1)
		println(key) // todo remove me
		var err error
		switch key {
		case "node_count":
			off, err = c.mustUint16(off, &c.meta.nodec)
		case "record_size":
			off, err = c.mustUint16(off, &c.meta.recSize)
		case "ip_version":
			off, err = c.mustUint16(off, &c.meta.ipVer)
		case "binary_format_major_version":
			off, err = c.mustUint16(off, &c.meta.bfmaj)
		case "binary_format_minor_version":
			off, err = c.mustUint16(off, &c.meta.bfmin)
		case "build_epoch":
			off, err = c.mustUint64(off, &c.meta.epoch)
		case "database_type":
			off, err = c.mustString(off, &c.meta.dbType)
		case "languages":
			ctrlb = c.bufm[off]
			off++
			et2 := entryType(ctrlb >> 5)
			if et2 == entryExtended {
				et2 = entryType(c.bufm[off] + 7)
				off++
			}
			if et2 != entryArray {
				return ErrMetaValueMustBeArray
			}
			size2 := ctrlb & 0x1f
			for j := 0; j < int(size2); j++ {
				var s string
				if off, err = c.mustString(off, &s); err != nil {
					break
				}
				c.meta.lang = append(c.meta.lang, s)
			}
		case "description":
			ctrlb = c.bufm[off]
			off++
			et2 := entryType(ctrlb >> 5)
			if et2 != entryMap {
				return ErrMetaValueMustBeMap
			}
			size2 := ctrlb & 0x1f
			for j := 0; j < int(size2); j++ {
				var k, v string
				if off, err = c.mustString(off, &k); err != nil {
					break
				}
				if off, err = c.mustString(off, &v); err != nil {
					break
				}
				c.meta.desc[k] = v
			}
		default:
			return fmt.Errorf("unknown meta key '%s'", key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *conn) mustUint16(off int, result *uint64) (int, error) {
	ctrlb := c.bufm[off]
	off++
	etype := entryType(ctrlb >> 5)
	if etype == entryExtended {
		if off > len(c.bufm) {
			return off, io.ErrUnexpectedEOF
		}
		etype = entryType(c.bufm[off] + 7)
		off++
	}
	if etype != entryUint16 {
		println(etype) // todo remove me
		return off, ErrMetaValueMustBeUint16
	}
	size := ctrlb & 0x1f
	v, _, err := decodeUint16(c.bufm, uint64(off), uint64(size))
	off += int(size)
	*result = uint64(v)
	return off, err
}

func (c *conn) mustUint64(off int, result *uint64) (int, error) {
	ctrlb := c.bufm[off]
	off++
	etype := entryType(ctrlb >> 5)
	if etype == entryExtended {
		if off > len(c.bufm) {
			return off, io.ErrUnexpectedEOF
		}
		etype = entryType(c.bufm[off] + 7)
		off++
	}
	if etype != entryUint64 {
		println(etype) // todo remove me
		return off, ErrMetaValueMustBeUint64
	}
	size := ctrlb & 0x1f
	v, _, err := decodeUint64(c.bufm, uint64(off), uint64(size))
	off += int(size)
	*result = v
	return off, err
}

func (c *conn) mustString(off int, result *string) (int, error) {
	ctrlb := c.bufm[off]
	off++
	etype := entryType(ctrlb >> 5)
	if etype == entryExtended {
		if off > len(c.bufm) {
			return off, io.ErrUnexpectedEOF
		}
		etype = entryType(c.bufm[off] + 7)
		off++
	}
	size := int(ctrlb & 0x1f)
	if etype == entryPointer {
		off1, _, err := decodePtr(c.bufm, uint64(off), uint64(size))
		if err != nil {
			return off, err
		}
		if off1 >= uint64(len(c.bufm)) {
			return off, io.ErrUnexpectedEOF
		}
		ctrlb = c.bufm[off1]
		if (ctrlb >> 5) == 1 {
			return off, ErrBadPointer
		}
		*result = byteconv.B2S(c.bufm[off1 : off1+uint64(size)])
		return int(off1), nil
	}
	if etype != entryString {
		println(etype)
		return off, ErrMetaValueMustBeString
	}
	*result = byteconv.B2S(c.bufm[off : off+size])
	off += size
	return off, nil
}

var (
	metaKeys = []string{
		"node_count",
		"record_size",
		"ip_version",
		"binary_format_major_version",
		"binary_format_minor_version",
		"build_epoch",
		"database_type",
		"languages",
		"description",
	}
	metaBKeys [][]byte
)

func init() {
	for i := range metaKeys {
		metaBKeys = append(metaBKeys, []byte(metaKeys[i]))
	}
}

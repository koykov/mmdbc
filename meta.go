package mmdbcli

import (
	"fmt"

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

	var off uint64
	ctrlb := c.bufm[off]
	etype := entryType(ctrlb >> 5)
	if etype != entryMap {
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
		etype1 := entryType(ctrlb >> 5)
		if etype1 != entryString {
			return ErrMetaKeyMustBeString
		}
		size1 := uint64(ctrlb & 0x1f)
		key := byteconv.B2S(c.bufm[off : off+size1])
		off += size1
		var err error
		switch key {
		case "node_count":
			off, err = decode(c.bufm, off, 0, 0, &c.meta.nodec)
		case "record_size":
			off, err = decode(c.bufm, off, 0, 0, &c.meta.recSize)
		case "ip_version":
			off, err = decode(c.bufm, off, 0, 0, &c.meta.ipVer)
		case "binary_format_major_version":
			off, err = decode(c.bufm, off, 0, 0, &c.meta.bfmaj)
		case "binary_format_minor_version":
			off, err = decode(c.bufm, off, 0, 0, &c.meta.bfmin)
		case "build_epoch":
			off, err = decode(c.bufm, off, 0, 0, &c.meta.epoch)
		case "database_type":
			off, err = decode(c.bufm, off, 0, 0, &c.meta.dbType)
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
				if off, err = decode(c.bufm, off, 0, 0, &s); err != nil {
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
				if off, err = decode(c.bufm, off, 0, 0, &k); err != nil {
					break
				}
				if off, err = decode(c.bufm, off, 0, 0, &v); err != nil {
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

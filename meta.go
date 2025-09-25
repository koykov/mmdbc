package mmdbc

import (
	"bytes"

	"github.com/koykov/version"
)

type Meta struct {
	desc    map[string]string
	dbType  string
	lang    []string
	bfver   version.Version64
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
	return uint64(m.bfver.Major())
}

func (m *Meta) BinaryFormatMinorVersion() uint64 {
	return uint64(m.bfver.Minor())
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
	m.bfver.Reset()
	m.epoch = 0
	m.ipVer = 0
	m.nodec = 0
	m.recSize = 0
}

func (c *conn) decodeMeta() error {
	for i := 0; i < len(metaKeys); i++ {
		key := metaKeys[i]
		idx := bytes.Index(c.bufm, metaBKeys[i])
		if idx == -1 {
			continue
		}
		switch key {
		case "node_count":
		case "record_size":
		case "ip_version":
		case "binary_format_major_version":
		case "binary_format_minor_version":
		case "build_epoch":
		case "database_type":
		case "languages":
		case "description":
		}
	}
	return nil
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

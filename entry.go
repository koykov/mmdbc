package mmdbcli

type entryType uint8

const (
	entryExtended entryType = iota
	entryPointer
	entryString
	entryDouble
	entryBytes
	entryUint16
	entryUint32
	entryMap
	entryInt32
	entryUint64
	entryUint128
	entryArray
	entryContainer
	entryEndMarker
	entryBool
	entryFloat
)

type entry struct {
	hasData      bool
	data         [16]byte
	offset       uint32
	offsetToNext uint32
	dataSize     uint32
	typ          uint32
}

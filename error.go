package mmdbcli

import "errors"

var (
	ErrMetaNotFound          = errors.New("meta not found")
	ErrMetaRootMustBeMap     = errors.New("meta root must be a map")
	ErrMetaKeyMustBeString   = errors.New("meta key must be a string")
	ErrMetaValueMustBeUint16 = errors.New("meta value must be uint16 number")
	ErrMetaValueMustBeUint32 = errors.New("meta value must be uint32 number")
	ErrMetaValueMustBeUint64 = errors.New("meta value must be uint64 number")
	ErrMetaValueMustBeString = errors.New("meta value must be string")
	ErrMetaValueMustBeMap    = errors.New("meta value must be map")
	ErrMetaValueMustBeArray  = errors.New("meta value must be array")
	ErrMetaEmpty             = errors.New("meta is empty")
	ErrBadPointerSize        = errors.New("bad pointer size")
	ErrBadPointer            = errors.New("bad pointer")
	ErrUnknownType           = errors.New("unknown type")
)

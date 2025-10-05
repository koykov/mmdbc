package mmdbcli

import "errors"

var (
	ErrMetaNotFound         = errors.New("meta not found")
	ErrMetaRootMustBeMap    = errors.New("meta root must be a map")
	ErrMetaKeyMustBeString  = errors.New("meta key must be a string")
	ErrMetaValueMustBeMap   = errors.New("meta value must be map")
	ErrMetaValueMustBeArray = errors.New("meta value must be array")
	ErrMetaEmpty            = errors.New("meta is empty")

	ErrBadPointerSize = errors.New("bad pointer size")
	ErrBadPointer     = errors.New("bad pointer")
	ErrUnknownType    = errors.New("unknown type")

	ErrInvalidPrefix  = errors.New("invalid prefix")
	ErrOverflowPrefix = errors.New("overflow prefix: use ipv6 over ipv4 database")
	ErrOverflowIPv6   = errors.New("bits overflow ipv6 (128 bits)")
)

package mmdbcli

import "errors"

var (
	ErrMetaNotFound          = errors.New("meta not found")
	ErrMetaRootMustBeMap     = errors.New("meta root must be a map")
	ErrMetaKeyMustBeString   = errors.New("meta key must be a string")
	ErrMetaValueMustBeUint16 = errors.New("meta value must be uint16 number")
	ErrMetaValueMustBeUint64 = errors.New("meta value must be uint64 number")
	ErrMetaValueMustBeString = errors.New("meta value must be string")
	ErrMetaEmpty             = errors.New("meta is empty")
)

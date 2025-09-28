package mmdbcli

import "errors"

var (
	ErrMetaNotFound         = errors.New("meta not found")
	ErrMetaRootMustBeMap    = errors.New("meta root must be a map")
	ErrMetaKeyMustBeString  = errors.New("meta key must be a string")
	ErrMetaValueMustBeUin16 = errors.New("meta value must be uint16 number")
	ErrMetaEmpty            = errors.New("meta is empty")
)

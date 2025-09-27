package mmdbcli

import "errors"

var (
	ErrMetaNotFound        = errors.New("meta not found")
	ErrMetaRootMustBeMap   = errors.New("meta root must be a map")
	ErrMetaKeyMustBeString = errors.New("meta key must be a string")
	ErrMetaEmpty           = errors.New("meta is empty")
)

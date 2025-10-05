package mmdbcli

type NetworkOption uint64

const (
	NetworkOptionIncludeAliased NetworkOption = 1 << iota
	NetworkOptionIncludeEmptyNetwork
	NetworkOptionSkipEmptyTuple
)

const (
	NetworkOptionNoOptions NetworkOption = 0
	NetworkOptionAll                     = NetworkOptionIncludeAliased | NetworkOptionIncludeEmptyNetwork | NetworkOptionSkipEmptyTuple
)

func (c *conn) EachNetwork(fn func(*Tuple) error) error {
	// todo implement me
	return nil
}

func (c *conn) EachNetworkWithOptions(fn func(*Tuple) error, options NetworkOption) error {
	// todo implement me
	return nil
}

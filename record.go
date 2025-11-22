package mmdbcli

import "github.com/koykov/indirect"

type Record struct {
	cnptr    uintptr
	off, pfx uint64
}

func (r *Record) Get(path string) *Value {
	cn := r.indirectConn()
	if cn == nil {
		return nil
	}
	return cn.lookup(r.off, path)
}

func (r *Record) indirectConn() *conn {
	if r.cnptr == 0 {
		return nil
	}
	return (*conn)(indirect.ToUnsafePtr(r.cnptr))
}

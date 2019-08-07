package sql2go

import "github.com/go-xorm/core"

type TmpData struct {
	Tables  []*core.Table
	Imports map[string]string
	Models  string
}

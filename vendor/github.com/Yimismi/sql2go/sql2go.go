package sql2go

import (
	"github.com/go-xorm/core"
)

func FromSql(sql string, args *convertArgs) ([]byte, error) {
	parse := ParseSql
	return generateCode(sql, parse, args)
}

func FromFile(fileName string, args *convertArgs) ([]byte, error) {
	parse := ParseSqlFile
	return generateCode(fileName, parse, args)
}

func generateCode(src string, parse func(string) ([]*core.Table, error), args *convertArgs) ([]byte, error) {
	tables, err := parse(src)
	if err != nil {
		return nil, err
	}
	goTmpl := NewGolangTmp(args)
	return goTmpl.GenerateGo(tables)
}

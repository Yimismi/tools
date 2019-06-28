package sql2go

import (
	"bytes"
	"fmt"
	"github.com/go-xorm/core"
	"github.com/knocknote/vitess-sqlparser/tidbparser/ast"
	tidbparser "github.com/knocknote/vitess-sqlparser/tidbparser/parser"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func ParseSqlFile(fileName string) ([]*core.Table, error) {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	sql := string(bs)
	return ParseSql(sql)
}

func ParseSql(sql string) ([]*core.Table, error) {

	tables := make([]*core.Table, 0)

	stmts, err := tidbparser.New().Parse(sql, "", "")
	if err != nil {
		return nil, err
	}
	for _, stmt := range stmts {
		if cstmt, ok := stmt.(*ast.CreateTableStmt); ok {
			tb, e := cvtDDL2Table(cstmt)
			if e != nil {
				fmt.Fprint(os.Stderr, e)
				continue
			}
			tables = append(tables, tb)
		}
	}

	return tables, nil
}

func cvtDDL2Table(cs *ast.CreateTableStmt) (*core.Table, error) {
	table := core.NewEmptyTable()
	table.Name = cs.Table.Name.String()
	table.StoreEngine = "InnoDB"
	for _, op := range cs.Options {
		switch op.Tp {
		// comment will be in `""` after parsing
		case ast.TableOptionComment:
			table.Comment = fomatStr(op.StrValue)
		case ast.TableOptionEngine:
			table.StoreEngine = op.StrValue
		}
	}
	// parse columns
	cols := make(map[string]*core.Column)
	colSeq := make([]string, 0)
	for _, c := range cs.Cols {
		col := new(core.Column)
		col.Indexes = make(map[string]int)
		col.Name = c.Name.Name.String()
		col.Nullable = true

		// parse columns type
		colType := c.Tp.String()
		cts := strings.Split(colType, "(")
		colName := cts[0]
		colType = strings.ToUpper(colName)
		var len1, len2 int
		if len(cts) == 2 {
			idx := strings.Index(cts[1], ")")
			if colType == core.Enum && cts[1][0] == '\'' { //enum
				options := strings.Split(cts[1][0:idx], ",")
				col.EnumOptions = make(map[string]int)
				for k, v := range options {
					v = strings.TrimSpace(v)
					v = strings.Trim(v, "'")
					col.EnumOptions[v] = k
				}
			} else if colType == core.Set && cts[1][0] == '\'' {
				options := strings.Split(cts[1][0:idx], ",")
				col.SetOptions = make(map[string]int)
				for k, v := range options {
					v = strings.TrimSpace(v)
					v = strings.Trim(v, "'")
					col.SetOptions[v] = k
				}
			} else {
				var err error
				lens := strings.Split(cts[1][0:idx], ",")
				len1, err = strconv.Atoi(strings.TrimSpace(lens[0]))
				if err != nil {
					return nil, err
				}
				if len(lens) == 2 {
					len2, err = strconv.Atoi(lens[1])
					if err != nil {
						return nil, err
					}
				}
			}
		}
		if colType == "FLOAT UNSIGNED" {
			colType = "FLOAT"
		}
		if colType == "DOUBLE UNSIGNED" {
			colType = "DOUBLE"
		}
		col.Length = len1
		col.Length2 = len2
		if _, ok := core.SqlTypes[colType]; ok {
			col.SQLType = core.SQLType{Name: colType, DefaultLength: len1, DefaultLength2: len2}
		} else {
			return nil, fmt.Errorf("Unknown colType %v", colType)
		}
		// parse columns type end

		for _, op := range c.Options {
			expr := ""
			if op.Expr != nil {
				var buf bytes.Buffer
				op.Expr.Format(&buf)
				expr = buf.String()
			}
			switch op.Tp {
			case ast.ColumnOptionNotNull:
				col.Nullable = false
			case ast.ColumnOptionDefaultValue:
				col.Default = fomatStr(expr)
				if col.Default == "" {
					col.DefaultIsEmpty = true
				}
			case ast.ColumnOptionComment:
				// comment will be in `""` after parsing
				col.Comment = fomatStr(expr)
			case ast.ColumnOptionAutoIncrement:
				col.IsAutoIncrement = true
			case ast.ColumnOptionPrimaryKey:
				col.IsPrimaryKey = true

			}
		}
		if col.SQLType.IsText() || col.SQLType.IsTime() {
			if col.Default != "" {
				col.Default = "'" + col.Default + "'"
			} else {
				if col.DefaultIsEmpty {
					col.Default = "''"
				}
			}
		}
		cols[col.Name] = col
		colSeq = append(colSeq, col.Name)
	}
	// parse columns end

	for _, cst := range cs.Constraints {
		switch cst.Tp {
		case ast.ConstraintPrimaryKey:
			for _, key := range cst.Keys {
				cols[key.Column.Name.String()].IsPrimaryKey = true
			}
		}
	}

	for _, name := range colSeq {
		table.AddColumn(cols[name])
	}

	return table, nil
}
func fomatStr(s string) string {
	if len(s) >= 2 && strings.Index(s, "\"") == 0 && strings.LastIndex(s, "\"") == len(s)-1 {
		return s[1 : len(s)-1]
	} else {
		return s
	}
}

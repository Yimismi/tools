### sql2go

---
This repository is adapted from ``reverse command`` of ``go-xorm/cmd``(https://github.com/go-xorm/cmd/blob/master/xorm/reverse.go).
However ``go-xorm/cmd`` generates go code from database-source while this repository converts create-statement of mysql to go struct.

### Installation

*[NOTE] Required Go version more than 1.9*
```
go get -u github.com/Yimismi/sql2go
```
---

### Examples
```$xslt
	sql := `
CREATE TABLE IF NOT EXISTS t_person (
  f_age INT(11) NULL,
  f_id INT(11) PRIMARY KEY AUTO_INCREMENT NOT NULL,
  f_name VARCHAR(30) NOT NULL,
  f_sex VARCHAR(2) NULL
  ) ENGINE=InnoDB;
`
	args := sql2go.NewConvertArgs().SetGenJson(true).
		SetPackageName("test").
		SetColPrefix("f_").
		SetTablePrefix("t_")

	code, err := sql2go.FromSql(sql, args)
	if err != nil {
		t.Error(err)
		return
	}
	f, err := os.Create("db_struct3.go")
	if err != nil {
		t.Error(err)
		return
	}
	defer f.Close()
	f.Write(code)
```
*The library can parse multiple Statements and will skip the statement which don't create tables*
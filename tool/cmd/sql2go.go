package cmd

import (
	"github.com/Yimismi/sql2go"
	"github.com/Yimismi/tools/tool"
	"github.com/Yimismi/tools/web"
	"github.com/gin-gonic/gin"
	"github.com/ngaut/log"
)

type Sql2goToolArgs struct {
	Sql         string `json:"src"`
	ColPrefix   string `json:"col_prefix"`
	TablePrefix string `json:"table_prefix"`
	GenJson     bool   `json:"gen_json"`
	GenXorm     bool   `json:"gen_xorm"`
	PackageName string `json:"package_name"`
	OtherTags   string `json:"other_tags"`
}

var sql2goToolArgsDesc = []*tool.ArgDesc{
	{Name: "col_prefix", Type: "string", DefaultValue: "", Desc: "列名前缀"},
	{Name: "table_prefix", Type: "string", DefaultValue: "", Desc: "表名前缀"},
	{Name: "gen_json", Type: "bool", DefaultValue: "true", Desc: "是否产生json tag", Optional: []bool{true, false}},
	{Name: "gen_xorm", Type: "bool", DefaultValue: "true", Desc: "是否产生xorm tag", Optional: []bool{true, false}},
	{Name: "other_tags", Type: "string", DefaultValue: "", Desc: "tags"},
}

func init() {
	web.RegisterTool(&Sql2goTool{
		Tool: tool.Tool{
			Name: "sql2go",
			Url:  "/tool/sql2go.html",
			Desc: "将mysql的create语句转化成go struct结构",
		}})
}

type Sql2goTool struct {
	tool.Tool
}

func NewSql2goToolArgs() *Sql2goToolArgs {
	return &Sql2goToolArgs{"", "", "", true, true, "db", ""}
}

func (t *Sql2goTool) GetArgsDesc() []*tool.ArgDesc {
	return sql2goToolArgsDesc
}
func (t *Sql2goTool) Usage() string {
	return ""
}

func (t *Sql2goTool) Exec(args *Sql2goToolArgs) ([]byte, error) {

	a := sql2go.NewConvertArgs().
		SetColPrefix(args.ColPrefix).
		SetTablePrefix(args.TablePrefix).
		SetGenJson(args.GenJson).
		SetGenXorm(args.GenXorm).
		SetOtherTags(args.OtherTags)
	return sql2go.FromSql(args.Sql, a)
}
func (t *Sql2goTool) Run(ctx *gin.Context) {
	args := NewSql2goToolArgs()
	err := ctx.BindJSON(args)
	log.Infof("--ip:%v....req:%v\n", ctx.ClientIP(), args)
	if err != nil {
		ctx.JSON(200, map[string]string{"error": err.Error()})
		return
	}
	bs, err := t.Exec(args)
	if err != nil {
		ctx.JSON(200, map[string]string{"error": err.Error()})
		return
	}
	ctx.JSON(200, map[string]string{"output": string(bs)})
}

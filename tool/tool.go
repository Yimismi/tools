package tool

import "github.com/gin-gonic/gin"

type ArgDesc struct {
	Name         string
	Type         string
	DefaultValue string
	Desc         string
	Optional     interface{}
}
type ToolInterface interface {
	Usage() string
	GetArgsDesc() []*ArgDesc
	GetName() string
	GetUrl() string
}
type WebTool interface {
	ToolInterface
	Run(ctx *gin.Context)
}
type Tool struct {
	Name string
	Url  string
	Desc string
}

func (t Tool) GetName() string {
	return t.Name
}
func (t Tool) GetUrl() string {
	return t.Url
}

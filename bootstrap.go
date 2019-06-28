//go:generate statik -src=./web -f

package main

import (
	_ "github.com/Yimismi/tools/tool/cmd"
	"github.com/Yimismi/tools/web"
)

const CONFPATH = "./config"

func main() {
	web.Run(CONFPATH)
}

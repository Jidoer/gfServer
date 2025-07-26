package main

import (
	_ "gfAdmin/internal/packed"

	_ "gfAdmin/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"gfAdmin/internal/cmd"
    _ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"


)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}

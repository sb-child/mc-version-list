package main

import (
	_ "mcvl/internal/packed"
	"os"

	_ "mcvl/internal/logic"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	"mcvl/internal/cmd"
)

func main() {
	g.Log().SetStdoutPrint(false)
	g.Log().SetWriter(os.Stderr)
	g.Log().SetWriterColorEnable(true)
	cmd.Main.Run(gctx.GetInitCtx())
}

package cmd

import (
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "Main command",
		Func:  MainCmd,
	}
)

func init() {
	Main.AddCommand(&gcmd.Command{
		Name:  "serve",
		Usage: "serve",
		Brief: "todo",
		Func:  ServeCmd,
	}, &gcmd.Command{
		Name:      "generate",
		Usage:     "generate",
		Arguments: []gcmd.Argument{{Name: "file", Short: "f", Orphan: true, Brief: "output file name or output to stdout"}},
		Brief:     "Generate version list and output",
		Func:      GenerateCmd,
	})
}

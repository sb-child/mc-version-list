package cmd

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gcmd"
)

func MainCmd(ctx context.Context, parser *gcmd.Parser) (err error) {
	return gerror.NewCode(gcode.CodeInvalidParameter, "please use generate command instead")
}

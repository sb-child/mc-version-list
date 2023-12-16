package cmd

import (
	"context"
	"fmt"
	"mcvl/internal/service"
	"mcvl/utility"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
)

func GenerateCmd(ctx context.Context, parser *gcmd.Parser) (err error) {
	fn := parser.GetOpt("file", "").String()
	vm, err := service.VerFetcher().Fetch(ctx)
	if err != nil {
		return gerror.Wrap(err, "Failed to fetch version data")
	}
	pm, err := service.ProtoFetcher().Fetch(ctx)
	if err != nil {
		return gerror.Wrap(err, "Failed to fetch protocol data")
	}
	sum := utility.Summarize(vm, pm)
	s, err := gjson.EncodeString(sum)
	if err != nil {
		return gerror.Wrap(err, "Failed to encode JSON")
	}
	if len(fn) <= 0 {
		g.Log().Info(ctx, "output to stdout...")
		fmt.Println(s)
	} else {
		g.Log().Info(ctx, "save to file")
		f, err := gfile.Create(fn)
		if err != nil {
			return gerror.Wrap(err, "Failed to create file")
		}
		g.Log().Infof(ctx, "saving to %s", f.Name())
		_, err = f.WriteString(s)
		if err != nil {
			return gerror.Wrap(err, "Failed to write file")
		}
		err = f.Close()
		if err != nil {
			return gerror.Wrap(err, "Failed to close file")
		}
	}
	g.Log().Infof(ctx, "Done.")
	return nil
}

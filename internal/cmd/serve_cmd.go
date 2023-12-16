package cmd

import (
	"context"
	"mcvl/internal/controller/hello"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

func ServeCmd(ctx context.Context, parser *gcmd.Parser) (err error) {
	// todo
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		group.Bind(
			hello.NewV1(),
		)
	})
	s.Run()
	return nil
}

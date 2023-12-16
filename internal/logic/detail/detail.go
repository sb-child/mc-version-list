package detail

import (
	"context"
	"mcvl/internal/service"
	"mcvl/utility"
	"net/url"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gctx"
)

type sDetail struct {
	cli *gclient.Client
	url string
}

func init() {
	service.RegisterDetail(New())
}

func New() *sDetail {
	x := sDetail{}
	if err := x.Init(); err != nil {
		g.Log().Fatalf(gctx.GetInitCtx(), "Failed to init Detail service: %s", err.Error())
	}
	return &x
}

func (x *sDetail) Init() error {
	x.cli = utility.DefaultClient()
	// x.cli.SetHeader("Referer", "https://zh.minecraft.wiki/w/%E5%9F%BA%E5%B2%A9%E7%89%88%E4%B8%93%E7%94%A8%E6%9C%8D%E5%8A%A1%E5%99%A8?variant=zh-cn")
	// x.cli.SetHeader("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	// x.cli.SetHeader("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	var err error
	x.url, err = url.PathUnescape("https://zh.minecraft.wiki/w/Module:Protocol_version/Versions")
	if err != nil {
		return gerror.Wrap(err, "URL unescape error")
	}
	return nil
}

func (x *sDetail) Fetch(ctx context.Context) (utility.Protocols, error) {
	g.Log().Infof(ctx, "Fetching %s", x.url)
	req, err := x.cli.Get(ctx, x.url)
	if err != nil {
		return nil, gerror.Wrap(err, "HTTP error")
	}
	defer req.Close()
	g.Log().Infof(ctx, "Reading content...")
	s := req.ReadAllString()
	g.Log().Infof(ctx, "Parsing...")
	pl := utility.FindProtocolFromMediawiki(s)
	if len(pl) <= 0 {
		g.Log().Warning(ctx, "Cannot parse the content")
		return nil, gerror.NewCode(gcode.CodeNotFound, "Cannot parse the content")
	}
	// for _, v := range pl {
	// 	g.Log().Infof(ctx, "protocol: %+v", v)
	// }
	return pl, nil
}

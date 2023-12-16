package utility

import (
	"net/http"

	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gtime"
)

func DefaultClient() *gclient.Client {
	x := gclient.New().
		SetAgent("curl/8.5.0").
		SetTimeout(gtime.S*5).
		SetBrowserMode(true).
		SetRetry(5, gtime.S).
		SetRedirectLimit(10).
		Use(func(c *gclient.Client, r *http.Request) (*gclient.Response, error) {
			r.Header.Del("Traceparent")
			resp, err := c.Next(r)
			return resp, err
		})
	return x
}

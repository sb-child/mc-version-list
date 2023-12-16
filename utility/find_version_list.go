package utility

import (
	"regexp"
	"strings"
	"sync"

	"github.com/gogf/gf/v2/util/gconv"
)

var versionListRegex = regexp.MustCompile(`(?s)({{History\|\|(.*?)\|(.*?)}})+`)
var versionPageRegex = regexp.MustCompile(`(?s)(\[\[(基岩版.*?)\]\])`)
var versionLinkRegex = regexp.MustCompile(`(?s)(link=基岩版(.*?)\|)`)

type Version struct {
	Version       string
	SimpleVersion string
	Page          string
	Released      bool
	Major         uint8
	Minor         uint8
	Build         uint8
	Revision      uint8
}

type Versions map[string]Version

func parseVersion(v []string, ch chan<- Version, wg *sync.WaitGroup) {
	defer wg.Done()
	item := Version{
		Version:  v[2],
		Page:     v[3],
		Released: true,
	}
	x := versionPageRegex.FindAllStringSubmatch(item.Page, 1)
	if vl := versionLinkRegex.FindAllStringSubmatch(item.Page, 1); len(vl) == 1 && len(vl[0]) == 3 {
		item.Page = "基岩版" + vl[0][2]
		item.SimpleVersion = vl[0][2]
	} else if len(x) <= 0 || len(x[0]) < 2 {
		item.Page = "基岩版" + item.Version
		item.SimpleVersion = item.Version
	} else {
		item.Page = x[0][2]
		item.SimpleVersion = strings.TrimPrefix(item.Page, "基岩版")
	}
	vs := strings.Split(item.Version, ".")
	if len(vs) >= 3 {
		item.Major = gconv.Uint8(vs[0])
		item.Minor = gconv.Uint8(vs[1])
		item.Build = gconv.Uint8(vs[2])
		if len(vs) == 4 {
			item.Revision = gconv.Uint8(vs[3])
		}
	}
	if strings.Contains(v[3], "测试版服务器软件") {
		item.Released = false
	}
	ch <- item
}

func FindVersionList(s string) Versions {
	x := versionListRegex.FindAllStringSubmatch(s, -1)
	if x == nil {
		return nil
	}
	r := make(Versions, len(x))
	ch := make(chan Version)
	wg := sync.WaitGroup{}
	for _, v := range x {
		if len(v) != 4 {
			continue
		}
		wg.Add(1)
		go parseVersion(v, ch, &wg)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for item := range ch {
		r[item.Page] = item
	}
	return r
}

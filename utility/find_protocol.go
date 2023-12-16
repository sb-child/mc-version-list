package utility

import (
	"html"
	"regexp"
	"sync"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/st3fan/html2text"
)

var protocolRegex = regexp.MustCompile(`(?s)ver\(pocket, '(.*?)', (.*?)\)`)

type Protocols map[string]Protocol

type Protocol struct {
	Version int64
	Page    string
}

func parseProtocol(v []string, ch chan<- Protocol, wg *sync.WaitGroup) {
	defer wg.Done()
	item := Protocol{
		Page: v[1],
	}
	if v[2] == "unknown" {
		item.Version = -1
	} else {
		item.Version = gconv.Int64(v[2])
	}
	ch <- item
}

func FindProtocolFromMediawiki(s string) Protocols {
	st := html2text.HTML2Text(s)
	st = html.UnescapeString(st)
	ss := protocolRegex.FindAllStringSubmatch(st, -1)
	if ss == nil {
		return nil
	}
	r := make(Protocols, len(ss))
	ch := make(chan Protocol)
	wg := sync.WaitGroup{}
	for _, v := range ss {
		if len(v) != 3 {
			continue
		}
		wg.Add(1)
		go parseProtocol(v, ch, &wg)
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

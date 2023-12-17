package utility

import (
	"sort"
	"strings"

	"github.com/gogf/gf/v2/os/gtime"
)

type SummarizedVersions map[string]SummarizedVersion

type SummarizedVersion struct {
	Version       string `json:"version"`
	SimpleVersion string `json:"simple_version"`
	Page          string `json:"page"`
	Released      bool   `json:"released"`
	Protocol      int64  `json:"protocol"`
	Major         uint8  `json:"major_v"`
	Minor         uint8  `json:"minor_v"`
	Build         uint8  `json:"build_v"`
	Revision      uint8  `json:"rev_v"`
}

type Summarized struct {
	VersionList         []string           `json:"versions"`
	Items               SummarizedVersions `json:"v"`
	ReverseVersionList  map[string]string  `json:"rv"`
	ReverseProtocolList map[int64][]string `json:"rp"`
	UpdatedAt           uint64             `json:"update_timestamp_ms"`
}

func Summarize(vm Versions, pm Protocols) Summarized {
	vl := make([]Version, 0, len(vm))
	for _, v := range vm {
		vl = append(vl, v)
	}
	sort.Slice(vl, func(i, j int) bool {
		return vl[i].Major < vl[j].Major ||
			(vl[i].Major == vl[j].Major &&
				vl[i].Minor < vl[j].Minor) ||
			(vl[i].Major == vl[j].Major &&
				vl[i].Minor == vl[j].Minor &&
				vl[i].Build < vl[j].Build) ||
			(vl[i].Major == vl[j].Major &&
				vl[i].Minor == vl[j].Minor &&
				vl[i].Build == vl[j].Build &&
				vl[i].Revision < vl[j].Revision)
	})
	vls := make([]string, 0, len(vm))
	svs := make(SummarizedVersions, len(vm))
	rv := make(map[string]string, len(vm))
	rp := make(map[int64][]string, len(vm))
	for _, v := range vl {
		vv := strings.ReplaceAll(v.Version, ".", "_")
		svv := strings.ReplaceAll(v.SimpleVersion, ".", "_")
		vls = append(vls, svv)
		sv := SummarizedVersion{
			Version:       v.Version,
			SimpleVersion: v.SimpleVersion,
			Page:          v.Page,
			Released:      v.Released,
			Protocol:      -2,
			Major:         v.Major,
			Minor:         v.Minor,
			Build:         v.Build,
			Revision:      v.Revision,
		}
		if p, ok := pm[v.Page]; ok {
			sv.Protocol = p.Version
		}
		svs[svv] = sv
		rv[vv] = svv
		if p, ok := rp[sv.Protocol]; ok {
			p = append(p, svv)
			rp[sv.Protocol] = p
		} else {
			rp[sv.Protocol] = []string{svv}
		}
	}
	return Summarized{
		VersionList:         vls,
		Items:               svs,
		UpdatedAt:           uint64(gtime.TimestampMilli()),
		ReverseVersionList:  rv,
		ReverseProtocolList: rp,
	}
}

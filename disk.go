package gosi

import (
	"encoding/json"
	"math"

	"github.com/shirou/gopsutil/v3/disk"

	"github.com/inhies/go-bytesize"
)

type DiskStat struct {
	Name        string `json:"name"`
	Total       string `json:"total"`
	Free        string `json:"free"`
	Used        string `json:"used"`
	UsedPercent uint   `json:"usedPercent"`
}

type DiskStats []DiskStat

func (d DiskStats) Json() []byte {
	j, _ := json.Marshal(d)
	return j
}

func Disk() DiskStats {
	ret := []DiskStat{}

	p, _ := disk.Partitions(true)

	bytesize.Format = "%.0f "
	for _, v := range p {
		if isVaildPartition(v) {
			d, _ := disk.Usage(v.Mountpoint)
			total := bytesize.New(float64(d.Total))
			free := bytesize.New(float64(d.Free))
			used := bytesize.New(float64(d.Used))
			usedPercent := math.Round(d.UsedPercent*10) / 10

			ret = append(ret, DiskStat{d.Path, total.String(), free.String(), used.String(), uint(usedPercent)})
		}
	}

	return ret
}

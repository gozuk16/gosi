package goss

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
	var ret []byte
	for _, v := range d {
		j, _ := json.Marshal(v)
		ret = append(ret, j...)
	}
	return ret
}

func Disk() DiskStats {
	ret := []DiskStat{}

	p, _ := disk.Partitions(true)

	bytesize.Format = "%.0f "
	for _, v := range p {
		//fmt.Println(v.Mountpoint, v.Device, v.Fstype, v.Opts)
		if isVaildPartition(v.Mountpoint, v.Opts) {
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

package mylib

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"

	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"

	"github.com/inhies/go-bytesize"
)

/*
type DiskInfo struct {
	name        string  `json:"name"`
	total       uint64  `json:"total"`
	free        uint64  `json:"free"`
	used        uint64  `json:"used"`
	usedPercent float64 `json:"usedPercent"`
}
*/

func Host() []byte {
	v, _ := host.Info()

	// convert to JSON. String() is also implemented
	fmt.Println(v)

	return []byte(v.String())
}

func Mem() []byte {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)

	return []byte(v.String())
}

func Disk() []byte {
	p, _ := disk.Partitions(true)

	// convert to JSON. String() is also implemented
	//fmt.Println(p)

	var disks []map[string]interface{}
	for _, v := range p {
		//b, _ := json.Marshal(v)
		d, _ := disk.Usage(v.Mountpoint)
		total := bytesize.New(float64(d.Total))
		free := bytesize.New(float64(d.Free))
		used := bytesize.New(float64(d.Used))
		bytesize.Format = "%.1f "
		usedPercent := math.Round(d.UsedPercent*10) / 10
		//fmt.Printf("%v %v %v", total, free, used)
		di := map[string]interface{}{
			"name":        d.Path,
			"total":       total.String(),
			"free":        free.String(),
			"used":        used.String(),
			"usedPercent": strconv.FormatFloat(usedPercent, 'f', -1, 64) + "%",
		}
		disks = append(disks, di)
	}

	j, _ := json.Marshal(disks)
	//fmt.Println(string(j))
	return j
}

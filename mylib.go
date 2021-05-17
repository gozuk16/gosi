package mylib

import (
	"encoding/json"
	"fmt"

	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
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
		di := map[string]interface{}{
			"name":        d.Path,
			"total":       d.Total,
			"free":        d.Free,
			"used":        d.Used,
			"usedPercent": d.UsedPercent,
		}
		disks = append(disks, di)
	}

	j, _ := json.Marshal(disks)
	//fmt.Println(string(j))
	return j
}

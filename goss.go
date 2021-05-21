package goss

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"

	"github.com/inhies/go-bytesize"
)

func Host() []byte {
	v, _ := host.Info()

	// convert to JSON. String() is also implemented
	fmt.Println(v)

	return []byte(v.String())
}

func Cpu() []byte {
	c, _ := cpu.Percent(time.Millisecond*200, false)
	core, _ := cpu.Percent(time.Millisecond*200, true)
	//fmt.Printf("%f%%\n", c)

	var cpu map[string]interface{}
	total := strconv.FormatFloat(c[0], 'f', 0, 64) + "%"
	var p = []string{}
	for _, v := range core {
		p = append(p, strconv.FormatFloat(v, 'f', 0, 64)+"%")
	}
	cpu = map[string]interface{}{
		"total":  total,
		"percpu": p,
	}
	j, _ := json.Marshal(cpu)

	return j
}

func Load() []byte {
	l, _ := load.Avg()
	//fmt.Printf("%f\n", l)

	var loadAvg map[string]interface{}
	loadAvg = map[string]interface{}{
		"load1":  strconv.FormatFloat(l.Load1, 'f', 2, 64),
		"load5":  strconv.FormatFloat(l.Load5, 'f', 2, 64),
		"load15": strconv.FormatFloat(l.Load15, 'f', 2, 64),
	}
	j, _ := json.Marshal(loadAvg)

	return j
}

func Mem() []byte {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	//fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	//fmt.Println(v)

	var mem map[string]interface{}
	total := bytesize.New(float64(v.Total))
	available := bytesize.New(float64(v.Available))
	used := bytesize.New(float64(v.Used))
	bytesize.Format = "%.1f "
	usedPercent := math.Round(v.UsedPercent*10) / 10
	//fmt.Printf("%v %v %v", total, free, used)
	mem = map[string]interface{}{
		"total":       total.String(),
		"available":   available.String(),
		"used":        used.String(),
		"usedPercent": strconv.FormatFloat(usedPercent, 'f', -1, 64) + "%",
	}
	j, _ := json.Marshal(mem)
	return j
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

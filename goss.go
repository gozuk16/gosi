package goss

import (
	"encoding/json"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"

	"github.com/inhies/go-bytesize"
)

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
		"usedPercent": uint(usedPercent),
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
		//bytesize.Format = "%.1f "
		bytesize.Format = "%.0f "
		usedPercent := math.Round(d.UsedPercent*10) / 10
		//fmt.Printf("%v %v %v", total, free, used)
		di := map[string]interface{}{
			"name":        d.Path,
			"total":       total.String(),
			"free":        free.String(),
			"used":        used.String(),
			"usedPercent": uint(usedPercent),
		}
		disks = append(disks, di)
	}

	j, _ := json.Marshal(disks)
	//fmt.Println(string(j))
	return j
}

// Processes プロセス情報をまとめて返す
func Processes(pids []int32) []map[string]interface{} {
	//result := [][]byte{}
	var result []map[string]interface{}
	for _, pid := range pids {
		//p, _ := process.NewProcess(pid)
		pi := Process(pid)
		result = append(result, pi)
	}

	//j, _ := json.Marshal(result)
	return result
}

// Process プロセス情報を返す
func Process(pid int32) map[string]interface{} {
	p, _ := process.NewProcess(pid)

	name, _ := p.Name()
	cpupercent, _ := p.CPUPercent()
	cpupercent = cpupercent * 100
	cputime, _ := p.Times()
	memory, _ := p.MemoryInfo()
	cmdline, _ := p.Cmdline()
	exe, _ := p.Exe()
	cwd, _ := p.Cwd()
	createtime, _ := p.CreateTime()
	isexists, _ := process.PidExists(pid)
	statuses, _ := p.Status()
	status := strings.Join(statuses, ", ")
	ppid, _ := p.Ppid()
	children, _ := p.Children()
	var procChildren []map[string]interface{}
	for _, c := range children {
		cn, _ := c.Name()
		ccmd, _ := c.Cmdline()
		cmem, _ := c.MemoryInfo() // TODO:rssを取り出して変換
		cproc := map[string]interface{}{
			"name":    cn,
			"cmdline": ccmd,
			"pid":     c.Pid,
			"mem":     cmem,
		}
		procChildren = append(procChildren, cproc)
	}
	//fmt.Printf("%v %v %v %v", name, memory, isexists, status)

	var proc map[string]interface{}
	proc = map[string]interface{}{
		"name":       name,
		"cpuPercent": math.Round(cpupercent*10) / 10,
		"cpuTotal":   math.Round(cputime.Total()*100) / 100,
		"cpuUser":    cputime.User,
		"cpuSystem":  cputime.System,
		"cpuIdle":    cputime.Idle,
		"cpuIowait":  cputime.Iowait,
		"vms":        bytesize.New(float64(memory.VMS)).String(),
		"rss":        bytesize.New(float64(memory.RSS)).String(),
		"swap":       bytesize.New(float64(memory.Swap)).String(),
		"cmdline":    cmdline,
		"exe":        exe,
		"cwd":        cwd,
		"createTime": time.Unix(createtime/1000, 0).Format(timeformat),
		"isExists":   isexists,
		"status":     status,
		"pid":        p.Pid,
		"ppid":       ppid,
		"children":   procChildren,
	}
	/*
		j, _ := json.Marshal(proc)
		return j
	*/
	return proc
}

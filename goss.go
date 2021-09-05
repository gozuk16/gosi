package gosi

import (
	"math"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/process"

	"github.com/inhies/go-bytesize"
)

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

package goss

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/load"
)

type CpuStat struct {
	Total  uint   `json:"total"`
	Num    int    `json:"num"`
	Load1  string `json:"load1"`
	Load5  string `json:"load5"`
	Load15 string `json:"load15"`
}

var cpupercent uint

func (c CpuStat) Json() []byte {
	j, _ := json.Marshal(c)
	return j
}

// Cpu CPU情報を取得
func Cpu() *CpuStat {
	ret := &CpuStat{}
	ret.Total = cpupercent
	ret.Num, _ = cpu.Counts(true)
	ret.Load1, ret.Load5, ret.Load15 = loadAvg()

	return ret
}

// RefreshCpu グローバル変数のCPU情報を更新
func RefreshCpu() {
	c, _ := cpu.Percent(time.Millisecond*1000, false)
	cpupercent = uint(c[0])
}

// load avarageをフォーマットして返す
func loadAvg() (load1 string, load5 string, load15 string) {
	l, _ := load.Avg()
	return strconv.FormatFloat(l.Load1, 'f', 2, 64), strconv.FormatFloat(l.Load5, 'f', 2, 64), strconv.FormatFloat(l.Load15, 'f', 2, 64)
}

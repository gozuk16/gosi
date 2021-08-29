package goss

import (
	"encoding/json"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

type CpuStat struct {
	Total uint `json:"total"`
	Num   int  `json:"num"`
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

	return ret
}

// RefreshCpu グローバル変数のCPU情報を更新
func RefreshCpu() {
	c, _ := cpu.Percent(time.Millisecond*1000, false)
	cpupercent = uint(c[0])
}

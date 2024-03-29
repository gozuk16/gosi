package gosi

import (
	"encoding/json"
	"math"
	"strconv"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/load"
)

type CpuStat struct {
	Total  float64 `json:"total"`
	Num    int     `json:"num"`
	Load1  string  `json:"load1"`
	Load5  string  `json:"load5"`
	Load15 string  `json:"load15"`
}

var cpupercent float64

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
	// 0はこの瞬間のCPU使用率
	c, _ := cpu.Percent(0, false)
	// 小数点第一位で四捨五入
	cpupercent = math.Round(c[0]*10) / 10
}

// load avarageをフォーマットして返す
func loadAvg() (load1 string, load5 string, load15 string) {
	l, _ := load.Avg()
	return strconv.FormatFloat(l.Load1, 'f', 2, 64), strconv.FormatFloat(l.Load5, 'f', 2, 64), strconv.FormatFloat(l.Load15, 'f', 2, 64)
}

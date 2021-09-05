package gosi

import (
	"encoding/json"
	"math"

	"github.com/shirou/gopsutil/v3/mem"

	"github.com/inhies/go-bytesize"
)

type MemStat struct {
	Total       string `json:"total"`
	Available   string `json:"available"`
	Used        string `json:"used"`
	UsedPercent uint   `json:"usedPercent"`
}

func (m MemStat) Json() []byte {
	j, _ := json.Marshal(m)
	return j
}

func Mem() *MemStat {
	ret := &MemStat{}

	v, _ := mem.VirtualMemory()

	bytesize.Format = "%.1f "
	ret.Total = bytesize.New(float64(v.Total)).String()
	ret.Available = bytesize.New(float64(v.Available)).String()
	ret.Used = bytesize.New(float64(v.Used)).String()
	ret.UsedPercent = uint(math.Round(v.UsedPercent*10) / 10)

	return ret
}

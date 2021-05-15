package mylib

import (
	"encoding/json"
	"fmt"

	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

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
	v, _ := disk.Partitions(true)

	// convert to JSON. String() is also implemented
	fmt.Println(v)

	d, _ := json.Marshal(v)
	//for i := 0; i < len(v); i++ {
	//	d[i] = []byte(v[i].String())
	//}
	return d
}

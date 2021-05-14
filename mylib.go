package mylib

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/mem"
)

func Mem() []byte {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)

	return []byte(v.String())
}

package gosi_test

import (
	//	"fmt"
	"fmt"
	"testing"
	"time"

	"github.com/gozuk16/gosi"
)

/*
func TestCpuTotal(t *testing.T) {
	gosi.RefreshCpu()
	i := gosi.Cpu()
	if i.Total <= 0 {
		t.Errorf("CPU Total = %d, Failed", i.Total)
	} else {
		fmt.Println("CPU Total:", i.Total)
	}
}
*/

func TestCpuStatJson(t *testing.T) {
	cpu := gosi.CpuStat{
		Total:  50,
		Num:    4,
		Load1:  "1.20",
		Load5:  "0.75",
		Load15: "0.60",
	}

	expected := `{"total":50,"num":4,"load1":"1.20","load5":"0.75","load15":"0.60"}`

	actual := cpu.Json()

	if string(actual) != expected {
		t.Errorf("Expected JSON: %s, but got: %s", expected, string(actual))
	}
}

func TestCpu(t *testing.T) {
	gosi.RefreshCpu()
	cpu := gosi.Cpu()

	if cpu.Total == 0 {
		t.Error("CPU total is zero")
	}
	if cpu.Num == 0 {
		t.Error("CPU num is zero")
	}
	if cpu.Load1 == "" {
		t.Error("CPU load1 is empty")
	}
	if cpu.Load5 == "" {
		t.Error("CPU load5 is empty")
	}
	if cpu.Load15 == "" {
		t.Error("CPU load15 is empty")
	}
}

func TestRefreshCpu(t *testing.T) {
	time.Sleep(1 * time.Second)
	gosi.RefreshCpu()
	time.Sleep(1 * time.Second)
	gosi.RefreshCpu()
	time.Sleep(1 * time.Second)
	gosi.RefreshCpu()

	cpu := gosi.Cpu()

	fmt.Printf("cpu.Total: %v\n", cpu.Total)
	if cpu.Total == 0 {
		t.Error("CPU total is zero")
	}
}

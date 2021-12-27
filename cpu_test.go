package gosi_test

import (
	"fmt"
	"testing"

	"github.com/gozuk16/gosi"
)

func TestCpuTotal(t *testing.T) {
	gosi.RefreshCpu()
	i := gosi.Cpu()
	if i.Total <= 0 {
		t.Errorf("CPU Total = %d, Failed", i.Total)
	} else {
		fmt.Println("CPU Total:", i.Total)
	}
}

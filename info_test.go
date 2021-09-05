package gosi_test

import (
	"testing"

	"github.com/gozuk16/gosi"
)

func TestCpuTemperatures(t *testing.T) {
	i := gosi.Info()
	if i.CpuTemperature == "" || i.CpuTemperature == "0℃" {
		t.Errorf("CPU Temperature = %s, Failed", i.CpuTemperature)
	}
}

package goss_test

import (
	"testing"

	"github.com/gozuk16/goss"
)

func TestCpuTemperatures(t *testing.T) {
	i := goss.Info()
	if i.CpuTemperature == "" || i.CpuTemperature == "0℃" {
		t.Errorf("CPU Temperature = %s, Failed", i.CpuTemperature)
	}
}

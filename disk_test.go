package gosi_test

import (
	"testing"

	"github.com/gozuk16/gosi"
)

func TestDiskStatsJson(t *testing.T) {
	diskStats := gosi.DiskStats{
		{
			Name:        "C:",
			Total:       "500 GB",
			Free:        "200 GB",
			Used:        "300 GB",
			UsedPercent: 60,
		},
		{
			Name:        "D:",
			Total:       "1 TB",
			Free:        "500 GB",
			Used:        "500 GB",
			UsedPercent: 50,
		},
	}

	expected := `[{"name":"C:","total":"500 GB","free":"200 GB","used":"300 GB","usedPercent":60},{"name":"D:","total":"1 TB","free":"500 GB","used":"500 GB","usedPercent":50}]`

	actual := diskStats.Json()

	if string(actual) != expected {
		t.Errorf("Expected JSON: %s, but got: %s", expected, string(actual))
	}
}

func TestDisk(t *testing.T) {
	diskStats := gosi.Disk()

	if len(diskStats) == 0 {
		t.Error("Disk stats is empty")
	}

	for _, diskStat := range diskStats {
		if diskStat.Name == "" {
			t.Error("Disk name is empty")
		}
		if diskStat.Total == "" {
			t.Error("Disk total is empty")
		}
		if diskStat.Free == "" {
			t.Error("Disk free is empty")
		}
		if diskStat.Used == "" {
			t.Error("Disk used is empty")
		}
		if diskStat.UsedPercent == 0 {
			t.Error("Disk used percent is zero")
		}
	}
}

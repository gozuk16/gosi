package gosi_test

import (
	//"encoding/json"
	"testing"

	"github.com/gozuk16/gosi"
)

func TestMemStatJson(t *testing.T) {
	mem := gosi.MemStat{
		Total:       "10 GB",
		Available:   "5 GB",
		Used:        "5 GB",
		UsedPercent: 50,
	}

	expected := `{"total":"10 GB","available":"5 GB","used":"5 GB","usedPercent":50}`

	actual := mem.Json()

	if string(actual) != expected {
		t.Errorf("Expected JSON: %s, but got: %s", expected, string(actual))
	}
}

func TestMem(t *testing.T) {
	mem := gosi.Mem()

	if mem.Total == "" {
		t.Error("Total memory is empty")
	}
	if mem.Available == "" {
		t.Error("Available memory is empty")
	}
	if mem.Used == "" {
		t.Error("Used memory is empty")
	}
	if mem.UsedPercent == 0 {
		t.Error("Used percent is zero")
	}
}


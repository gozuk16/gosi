package gosi_test

import (
	"encoding/json"
	"testing"

	"github.com/gozuk16/gosi"
)

func TestInfo(t *testing.T) {
	info := gosi.Info()

	if info == nil {
		t.Error("Expected non-nil InfoStat, got nil")
	}

	if info.Hostname == "" {
		t.Error("Expected non-empty Hostname, got empty")
	}

	// Json メソッドのテスト
	jsonData := info.Json()
	var parsedInfo gosi.InfoStat
	err := json.Unmarshal(jsonData, &parsedInfo)
	if err != nil {
		t.Errorf("Error decoding JSON: %v", err)
	}

	// Json メソッドで得られた JSON データが元の情報と一致するか確認
	if parsedInfo.Hostname != info.Hostname {
		t.Errorf("Expected Hostname %s, got %s", info.Hostname, parsedInfo.Hostname)
	}
}

func TestInfoIpAddresses(t *testing.T) {
	info := gosi.Info()

	// IpAddres フィールドが nil でないことを確認
	if info.IpAddres == nil {
		t.Error("Expected non-nil IpAddres, got nil")
	}

	// IpAddres フィールドが空でないことを確認
	if len(info.IpAddres) == 0 {
		t.Error("Expected non-empty IpAddres, got empty")
	}

	// IpAddres の各要素に対してアサーションを行う
	for _, ipAddr := range info.IpAddres {
		if ipAddr.Name == "" {
			t.Error("Expected non-empty Name in IpAddr, got empty")
		}

		if ipAddr.IpAddr == "" {
			t.Error("Expected non-empty IpAddr in IpAddr, got empty")
		}
	}
}

// TestUptime2String privateだけどuptime2stringのテストをしておく
func TestUptime2String(t *testing.T) {

	// uptimeが1日以上の場合
	uptimeOverOneDay := uint64(172800) // 2 days
	expectedOverOneDay := "2 days, 0:00"
	resultOverOneDay := gosi.ExportedUptime2String(uptimeOverOneDay)
	if resultOverOneDay != expectedOverOneDay {
		t.Errorf("Expected %s for uptime over one day, got %s", expectedOverOneDay, resultOverOneDay)
	}

	// uptimeが1日未満の場合
	uptimeLessThanOneDay := uint64(7200) // 2 hours
	expectedLessThanOneDay := "2:00"
	resultLessThanOneDay := gosi.ExportedUptime2String(uptimeLessThanOneDay)
	if resultLessThanOneDay != expectedLessThanOneDay {
		t.Errorf("Expected %s for uptime less than one day, got %s", expectedLessThanOneDay, resultLessThanOneDay)
	}
}

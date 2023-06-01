package gosi

import (
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/shirou/gopsutil/v3/host"
	psnet "github.com/shirou/gopsutil/v3/net"
)

//const timeformat = "2006/01/02 15:04:05"
const timeformat = "2006/01/02 15:04:05.000000000"

type IpAddr struct {
	Name   string `json:"name"`
	IpAddr string `json:"ipaddr"`
}

type InfoStat struct {
	Hostname        string   `json:"hostname"`
	OS              string   `json:"os"`
	Platform        string   `json:"platform"`
	PlatformFamily  string   `json:"platformFamily"`
	PlatformVersion string   `json:"platformVersion"`
	KernelArch      string   `json:"kernelArch"`
	Uptime          string   `json:"uptime"`
	BootTime        string   `json:"bootTime"`
	ServerTime      string   `json:"serverTime"`
	CpuTemperature  string   `json:"cpuTemperature"`
	IpAddres        []IpAddr `json:"ipaddr"`
}

func (s InfoStat) Json() []byte {
	j, _ := json.Marshal(s)
	return j
}

// Info ホスト情報を取得
func Info() *InfoStat {
	ret := &InfoStat{}

	i, _ := host.Info()

	ret.Hostname = i.Hostname
	ret.OS = i.OS
	ret.Platform = i.Platform
	ret.PlatformFamily = i.PlatformFamily
	ret.PlatformVersion = i.PlatformVersion
	ret.KernelArch = i.KernelArch
	ret.Uptime = uptime2string(i.Uptime)
	//ret.BootTime = time.Unix(int64(i.BootTime), 0).Format(timeformat)
	ret.BootTime = time.Unix(int64(i.BootTime), int64(i.BootTime)%1000000000).Format(timeformat)
	ret.ServerTime = time.Now().Format(timeformat)

	n, _ := psnet.Interfaces()
	ipaddres := make([]IpAddr, 0)
	for _, v := range n {
		if len(v.Addrs) > 0 {
			for _, a := range v.Addrs {
				ipaddr, ipnet, err := net.ParseCIDR(a.Addr)
				if err != nil {
					fmt.Println(err)
				}
				if ipnet.IP.To4() != nil && !ipnet.IP.IsLoopback() && !ipnet.IP.IsLinkLocalUnicast() {
					ipaddres = append(ipaddres, IpAddr{v.Name, ipaddr.String()})
				}
			}
		}
	}

	ret.IpAddres = ipaddres

	// 仮想環境では温度が取得できないが、WindowsではVirtualizationSystemが取れないのでとりあえずコメントにしておく
	//ret.CpuTemperature, _ = getTemperatures()

	return ret
}

// uptime2string uptime(経過秒)をuptimeと同じ"0 days, 00:00"形式に変換する
func uptime2string(uptime uint64) string {
	const oneDay int = 60 * 60 * 24

	if int(uptime) > oneDay {
		day := int(uptime) / oneDay
		secondsOfTheDay := day * oneDay
		d := time.Duration(int(uptime)-secondsOfTheDay) * time.Second
		d = d.Round(time.Minute)
		h := d / time.Hour
		d -= h * time.Hour
		m := d / time.Minute
		return fmt.Sprintf("%d days, %d:%02d", day, h, m)
	} else {
		d := time.Duration(int(uptime)) * time.Second
		d = d.Round(time.Minute)
		h := d / time.Hour
		d -= h * time.Hour
		m := d / time.Minute
		return fmt.Sprintf("%d:%02d", h, m)
	}
}

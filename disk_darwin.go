package gosi

import (
	//"log"
	"strings"

	"github.com/shirou/gopsutil/v3/disk"
)

// isVaildPartition システムで予約されているパーティション、ネットワークマウントやGoogleDrive、dmg(Optsがro+nodevのもの)をfalseで返す
func isVaildPartition(p disk.PartitionStat) bool {
	if p.Mountpoint == "/dev" ||
		p.Mountpoint == "/System/Volumes/VM" ||
		p.Mountpoint == "/System/Volumes/Preboot" ||
		p.Mountpoint == "/private/var/vm" ||
		p.Mountpoint == "/Volumes/Recovery" ||
		p.Mountpoint == "/System/Volumes/Data/home" ||
		strings.Contains(p.Mountpoint, "/Volumes/.timemachine") ||
		strings.Contains(p.Mountpoint, "/Volumes/com.apple.TimeMachine") ||
		strings.Contains(p.Mountpoint, "/System/Volumes/Update") {
		return false
	} else if p.Fstype == "smbfs" || p.Fstype == "dfsfuse_DFS" {
		return false
	} else if len(p.Opts) > 1 {
		dmg := 0
		for _, o := range p.Opts {
			if o == "nodev" || o == "ro" {
				dmg++
			}
		}
		if dmg == 2 {
			return false
		}
	}
	//log.Println(p.Mountpoint, p.Device, p.Fstype, p.Opts)
	return true
}

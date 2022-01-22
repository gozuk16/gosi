package gosi

import (
	"github.com/shirou/gopsutil/v3/disk"
)

// isVaildPartition システムで予約されているパーティション、ネットワークマウントや特殊なものをfalseで返す
// ネットワークドライブはgopsutilでは最初から除外されている
// GoogleDriveの判定方法が分からないのでいまのところ全部trueで返す
func isVaildPartition(p disk.PartitionStat) bool {
	return true
}

package goss

// isVaildPartition システムで予約されているパーティション、ネットワークマウントやdmg(Optsがnodevのもの)をfalseで返す
func isVaildPartition(name string, opts []string) bool {
	if name == "/dev" ||
		name == "/System/Volumes/VM" ||
		name == "/System/Volumes/Preboot" ||
		name == "/System/Volumes/Update" ||
		name == "/System/Volumes/Data/home" {
		return false
	} else if len(opts) > 0 {
		for _, v := range opts {
			if v == "nodev" {
				return false
			}
		}
	}
	return true
}

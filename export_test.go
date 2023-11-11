package gosi

// ExportedUptime2String は uptime2string をテストから呼び出すためのエクスポートされた関数です。
func ExportedUptime2String(uptime uint64) string {
	return uptime2string(uptime)
}

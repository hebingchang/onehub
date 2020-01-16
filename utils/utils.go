package utils

import "fmt"

func QuotaFormat(quota_int int64) string {
	quota := float64(quota_int)
	if quota / 1024.0 < 1 {
		return fmt.Sprintf("%.2f Bytes", quota)
	}
	if quota / 1024.0 / 1024.0 < 1 {
		return fmt.Sprintf("%.2f KB", quota / 1024.0)
	}
	if quota / 1024.0 / 1024.0 / 1024.0 < 1 {
		return fmt.Sprintf("%.2f MB", quota / 1024.0 / 1024.0)
	}
	if quota / 1024.0 / 1024.0 / 1024.0 / 1024.0 < 1 {
		return fmt.Sprintf("%.2f GB", quota / 1024.0 / 1024.0 / 1024.0)
	}
	return fmt.Sprintf("%.2f TB", quota / 1024.0 / 1024.0 / 1024.0 / 1024.0)
}

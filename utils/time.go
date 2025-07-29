package utils

import (
	"time"
)

// WIBTimeNow mengembalikan waktu sekarang dalam zona Asia/Jakarta (WIB)
func WIBTimeNow() time.Time {
	loc, err := time.LoadLocation("Asia/Jakarta")
    if err != nil {
		Warn("Falling back to manual +07 timezone due to missing Asia/Jakarta")
        loc = time.FixedZone("WIB", 7*3600)
    }
	return time.Now().In(loc)
}

// WIBTimestamp mengembalikan unix timestamp (detik) zona WIB
func WIBTimestamp() int64 {
	return WIBTimeNow().Unix()
}

// WIBTimestampString mengembalikan string waktu WIB dengan format yyyy-mm-dd HH:MM:SS
func WIBTimestampString() string {
	return WIBTimeNow().Format("2006-01-02 15:04:05")
}

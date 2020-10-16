package timeutil

import (
	"time"
)

const (
	Datetime14Layout      = "20060102150405"
	Datetime8Layout       = "20060102"
	Datetime6Layout       = "200601"
	YYYYMMDDHHMMSS_LAYOUT = "2006-01-02 15:04:05"
	YYYYMMDDHHMM_LAYOUT   = "2006-01-02 15:04"
	YYYYMMDD_LAYOUT       = "2006-01-02"
)

var (
	// ShangHaiLOC, _ = time.LoadLocation("Asia/Shanghai")
	EmptyTime, _   = time.Parse("2006-01-02 15:04:05 Z0700 MST", "1979-11-30 00:00:00 +0000 GMT")
)

// yyyy-MM-dd hh:mm:ss 年-月-日 时:分:秒
func FmtDatetimeString(t time.Time) string {
	return t.Format(YYYYMMDDHHMMSS_LAYOUT)
}

// yyyy-MM-dd hh:mm 年-月-日 时:分
func FmtDatetimeMString(t time.Time) string {
	return t.Format(YYYYMMDDHHMM_LAYOUT)
}

// yy-MM-dd 年-月-日
func FmtDateString(t time.Time) string {
	return t.Format(YYYYMMDD_LAYOUT)
}

// yyyyMMddhhmmss 年月日时分秒
func FmtDatetime14String(t time.Time) string {
	return t.Format(Datetime14Layout)
}

// yyyyMMdd 年月日
func FmtDatetime8String(t time.Time) string {
	return t.Format(Datetime8Layout)
}

// yyyyMM  年月
func FmtDatetime6String(t time.Time) string {
	return t.Format(Datetime6Layout)
}


func ComputeEndTime(times int, unit string) time.Time {
	ctime := time.Now()
	switch unit {
	case "second":
		return ctime.Add(time.Second * time.Duration(times))
	case "minute":
		return ctime.Add(time.Minute * time.Duration(times))
	case "hour":
		return ctime.Add(time.Hour * time.Duration(times))
	case "day":
		return ctime.Add(time.Hour * 24 * time.Duration(times))
	case "week":
		return ctime.Add(time.Hour * 24 * 7 * time.Duration(times))
	case "month":
		return ctime.Add(time.Hour * 24 * 30 * time.Duration(times))
	case "year":
		return ctime.Add(time.Hour * 24 * 365 * time.Duration(times))
	default:
		return ctime
	}
}

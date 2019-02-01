package utils

import "time"

//GetUTCTime returns utc ms for current time
func GetUTCTimeNow() uint64 {
	return uint64(time.Now().Unix() * 1000)
}

func MakeTimestamp(t time.Time) uint64 {
	if t.UnixNano() < 0 {
		return uint64(0)
	}
	return uint64(t.UnixNano() / int64(time.Millisecond))
}

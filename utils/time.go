package utils

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
)

var DefaultTimeLayout = "2006-01-02 15:04"
var YMDHIS = "2006-01-02 15:04:05"

func FormatTime(src string) (time.Time, error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return time.ParseInLocation("2006-01-02T15:04:05", src, loc)
}

// MakeTimestamp 生成毫秒时间戳
func MakeTimestamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// MakeTimestampCurrent 生成当前时间的毫秒时间戳
func MakeCurrnetTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func TimeToPBTimestamp(t time.Time) *google_protobuf.Timestamp {
	ts, _ := ptypes.TimestampProto(t)
	return ts
}

func PBTimestampToString(ts *google_protobuf.Timestamp) string {
	return ptypes.TimestampString(ts)
}

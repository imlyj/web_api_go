package utils

import (
	"fmt"
	"reflect"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"
const TimeShortFormat = "2006-01-02"

func Int2str(value int64)string{
	return time.Unix(value, 0).Format(TimeFormat)
}

func Str2int(value string) int64{
	var t time.Time
	if len(value) == len(TimeShortFormat){
		t ,_ = time.Parse(TimeShortFormat,value)
	}else{
		t ,_ = time.Parse(TimeFormat,value)
	}
	return t.Unix()
}

// convert any numeric value to int64
func ToInt64(value interface{}) (d int64, err error) {
	val := reflect.ValueOf(value)
	switch value.(type) {
		case int, int8, int16, int32, int64:
			d = val.Int()
		case uint, uint8, uint16, uint32, uint64:
			d = int64(val.Uint())
		default:
			err = fmt.Errorf("ToInt64 need numeric not `%T`", value)
	}
	return
}

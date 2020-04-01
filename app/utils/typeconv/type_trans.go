package typeconv

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

//nolint:gosimple
func ToInt64(data interface{}) (res int64, err error) {
	val := reflect.ValueOf(data)
	switch data.(type) {
	case int, int8, int16, int32, int64:
		res = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		res = int64(val.Uint())
	case float64:
		res = int64(data.(float64))
	case float32:
		res = int64(data.(float32))
	case string:
		res, err = strconv.ParseInt(strings.TrimSpace(data.(string)), 10, 64)
	case []byte:
		res, err = strconv.ParseInt(strings.TrimSpace(string(data.([]byte))), 10, 64)
	default:
		res, err = strconv.ParseInt(fmt.Sprintf("%v", data), 10, 64)
	}
	return
}

//nolint:gosimple
func ToInt32(data interface{}) (res int32, err error) {
	val := reflect.ValueOf(data)
	switch data.(type) {
	case int, int8, int16, int32, int64:
		res = int32(val.Int())
	case uint, uint8, uint16, uint32, uint64:
		res = int32(val.Uint())
	case float64:
		res = int32(data.(float64))
	case float32:
		res = int32(data.(float32))
	case string:
		res64, _ := strconv.ParseInt(strings.TrimSpace(data.(string)), 10, 64)
		res = int32(res64)
	case []byte:
		var res64 int64
		res64, err = strconv.ParseInt(strings.TrimSpace(string(data.([]byte))), 10, 64)
		res = int32(res64)
	default:
		var res64 int64
		res64, err = strconv.ParseInt(fmt.Sprintf("%v", data), 10, 64)
		res = int32(res64)
	}
	return
}

//nolint:gosimple
func ToUInt64(data interface{}) (res uint64, err error) {
	val := reflect.ValueOf(data)
	switch data.(type) {
	case int, int8, int16, int32, int64:
		res = uint64(val.Int())
	case uint, uint8, uint16, uint32, uint64:
		res = uint64(val.Uint())
	case float64:
		res = uint64(data.(float64))
	case float32:
		res = uint64(data.(float32))
	case string:
		res, err = strconv.ParseUint(strings.TrimSpace(data.(string)), 10, 64)
	case []byte:
		res, err = strconv.ParseUint(strings.TrimSpace(string(data.([]byte))), 10, 64)
	default:
		res, err = strconv.ParseUint(fmt.Sprintf("%v", data), 10, 64)
	}
	return
}

//nolint:gosimple
func ToInt(data interface{}) (res int, err error) {
	val := reflect.ValueOf(data)
	switch data.(type) {
	case int, int8, int16, int32, int64:
		res = int(val.Int())
	case uint, uint8, uint16, uint32, uint64:
		res = int(val.Uint())
	case float64:
		res = int(data.(float64))
	case float32:
		res = int(data.(float32))
	case string:
		res, err = strconv.Atoi(strings.TrimSpace(data.(string)))
	case []byte:
		res, err = strconv.Atoi(strings.TrimSpace(string(data.([]byte))))
	default:
		res, err = strconv.Atoi(fmt.Sprintf("%v", data))
	}
	return
}

//nolint:gosimple
func ToDateTime(data interface{}) (res time.Time, err error) {
	switch data.(type) {
	case []byte:
		res, err = time.ParseInLocation("2006-01-02 15:04:05", strings.TrimSpace(string(data.([]byte))), time.Local)
	case string:
		res, err = time.ParseInLocation("2006-01-02 15:04:05", strings.TrimSpace(data.(string)), time.Local)
	default:
		res, err = time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("%v", data), time.Local)
	}
	return
}

//nolint:gosimple
func ToDate(data interface{}) (res time.Time, err error) {
	switch data.(type) {
	case []byte:
		res, err = time.ParseInLocation("2006-01-02", strings.TrimSpace(string(data.([]byte))), time.Local)
	case string:
		res, err = time.ParseInLocation("2006-01-02", strings.TrimSpace(data.(string)), time.Local)
	default:
		res, err = time.ParseInLocation("2006-01-02", fmt.Sprintf("%v", data), time.Local)
	}
	return
}

//nolint:gosimple
func ToFloat32(data interface{}) (res float32, err error) {
	val := reflect.ValueOf(data)
	switch data.(type) {
	case int, int8, int16, int32, int64:
		res = float32(val.Int())
	case uint, uint8, uint16, uint32, uint64:
		res = float32(val.Uint())
	case float64:
		res = float32(data.(float64))
	case float32:
		res = data.(float32)
	case string:
		var res64 float64
		res64, err = strconv.ParseFloat(strings.TrimSpace(data.(string)), 32)
		res = float32(res64)
	default:
		var res64 float64
		res64, err = strconv.ParseFloat(fmt.Sprintf("%v", data), 32)
		res = float32(res64)
	}
	return
}

//nolint:gosimple
func ToFloat64(data interface{}) (res float64, err error) {
	val := reflect.ValueOf(data)
	switch data.(type) {
	case int, int8, int16, int32, int64:
		res = float64(val.Int())
	case uint, uint8, uint16, uint32, uint64:
		res = float64(val.Uint())
	case float64:
		res = data.(float64)
	case float32:
		res = float64(data.(float32))
	case string:
		res, err = strconv.ParseFloat(strings.TrimSpace(data.(string)), 64)
	default:
		res, err = strconv.ParseFloat(fmt.Sprintf("%v", data), 64)
	}
	return
}

//nolint:gosimple
func ToString(data interface{}) (res string) {
	switch v := data.(type) {
	case float64:
		res = strconv.FormatFloat(data.(float64), 'f', 6, 64)
	case float32:
		res = strconv.FormatFloat(float64(data.(float32)), 'f', 6, 32)
	case int:
		res = strconv.FormatInt(int64(data.(int)), 10)
	case int64:
		res = strconv.FormatInt(data.(int64), 10)
	case uint:
		res = strconv.FormatUint(uint64(data.(uint)), 10)
	case uint64:
		res = strconv.FormatUint(data.(uint64), 10)
	case uint32:
		res = strconv.FormatUint(uint64(data.(uint32)), 10)
	case string:
		res = data.(string)
	case []byte:
		res = string(v)
	default:
		res = ""
	}
	return
}

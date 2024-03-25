package strutil

import (
	"fmt"
	"github.com/lucky-xin/xyz-common-go/value"
	"strconv"
)

func IsEmpty(s string) bool {
	return s == ""
}

func IsNotEmpty(s string) bool {
	return s != ""
}

func ToString(val interface{}) string {
	if value.IsNil(val) {
		return ""
	}
	return fmt.Sprint(val)
}

func ToInt8(s string, defaultVal int8) int8 {
	if IsEmpty(s) {
		return defaultVal
	}
	if atoi, err := strconv.Atoi(s); err == nil {
		return int8(atoi)
	}
	return defaultVal
}

func ToInt16(s string, defaultVal int16) int16 {
	if IsEmpty(s) {
		return defaultVal
	}
	if atoi, err := strconv.Atoi(s); err == nil {
		return int16(atoi)
	}
	return defaultVal
}

func ToInt32(s string, defaultVal int32) int32 {
	if IsEmpty(s) {
		return defaultVal
	}
	if atoi, err := strconv.Atoi(s); err == nil {
		return int32(atoi)
	}
	return defaultVal
}

func ToInt64(s string, defaultVal int64) int64 {
	if IsEmpty(s) {
		return defaultVal
	}
	if atoi, err := strconv.Atoi(s); err == nil {
		return int64(atoi)
	}
	return defaultVal
}

func ToInt(s string, defaultVal int) int {
	if IsEmpty(s) {
		return defaultVal
	}
	if atoi, err := strconv.Atoi(s); err == nil {
		return atoi
	}
	return defaultVal
}

func ToFloat32(s string, defaultVal float32) float32 {
	if IsEmpty(s) {
		return defaultVal
	}
	if atoi, err := strconv.ParseFloat(s, 32); err == nil {
		return float32(atoi)
	}
	return defaultVal
}

func ToFloat64(s string, defaultVal float64) float64 {
	if IsEmpty(s) {
		return defaultVal
	}
	if atoi, err := strconv.ParseFloat(s, 64); err == nil {
		return atoi
	}
	return defaultVal
}

func ToBool(s string, defaultVal bool) bool {
	if IsEmpty(s) {
		return defaultVal
	}
	if atoi, err := strconv.ParseBool(s); err == nil {
		return atoi
	}
	return defaultVal
}

package collutil

import (
	"github.com/lucky-xin/xyz-common-go/strutil"
)

func Map(values ...interface{}) (m map[string]interface{}) {
	length := len(values)
	m = make(map[string]interface{}, length/2)
	for i := 0; i+1 < length; i = i + 2 {
		m[values[i].(string)] = values[i+1]
	}
	return
}

func StrVal(m map[string]interface{}, k, defaultVal string) string {
	if m == nil {
		return defaultVal
	}
	v := m[k]
	return strutil.ToString(v)
}

func Int8Val(m map[string]interface{}, k string, defaultVal int8) int8 {
	if m == nil {
		return defaultVal
	}
	v := m[k]
	return strutil.ToInt8(strutil.ToString(v), defaultVal)
}

func Int16Val(m map[string]interface{}, k string, defaultVal int16) int16 {
	if m == nil {
		return defaultVal
	}
	v := m[k]
	return strutil.ToInt16(strutil.ToString(v), defaultVal)
}

func Int32Val(m map[string]interface{}, k string, defaultVal int32) int32 {
	if m == nil {
		return defaultVal
	}
	v := m[k]
	return strutil.ToInt32(strutil.ToString(v), defaultVal)
}

func Int64Val(m map[string]interface{}, k string, defaultVal int64) int64 {
	if m == nil {
		return defaultVal
	}
	v := m[k]
	return strutil.ToInt64(strutil.ToString(v), defaultVal)
}

func ToInt(m map[string]interface{}, k string, defaultVal int) int {
	if m == nil {
		return defaultVal
	}
	v := m[k]
	return strutil.ToInt(strutil.ToString(v), defaultVal)
}

func ToFloat32(m map[string]interface{}, k string, defaultVal float32) float32 {
	if m == nil {
		return defaultVal
	}
	v := m[k]
	return strutil.ToFloat32(strutil.ToString(v), defaultVal)
}

func ToFloat64(m map[string]interface{}, k string, defaultVal float64) float64 {
	if m == nil {
		return defaultVal
	}
	v := m[k]
	return strutil.ToFloat64(strutil.ToString(v), defaultVal)
}

func ToBool(m map[string]interface{}, k string, defaultVal bool) bool {
	if m == nil {
		return defaultVal
	}
	v := m[k]
	return strutil.ToBool(strutil.ToString(v), defaultVal)
}

package env

import (
	"github.com/lucky-xin/xyz-common-go/strutil"
	"os"
	"strconv"
	"strings"
)

func GetString(key, defaultVal string) string {
	if val, ex := os.LookupEnv(key); ex {
		return val
	}
	return defaultVal
}

func GetInt32(key string, defaultVal int32) int32 {
	if val, ex := os.LookupEnv(key); ex {
		return strutil.ToInt32(val, defaultVal)
	}
	return defaultVal
}

func GetInt64(key string, defaultVal int64) int64 {
	if val, ex := os.LookupEnv(key); ex {
		return strutil.ToInt64(val, defaultVal)
	}
	return defaultVal
}

func GetInt(key string, defaultVal int) int {
	if val, ex := os.LookupEnv(key); ex {
		return strutil.ToInt(val, defaultVal)
	}
	return defaultVal
}

func GetBool(key string, defaultVal bool) bool {
	if val, ex := os.LookupEnv(key); ex {
		return strutil.ToBool(val, defaultVal)
	}
	return defaultVal
}

func GetFloat32(key string, defaultVal float32) float32 {
	if val, ex := os.LookupEnv(key); ex {
		return strutil.ToFloat32(val, defaultVal)
	}
	return defaultVal
}

func GetFloat64(key string, defaultVal float64) float64 {
	if val, ex := os.LookupEnv(key); ex {
		return strutil.ToFloat64(val, defaultVal)
	}
	return defaultVal
}

func GetStringArray(key string, defaultVal []string) []string {
	if val, ex := os.LookupEnv(key); ex {
		return strings.Split(val, ",")
	}
	return defaultVal
}

func GetInt32Array(key string, defaultVal []int32) []int32 {
	if val, ex := os.LookupEnv(key); ex {
		split := strings.Split(val, ",")
		var result []int32
		for i := range split {
			if intVar, err := strconv.Atoi(split[i]); err == nil {
				result = append(result, int32(intVar))
			}
		}
		return result
	}
	return defaultVal
}

func GetInt64Array(key string, defaultVal []int64) []int64 {
	if val, ex := os.LookupEnv(key); ex {
		split := strings.Split(val, ",")
		var result []int64
		for i := range split {
			if intVar, err := strconv.Atoi(split[i]); err == nil {
				result = append(result, int64(intVar))
			}
		}
		return result
	}
	return defaultVal
}

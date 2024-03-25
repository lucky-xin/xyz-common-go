package collections

func ContainsString(array []string, target string) bool {
	for _, s := range array {
		if s == target {
			return true
		}
	}
	return false
}

func ContainsInt32(array []int32, target int32) bool {
	for _, s := range array {
		if s == target {
			return true
		}
	}
	return false
}

func ContainsInt64(array []int64, target int64) bool {
	for _, s := range array {
		if s == target {
			return true
		}
	}
	return false
}

package pointer

// Ptr returns a pointer to a variable holding the supplied bool constant
func Ptr[T any](x T) *T {
	return &x
}

// PtrVal returns the bool value pointed to by p or defaultVal if p is nil
func PtrVal[T any](p, defaultVal *T) T {
	if p == nil {
		return *defaultVal
	}
	return *p
}

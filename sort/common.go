package sort

type builtInComparable interface {
	string | int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr
}

// Above 大于 排序时递增
func Above[T builtInComparable](a, b T) bool {
	return a > b
}

// Lower 小于 排序时递减
func Lower[T builtInComparable](a, b T) bool {
	return a < b
}

// Equal 相等
func Equal[T comparable](a, b T) bool {
	return a == b
}

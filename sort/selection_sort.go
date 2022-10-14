package sort

// SelectionSort 选择排序
func SelectionSort[T any](a []T, cmp func(a, b T) bool) {
	for i := 0; i < len(a); i++ {
		var pos int
		for j := i + 1; j < len(a); j++ {
			if cmp(a[j], a[i]) {
				pos = j
			}
		}

		a[pos], a[i] = a[i], a[pos]
	}
}

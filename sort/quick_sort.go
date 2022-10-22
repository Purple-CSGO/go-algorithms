package sort

// QuickSort 快速排序
func QuickSort[T any](a []T, left int, right int, cmp func(a, b T) bool) {
	if left < 0 || right >= len(a) {
		return
	}

	if left < right {
		pivot := a[right]

		i := left
		for j := left; j < right; j++ {
			if cmp(pivot, a[j]) {
				a[j], a[i] = a[i], a[j]
				i++
			}
		}

		a[i], a[right] = pivot, a[i]

		QuickSort[T](a, left, i-1, cmp)
		QuickSort[T](a, i+1, right, cmp)
	}
}

// QuickSortAll 快速排序所有元素
func QuickSortAll[T any](a []T, cmp func(a, b T) bool) {
	QuickSort[T](a, 0, len(a)-1, cmp)
}

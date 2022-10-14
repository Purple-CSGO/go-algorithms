package sort

// BubbleSort 冒泡排序
// @Params
//  a   数组
//  cmp 比较方法 a>b ==> true
func BubbleSort[T any](a []T, cmp func(a, b T) bool) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if cmp(a[i], a[j]) {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

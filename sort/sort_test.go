package sort

import (
	"testing"
)

var unsorted = []int{-4, 2, 3, 15, 5, 91, 6657, 114, 514, 1}
var sorted = []int{-4, 1, 2, 3, 5, 15, 91, 114, 514, 6657}

func TestBubbleSort(t *testing.T) {
	var nums []int
	copy(nums, unsorted)
	BubbleSort[int](nums, Above[int])
	for i := range nums {
		if sorted[i] != nums[i] {
			t.Fatalf("BubbleSort test fail\nunsorted:%+v\nnums:%+v\ntarget:%+v\n", unsorted, nums, sorted)
		}
	}
}

func TestQuickSort(t *testing.T) {
	var nums = unsorted

	QuickSortAll[int](nums, Above[int])

	for i := range nums {
		if sorted[i] != nums[i] {
			t.Fatalf("QuickSort test fail\nunsorted:%+v\nnums:%+v\ntarget:%+v\n", unsorted, nums, sorted)
		}
	}
}

func TestSelectionSort(t *testing.T) {
	var nums []int
	copy(nums, unsorted)
	SelectionSort[int](nums, Above[int])
	for i := range nums {
		if sorted[i] != nums[i] {
			t.Fatalf("SelectionSort test fail\nunsorted:%+v\nnums:%+v\ntarget:%+v\n", unsorted, nums, sorted)
		}
	}
}

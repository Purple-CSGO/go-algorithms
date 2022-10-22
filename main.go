package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	// 1. 让数组左侧均<=target 确定边界
	slow, n := 0, len(candidates)
	for fast := 0; fast < n; fast++ {
		if candidates[fast] <= target {
			candidates[slow], candidates[fast] = candidates[fast], candidates[slow]
			slow++
		}
	}

	// 2. 回溯法
	var res [][]int
	comb := []int{}
	var dfs func(target, idx int)

	dfs = func(target, idx int) {
		if idx == slow {
			return
		}
		if target == 0 {
			res = append(res, append([]int(nil), comb...))
			return
		}
		// 直接跳过
		dfs(target, idx+1)
		// 选择当前数
		if target-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)

	return res
}

func main() {
	fmt.Printf("%+v", combinationSum([]int{7, 3, 2}, 18))
}

package main

import "sort"

// 按照起始值排序
func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}
	var res [][]int
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		last := res[len(res)-1]
		current := intervals[i]
		if current[0] <= last[1] {
			if current[1] > last[1] {
				res[len(res)-1][1] = current[1]
			}
		} else {
			res = append(res, current)
		}
	}
	return res
}

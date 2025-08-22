package main

func intersection(nums1 []int, nums2 []int) []int {
	count := make(map[int]int)
	result := make([]int, 0)
	for _, num := range nums1 {
		count[num]++
	}
	for _, num := range nums2 {
		if _, ok := count[num]; ok {
			result = append(result, num)
			delete(count, num)
		}
	}
	return result
}

// []int{1, 2, 2, 1} and []int{2, 2} should return []int{2}
// count[1]=2, count[2]=2
//

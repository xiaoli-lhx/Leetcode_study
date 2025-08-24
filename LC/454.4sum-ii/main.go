package main

//	func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
//		m := make(map[int]int)
//		n := len(nums1)
//		count := 0
//		for i := 0; i < n; i++ {
//			for j := 0; j < n; j++ {
//				m[nums1[i]+nums2[j]]++
//			}
//		}
//		for i := 0; i < n; i++ {
//			for j := 0; j < n; j++ {
//				if _, ok := m[0-(nums3[i]+nums4[j])]; ok {
//					count += m[0-(nums3[i]+nums4[j])]
//				}
//			}
//		}
//
//		return count
//	}
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := make(map[int]int, len(nums1)*len(nums2))
	count := 0
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			m[v1+v2]++
		}
	}
	for _, v1 := range nums3 {
		for _, v2 := range nums4 {
			count += m[-(v1 + v2)]
		}
	}
	return count
}

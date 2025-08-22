package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	// 步骤一：对数组进行排序，这是双指针法的基础。
	sort.Ints(nums)
	var result [][]int
	n := len(nums)

	// 步骤二：固定前两个数 nums[i] 和 nums[j]。
	// 第一个循环，固定 nums[i]。
	// 我们只需要遍历到倒数第四个元素，因为后面需要至少三个元素。
	for i := 0; i < n-3; i++ {
		// 避免重复：如果 nums[i] 和前一个数相同，则跳过。
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 第二个循环，固定 nums[j]。
		// 我们只需要遍历到倒数第三个元素。
		for j := i + 1; j < n-2; j++ {
			// 避免重复：如果 nums[j] 和前一个数相同，则跳过。
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			// 步骤三：使用双指针法寻找另外两个数。
			// left 指向 j 的下一个元素，right 指向数组末尾。
			left, right := j+1, n-1

			for left < right {
				// 计算当前四个数的和
				sum := nums[i] + nums[j] + nums[left] + nums[right]

				if sum == target {
					// 找到了一个符合条件的四元组。
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})

					// 避免重复：移动 left 和 right 指针，直到它们指向的元素与当前元素不同。
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}

					// 正常移动指针继续寻找其他组合。
					left++
					right--

				} else if sum < target {
					// 如果和小于目标值，说明需要一个更大的数，移动 left 指针。
					left++
				} else { // sum > target
					// 如果和大于目标值，说明需要一个更小的数，移动 right 指针。
					right--
				}
			}
		}
	}

	return result
}

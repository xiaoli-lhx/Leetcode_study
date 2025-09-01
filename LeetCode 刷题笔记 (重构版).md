# `LeetCode` 刷题笔记 (重构版)

这是一份根据练习记录重构的刷题笔记，主要分为**核心算法思想**、**基础数据结构**和**设计类问题**三大模块，旨在提供一个更清晰、更系统化的复习框架。

## 目录

[TOC]

## 一、核心算法思想

### 双指针 (Two Pointers)

双指针是一种通过维护两个指针在序列中同向或相向移动，来降低时间复杂度的技巧。

#### 88. [合并两个有序数组]()

**思路**：**从后往前**进行合并。因为 `nums1` 的尾部是空闲空间，从后往前填充，我们永远不会覆盖掉 `nums1` 中尚未被处理的有效元素。

这个方法被称为**逆向双指针法**。

~~~go
package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	// nums1 []int{2,3,4,0,0,0}
	// m int 3
	// nums2 []int{1,5,6}
	// n int 3
	// Output: nums1 []int{1,2,3,4,5,6}
	i := m - 1
	j := n - 1
	k := m + n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
    // 如果nums2还有未被处理的数，那么一定是小于了nums1的最小的数，仅需放在nums1前面即可
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}
~~~

#### 125. [验证回文串](https://leetcode.cn/problems/valid-palindrome/)

**思路**：**初始化指针**：定义两个指针，`left` 指向字符串的开头，`right` 指向字符串的末尾。

**循环比较**：当 `left` 指针在 `right` 指针左边时，进行循环。

- **跳过无效字符**：从左向右移动 `left` 指针，直到它指向一个字母或数字。同样，从右向左移动 `right` 指针，直到它也指向一个字母或数字。
- **比较字符**：将 `left` 和 `right` 指针所指向的字符统一转换为小写进行比较。
  - 如果两个字符不相等，说明这个字符串不是回文串，直接返回 `false`。
  - 如果相等，则将 `left` 指针向右移动一位，`right` 指针向左移动一位，继续下一轮比较。

**得出结论**：如果循环正常结束（即 `left` 越过了 `right`），说明所有对应的字符都相等，该字符串是回文串，返回 `true`。

~~~go
func isPalindrome(s string) bool {
    left,right:=0,len(s)-1
    for left<right{
        // unicode.IsLetter() 判断是否为字母
        // unicode.IsDigit() 判断是否为数字
        for left<right && !unicode.IsLetter(rune(s[left])) && !unicode.IsDigit(rune(s[left])){
            left++
        }
        for left<right && !unicode.IsLetter(rune(s[right])) && !unicode.IsDigit(rune(s[right])){
            right--
        }
        // unicode.ToLower() 转化为小写字母
        if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])){
            return false
        }
        left++
        right--
    }
    return true
}
~~~

#### 26. [删除有序数组中的重复项](https://leetcode.cn/problems/remove-duplicates-from-sorted-array/)

**思路**：使用快慢指针（从1开始，数组中的第一个元素不可能是重复的），快指针遍历数组找新的不重复的元素，慢指针指向下一个唯一元素应该被放置的位置。它也代表了到目前为止唯一元素的数量。

~~~go
func removeDuplicates(nums []int) int {
	n := len(nums)
	// 数组第一个元素肯定是唯一的
	slow := 1
	for fast := 1; fast < n; fast++ {
		// 如果当前元素和上一个元素不相同，则将当前元素复制到慢指针位置
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
~~~

#### [11. 盛水最多的容器](https://leetcode.cn/problems/container-with-most-water/description/)

**思路:** 使用左右双指针，分别指向数组首尾。每次计算面积并更新最大值，然后将指向**较短板**的指针向内移动，因为移动长板不可能使面积增大。

```
func maxArea(height []int) int {
    maxArea := 0
    left, right := 0, len(height)-1
    for left < right {
        area := min(height[left], height[right]) * (right - left)
        if area > maxArea {
            maxArea = area
        }
        if height[left] > height[right] {
            right--
        } else {
            left++
        }
    }
    return maxArea
}
```

#### [167. 两数之和II-输入有序数组](https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/description/)

**思路:** 利用数组有序性，使用左右双指针。若两数之和小于 `target`，左指针右移；若大于 `target`，右指针左移。

```
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}
```

#### [283. 移动零](https://leetcode.cn/problems/move-zeroes/description/)

**思路:** 快慢指针。`slow` 指针指向下一个非零元素应存放的位置，`fast` 指针遍历数组。当 `nums[fast]` 不为零时，将其值与 `nums[slow]` 交换并移动 `slow`。

```go
func moveZeroes(nums []int) {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}
```

#### [27. 移除元素](https://leetcode.cn/problems/remove-element/description/)

**思路:** 快慢指针。`slow` 指针记录非 `val` 元素的个数，`fast` 遍历数组。当 `nums[fast]` 不等于 `val` 时，将其值赋给 `nums[slow]` 并移动 `slow`。

```
func removeElement(nums []int, val int) int {
    slow := 0
    for fast := 0; fast < len(nums); fast++ {
        if nums[fast] != val {
            nums[slow] = nums[fast]
            slow++
        }
    }
    return slow
}
```

#### [977. 有序数组的平方](https://leetcode.cn/problems/squares-of-a-sorted-array/description/)

**思路:** 双指针法。由于原数组有序，平方后的最大值一定在原数组的两端。设置左右指针分别指向首尾，比较两者平方大小，将较大者从后往前放入新数组。

```
func sortedSquares(nums []int) []int {
   n := len(nums)
   result := make([]int, n)
   left, right, k := 0, n-1, n-1
   for left <= right {
       lm, rm := nums[left]*nums[left], nums[right]*nums[right]
       if lm > rm {
           result[k] = lm
           left++
       } else {
           result[k] = rm
           right--
       }
       k--
   }
   return result
}
```

#### [15. 三数之和](https://leetcode.cn/problems/3sum/)

**思路:** 排序 + 双指针。先对数组排序，然后遍历数组，固定一个数 `nums[i]`，再用双指针 `left` 和 `right` 在 `i` 之后的区间内寻找和为 `-nums[i]` 的两个数。注意处理重复解。

```
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	n := len(nums)
	for i := 0; i < n-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}
```

#### [18. 四数之和](https://leetcode.cn/problems/4sum/)

**思路:** 排序 + 双指针。与三数之和类似，固定两个数 `nums[i]` 和 `nums[j]`，然后使用双指针在 `j` 之后的区间内寻找另外两个数。注意剪枝和去重。

```
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int
	n := len(nums)
	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left, right := j+1, n-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if sum > target {
					right--
				} else {
					left++
				}
			}
		}
	}
	return res
}
```

### 滑动窗口 (Sliding Window)

滑动窗口是双指针的一种特例，用于解决子数组/子串问题。两个指针维护一个“窗口”，根据条件移动右指针扩大窗口，移动左指针缩小窗口。

#### 76.[最小覆盖字串](https://leetcode.cn/problems/minimum-window-substring/)

##### **思路**:滑动窗口（欠债与还债模型）

可以把这个问题想象成一个“还债”的过程：

1. 字符串 `t` 是我们的**“债务清单”(`need` map)**，记录了我们需要“偿还”哪些字符，以及各需要多少个。
2. 我们在字符串 `s` 上移动的**窗口 (`window` map)**，就是我们用来“还债”的资产。
3. 算法的目标是，找到能**刚好还清所有债务**的**最小**的一段连续资产。

###### 关键变量

- `need` 哈希表：债务清单，记录 `t` 中各字符的数量。
- `window` 哈希表：记录当前窗口内，相关字符的数量。
- `left`, `right` 指针：构成滑动窗口的左右边界。
- `match` 计数器：记录**已还清**的字符种类数。

###### 算法流程：两阶段循环

整个过程就是 `right` 指针不断右移，而 `left` 指针在满足条件时右移。

**第一阶段：扩大窗口，努力“还债”**

- `right` 指针向右移动，把新字符纳入 `window`。
- 如果新纳入的字符是 `need` 中的一员：
  - `window` 中对应字符数+1。
  - 如果 `window` 中该字符的数量**恰好等于** `need` 中的数量，说明该种字符的“债务”刚好还清，`match` 计数器+1。

**第二阶段：收缩窗口，寻求“最优”**

- 一旦 `match` 的值等于 `need` 中不同字符的总数，说明**所有债务都已还清**，找到了一个可行的解。
- 此时，我们尝试**收缩窗口**：
  - `left` 指针向右移动，将字符移出 `window`。
  - 如果移出的字符是 `need` 中的一员：
    - 如果它在 `window` 中的数量**从“刚好还清”变为“又欠下了”**，那么 `match` 计数器-1。
    - `window` 中对应字符数-1。
  - 在每次成功收缩时（即 `left` 右移前），都记录并更新一次“最小覆盖子串”的长度和位置。
  - 当 `match` 值不再满足条件时，停止收缩，回到第一阶段，继续移动 `right` 寻找下一个还清债务的时刻。

**总结**：整个算法就是 `right` 指针一路向东探索，`left` 指针则在窗口满足条件时，尽力向右收缩以找到最短的距离

~~~Go
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	
	// need 记录了 t 中每个字符需要的数量
	need := make(map[byte]int)
	// window 记录了当前窗口中每个字符的数量
	window := make(map[byte]int)

	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left, right := 0, 0
	// match 记录了窗口中满足 need 条件的字符种类数
	match := 0
	// minLen 记录了最小子串的长度，初始化为一个不可能的大值
	minLen := len(s) + 1
	// start 记录了最小子串的起始位置
	start := 0

	for right < len(s) {
		// c1 是即将移入窗口的字符
		c1 := s[right]
		right++ // 扩大窗口

		// --- 更新窗口和匹配状态 ---
		if _, ok := need[c1]; ok {
			window[c1]++
			if window[c1] == need[c1] {
				// 这种字符的数量已经满足要求了
				match++
			}
		}

		// --- 当所有字符都满足要求时，开始收缩窗口 ---
		for match == len(need) {
			// 找到了一个有效的覆盖子串，更新最小长度
			if right-left < minLen {
				minLen = right - left
				start = left
			}

			// c2 是即将移出窗口的字符
			c2 := s[left]
			left++ // 收缩窗口

			// --- 更新窗口和匹配状态 ---
			if _, ok := need[c2]; ok {
				if window[c2] == need[c2] {
					// 移出 c2 后，这种字符的数量将不再满足要求
					match--
				}
				window[c2]--
			}
		}
	}

	// 如果 minLen 没有被更新过，说明没有找到有效的子串
	if minLen == len(s)+1 {
		return ""
	}
	
	return s[start : start+minLen]
}

~~~



#### 3. [无重复字符的最长字串](https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/)

**思路**:滑动窗口 + 哈希表

想象有一个**可伸缩的窗口**在字符串上从左向右滑动。这个窗口始终努力保持内部没有重复字符，并在这个过程中记录下它所能达到的最大宽度。

具体步骤：

1. **定义窗口**：用 `left` 和 `right` 两个指针作为窗口的左右边界。`right` 负责向右探索，`left` 在必要时向右收缩。

2. **定义“记事本”**：使用一个哈希表（`map`）作为记事本，记录每个字符**最近一次出现的位置**。

3. **滑动过程**：

   - `right` 指针不断向右移动，考察新字符。

   - **遇到新字符时，问记事本**：这个字符上次出现过吗？

     - **如果出现过，且在当前窗口内**：说明窗口内有了重复。必须收缩窗口。将 `left` 指针直接“跳”到**重复字符上次出现位置的下一位**。

     - **如果不重复**：窗口可以安全地向右扩大。`left` 指针不动。

   - 每移动一步 `right`，都计算一下当前窗口的长度 (`right - left + 1`)，并更新记录到的最大长度。

   - **更新记事本**：在每一步结束时，都要在记事本上更新当前字符的最新位置。

总结

这个方法就像一个“健忘”的检查员，他只关心当前窗口内的字符是否重复。一旦发现重复，他就把窗口的起点移到旧字符之后，假装“忘记”了旧字符以及它之前的所有内容，然后继续检查。通过这种方式，我们用一次遍历就解决了问题，效率非常高。

~~~go
func lengthOfLongestSubstring(s string) int {
	// lastSeen 哈希表：key 是字符，value 是该字符最近一次出现的索引
	lastSeen := make(map[rune]int)
	maxLength := 0
	// left 是滑动窗口的左边界
	left := 0

	// right 是滑动窗口的右边界
	for right, char := range s {
		// 检查当前字符 char 是否在哈希表中，并且其上次出现的位置是否在当前窗口内
		if lastIndex, found := lastSeen[char]; found && lastIndex >= left {
			// 如果是，说明遇到了重复字符。
			// 我们需要收缩窗口，将左边界移动到重复字符上次出现位置的下一个位置。
			left = lastIndex + 1
		}

		// 计算当前窗口的长度
		currentLength := right - left + 1
		// 更新最大长度
		if currentLength > maxLength {
			maxLength = currentLength
		}

		// 无论如何，都要更新当前字符的最新位置
		lastSeen[char] = right
	}

	return maxLength
}
~~~



#### [209. 长度最小的子数组](https://leetcode.cn/problems/minimum-size-subarray-sum/description/)

**思路:** 使用滑动窗口。用 `right` 指针扩大窗口，当窗口内元素和 `sum >= target` 时，记录长度并收缩 `left` 指针，直到 `sum < target`，然后继续扩大窗口。

```
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	left, sum, minlen := 0, 0, n+1
	for right := 0; right < n; right++ {
		sum += nums[right]
		for sum >= target {
			count := right - left + 1
			if count < minlen {
				minlen = count
			}
			sum -= nums[left]
			left++
		}
	}
	if minlen == n+1 {
		return 0
	}
	return minlen
}
```

### 二分查找 (Binary Search)

适用于**有序**序列的查找，每次将搜索范围缩小一半，时间复杂度为 O(log n)。

#### [704. 二分查找](https://leetcode.cn/problems/binary-search/description/)

```
func search(nums []int, target int) int {
    right := len(nums) - 1
    left := 0
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}
```

### 贪心算法 (Greedy Algorithm)

**核心思想:** 贪心的本质是选择每一阶段的局部最优，从而达到全局最优。

#### [455. 分发饼干](https://leetcode.cn/problems/assign-cookies/description/)

**贪心策略:** 为了满足更多孩子，应该用尺寸尽量小的饼干去满足胃口最小的孩子，或者用尺寸最大的饼干去满足胃口最大的孩子。

```
func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)
    index := len(s) - 1
    count := 0
    for i := len(g) - 1; i >= 0; i-- {
        if index >= 0 && s[index] >= g[i] {
            count++
            index--
        }
    }
    return count
}
```

#### [122. 买卖股票的最佳时机Ⅱ](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/)

**贪心策略:** 只要明天的价格比今天高，就今天买入，明天卖出，将所有正利润累加起来。

```
func maxProfit(prices []int) int {
    profit := 0
    for i := 0; i < len(prices)-1; i++ {
        if prices[i] < prices[i+1] {
            profit += prices[i+1] - prices[i]
        }
    }
    return profit
}
```

#### [376. 摆动序列](https://leetcode.cn/problems/wiggle-subsequence/description/)

**贪心策略:** 统计数组中的“峰”和“谷”的数量。在单调递增或递减的序列中，只需要保留端点即可。

```
func wiggleMaxLength(nums []int) int {
    if len(nums) < 2 {
        return len(nums)
    }
    up, down := 1, 1
    for i := 1; i < len(nums); i++ {
        if nums[i]-nums[i-1] > 0 {
            up = down + 1
        } else if nums[i] < nums[i-1] {
            down = up + 1
        }
    }
    if up > down {
        return up
    }
    return down
}
```

#### [53. 最大子数组和](https://leetcode.cn/problems/maximum-subarray/description/)

**贪心策略 (Kadane's Algorithm):** 如果当前累加的子数组和为负数，那么它对后续的结果只会产生负贡献，应该舍弃它，从下一个元素重新开始累加。

```
func maxSubArray(nums []int) int {
    currentSum := 0
    maxSum := nums[0]
    for i := 0; i < len(nums); i++ {
        currentSum += nums[i]
        if currentSum > maxSum {
            maxSum = currentSum
        }
        if currentSum < 0 {
            currentSum = 0
        }
    }
    return maxSum
}
```

## 二、基础数据结构

### 数组 (Array)

#### 238.[除自身以外数组的乘积](https://leetcode.cn/problems/product-of-array-except-self/)

**思路**：前缀积与后缀积

既然不能用除法，我们就不能先计算总乘积再逐个除去。正确的思路是，对于每一个位置 `i`，它最终的结果其实是**它左边所有元素的乘积** 乘以 **它右边所有元素的乘积**。

~~~go
func productExceptSelf(nums []int) []int {
	n := len(nums)
	// 初始化结果数组
	answer := make([]int, n)

	// --- 第一次遍历：计算前缀积 ---
	// answer[i] 先用来存储 nums[i] 左侧所有元素的乘积
	// 对于第一个元素，其左侧没有元素，我们认为乘积为 1
	answer[0] = 1
	for i := 1; i < n; i++ {
		// i 的前缀积 = (i-1 的前缀积) * nums[i-1]
		answer[i] = answer[i-1] * nums[i-1]
	}

	// --- 第二次遍历：计算后缀积并得出最终结果 ---
	// 我们需要一个变量来从右到左地累计后缀积
	// 对于最后一个元素，其右侧没有元素，初始后缀积为 1
	suffixProduct := 1
	for i := n - 1; i >= 0; i-- {
		// 对于位置 i，最终结果 = (之前算好的前缀积) * (当前的后缀积)
		answer[i] = answer[i] * suffixProduct
		// 更新后缀积，为下一个位置（左边一个）做准备
		suffixProduct = suffixProduct * nums[i]
	}

	return answer
}
~~~



#### [66. 加一](https://leetcode.cn/problems/plus-one/description/)

**思路:** 从数组末尾开始遍历，处理进位。如果所有位都是9，则需要在数组头部加一个1。

```
func plusOne(digits []int) []int {
    n := len(digits) - 1
    for i := n; i >= 0; i-- {
        if digits[i] < 9 {
            digits[i]++
            return digits
        }
        digits[i] = 0
    }
    return append([]int{1}, digits...)
}
```

#### [121. 买卖股票的最佳时机](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/)

**思路:** 一次遍历，记录迄今为止的最低价格 `minPrice`，并不断计算当前价格与 `minPrice` 的差值来更新最大利润 `maxProfit`。

```
func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }
    minPrice := prices[0]
    maxProfit := 0
    for i := 0; i < len(prices); i++ {
        if minPrice > prices[i] {
            minPrice = prices[i]
        } else {
            profit := prices[i] - minPrice
            if maxProfit < profit {
                maxProfit = profit
            }
        }
    }
    return maxProfit
}
```

#### [136. 只出现一次的数字](https://leetcode.cn/problems/single-number/)

**思路:** 使用异或运算(XOR)。任何数与自身异或为0，任何数与0异或为自身。所有数字异或一遍，成对的数字会抵消，最后剩下只出现一次的数字。

```
func singleNumber(nums []int) int {
    result := 0
    for _, num := range nums {
        result ^= num
    }
    return result
}
```

#### [59. 螺旋矩阵II](https://leetcode.cn/problems/spiral-matrix-ii/description/)

**思路:** 模拟螺旋填充过程。定义 `top`, `bottom`, `left`, `right` 四个边界，按“从左到右 -> 从上到下 -> 从右到左 -> 从下到上”的顺序填充，每完成一圈就收缩相应的边界。

```
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	num := 1
	top, bottom, left, right := 0, n-1, 0, n-1
	for num <= n*n {
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--
		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}
		bottom--
		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++
	}
	return matrix
}
```

### 链表 (Linked List)

#### [707. 设计链表](https://leetcode.cn/problems/design-linked-list/description/)

**问题描述:** 设计并实现自己的链表，支持 `get`, `addAtHead`, `addAtTail`, `addAtIndex`, `deleteAtIndex` 等操作。

```
type MyLinkedList struct {
	dummyHead *ListNode // 虚拟头结点
	size      int       // 链表长度
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		dummyHead: &ListNode{},
		size:      0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	current := this.dummyHead.Next
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.size {
		return
	}
	newNode := &ListNode{Val: val}
	prev := this.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	newNode.Next = prev.Next
	prev.Next = newNode
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	prev := this.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	prev.Next = prev.Next.Next
	this.size--
}
```

#### [203. 移除链表元素](https://leetcode.cn/problems/remove-linked-list-elements/)

**思路:** 使用虚拟头节点（Dummy Head）统一并简化删除逻辑，特别是处理头节点需要被删除的情况。

```
func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Next: head}
	prev := dummyHead
	curr := head
	for curr != nil {
		if curr.Val == val {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
		curr = curr.Next
	}
	return dummyHead.Next
}
```

#### [206. 反转链表](https://leetcode.cn/problems/reverse-linked-list/description/)

**迭代法:**

```
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		nextTmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTmp
	}
	return prev
}
```

**递归法:**

```
func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
```

#### [24. 两两交换链表中的节点](https://leetcode.cn/problems/swap-nodes-in-pairs/)

**迭代法:** 使用虚拟头节点，三个指针 `cur`, `first`, `second` 配合完成交换，`cur` 每次前进两步。

```
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		first := cur.Next
		second := cur.Next.Next
		first.Next = second.Next
		second.Next = first
		cur.Next = second
		cur = first
	}
	return dummy.Next
}
```

#### [19. 删除链表的倒数第N个节点](https://leetcode.cn/problems/remove-nth-node-from-end-of-list/description/)

**思路:** 快慢指针。快指针先走 `n+1` 步，然后快慢指针同步前进。当快指针到达 `nil` 时，慢指针正好指向待删除节点的前一个节点。

```
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
```

#### [面试题 02.07. 链表相交](https://leetcode.cn/problems/intersection-of-two-linked-lists-lcci/)

**思路:** 双指针法。两个指针 `pA`, `pB` 分别遍历链表 A 和 B。当一个指针到达末尾时，它就跳转到另一个链表的头部继续遍历。这样，两个指针走过的总路程相等 (`lenA+lenB`)，如果存在交点，它们必会在交点相遇。

```
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }
    pA, pB := headA, headB
    for pA != pB {
        if pA == nil {
            pA = headB
        } else {
            pA = pA.Next
        }
        if pB == nil {
            pB = headA
        } else {
            pB = pB.Next
        }
    }
    return pA
}
```

#### [142. 环形链表 II](https://leetcode.cn/problems/linked-list-cycle-ii/)

**思路:** 快慢指针。快指针一次走两步，慢指针一次一步。若相遇，则有环。相遇后，将一个指针放回头节点，另一个指针留在相遇点，然后两个指针都以一次一步的速度前进，再次相遇的点即为环的入口。

```
func detectCycle(head *ListNode) *ListNode {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
        if slow == fast {
            for slow != head {
                slow = slow.Next
                head = head.Next
            }
            return head
        }
    }
    return nil
}
```

### 字符串 (String)

#### [344. 反转字符串](https://leetcode.cn/problems/reverse-string/description/)

**思路:** 双指针，一个指向头，一个指向尾，两两交换直到相遇。

```
func reverseString(s []byte) {
    n := len(s)
    for i := 0; i < n/2; i++ {
        s[i], s[n-i-1] = s[n-i-1], s[i]
    }
}
```

#### [541. 反转字符串II](https://leetcode.cn/problems/reverse-string-ii/description/)

**思路:** 遍历字符串，步长为 `2k`。在每个 `2k` 的块内，反转前 `k` 个字符。注意处理末尾不足 `2k` 或 `k` 个字符的边界情况。

```
func reverseStr(s string, k int) string {
	bytes := []byte(s)
	n := len(bytes)
	for i := 0; i < n; i += 2 * k {
		left := i
		right := i + k - 1
		if right >= n {
			right = n - 1
		}
		for left < right {
			bytes[left], bytes[right] = bytes[right], bytes[left]
			left++
			right--
		}
	}
	return string(bytes)
}
```

#### [151. 反转字符串中的单词](https://leetcode.cn/problems/reverse-words-in-a-string/)

**库函数解法:**

```
import "strings"

func reverseWords(s string) string {
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}
```

#### [28. 找出字符串中第一个匹配项的下标](https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/description/)

**暴力匹配:**

```
func strStr(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}
	for i := 0; i <= n-m; i++ {
		if haystack[i:i+m] == needle {
			return i
		}
	}
	return -1
}
```

### 哈希表 (Hash Table)

#### 560.  [和为K的子数组](https://leetcode.cn/problems/subarray-sum-equals-k/)

**思路**：前缀和＋哈希表，**当前前缀和 - 起始前缀和 = k --> 当前前缀和 - k = 起始前缀和**

>比如，对于数组 `nums = [3, 4, 7, 2]`：
>
>- 到第一个数 `3` 的前缀和是 `3`。
>- 到第二个数 `4` 的前缀和是 `3 + 4 = 7`。
>- 到第三个数 `7` 的前缀和是 `7 + 7 = 14`。
>- 到第四个数 `2` 的前缀和是 `14 + 2 = 16`。
>
>这个“累计尺”最大的用处，就是能立刻算出**任意一段**的和。比如，我们要算子数组 `[4, 7]` 的和（也就是从第2个到第3个数）。我们可以用第3个位置的前缀和 (14) 减去 **它开始位置之前** 的前缀和 (3)，得到 `14 - 3 = 11`。这和 `4 + 7` 的结果是一样的！

当我们在数组中一路计算“当前前缀和”（也就是公式里的“终点的前缀和”）时，我们只需要回头看看，**我们以前见没见过一个前缀和，它的值恰好是“当前前缀和 - k”**。

如果见过，那么从那个“见过的位置”到我们“当前的位置”，中间夹着的子数组的和就一定是 `k`！

**现在问题来了：** 当我们遍历数组时，如何能快速地知道“某个值的前缀和”我们以前见没见过，以及见过几次呢？

这就是**哈希表 (Map)** 登场的时候了。我们可以用一个哈希表，在遍历的时候不断记录：

- **键 (Key)**: 某个前缀和的值。
- **值 (Value)**: 这个值到目前为止出现过几次。

这样，每当我们算出一个新的“当前前缀和”，我们就可以去哈希表里快速查询 `当前前缀和 - k` 是不是存在，从而瞬间知道有多少个满足条件的子数组。

~~~go
func subarraySum(nums []int, k int) int {
	count := 0
	sum := 0
	m := make(map[int]int)
	m[0] = 1
	for _, num := range nums {
		sum += num
		if _, ok := m[sum-k]; ok {
			count++
		}
		m[sum]++
	}
	return count
}
~~~



#### [1. 两数之和](https://leetcode.cn/problems/two-sum/)

**思路:** 使用哈希表存储已经遍历过的数字及其下标。对于每个数字，检查 `target - num` 是否在哈希表中。

```
func twoSum(nums []int, target int) []int {
    m := map[int]int{}
    for i, num := range nums {
        if v, ok := m[target-num]; ok {
            return []int{v, i}
        }
        m[num] = i
    }
    return nil
}
```

#### [217. 存在重复元素](https://leetcode.cn/problems/contains-duplicate/)

**思路:** 使用哈希集合（`map[int]bool` 或 `map[int]struct{}`）记录出现过的数字。

```
func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
    for _, num := range nums {
        if m[num] {
            return true
        }
        m[num] = true
    }
    return false
}
```

#### [219. 存在重复元素Ⅱ](https://leetcode.cn/problems/contains-duplicate-ii/description/)

**思路:** 使用哈希表记录数字及其最近出现的下标。

```
func containsNearbyDuplicate(nums []int, k int) bool {
    m := map[int]int{}
    for i, num := range nums {
        if lastIndex, ok := m[num]; ok {
            if i-lastIndex <= k {
                return true
            }
        }
        m[num] = i
    }
    return false
}
```

#### [242. 有效的字母异位词](https://leetcode.cn/problems/valid-anagram/description/)

**思路:** 使用一个数组或哈希表记录 `s` 中每个字符的出现频率，然后遍历 `t` 进行抵消。

```
// 仅包含小写字母
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := [26]int{}
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if count[i] != 0 {
			return false
		}
	}
	return true
}
```

#### [349. 两个数组的交集](https://leetcode.cn/problems/intersection-of-two-arrays/)

**思路:** 用一个哈希集合存储 `nums1` 的所有元素，然后遍历 `nums2`，检查元素是否存在于集合中。

```
func intersection(nums1 []int, nums2 []int) []int {
    set := make(map[int]struct{})
    result := make([]int, 0)
    for _, num := range nums1 {
        set[num] = struct{}{}
    }
    for _, num := range nums2 {
        if _, ok := set[num]; ok {
            result = append(result, num)
            delete(set, num)
        }
    }
    return result
}
```

#### [202. 开心数](https://leetcode.cn/problems/happy-number/)

**思路:** 在计算过程中，用哈希表记录出现过的数字。如果某个数字重复出现，说明进入了循环。

```
func isHappy(n int) bool {
	seen := make(map[int]bool)
	for n != 1 && !seen[n] {
		seen[n] = true
		n = getNext(n)
	}
	return n == 1
}
func getNext(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}
```

#### [454. 四数相加II](https://leetcode.cn/problems/4sum-ii/description/)

**思路:** 分组+哈希表。先计算 `nums1` 和 `nums2` 中所有元素的和，存入哈希表。再遍历 `nums3` 和 `nums4`，查找 `target - (v3+v4)` 是否在哈希表中。

```
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := make(map[int]int)
	count := 0
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			m[v1+v2]++
		}
	}
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			count += m[-(v3 + v4)]
		}
	}
	return count
}
```

#### [383. 赎金信](https://leetcode.cn/problems/ransom-note/description/)

**思路:** 用哈希表或数组统计 `magazine` 中每个字符的数量，然后遍历 `ransomNote`，消耗相应字符的计数。

```
func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	charCount := make(map[rune]int)
	for _, c := range magazine {
		charCount[c]++
	}
	for _, c := range ransomNote {
		if charCount[c] > 0 {
			charCount[c]--
		} else {
			return false
		}
	}
	return true
}
```

### 栈与队列(Stack and Queue)

#### 347.[前k个高频元素](https://leetcode.cn/problems/top-k-frequent-elements/)

**思路**: 详见[Go语言解[前k个高频元素]：从排序到堆的深度探索](https://xiaoli-lhx.github.io/go%E8%AF%AD%E8%A8%80/go%E5%9F%BA%E7%A1%80/go%E8%AF%AD%E8%A8%80%E5%A0%86/)

~~~go
package main

import (
	"container/heap"
	"sort"
)

type Pair struct {
	Number int
	Count  int
}

type IHeap []Pair

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i].Count < h[j].Count }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}
func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	// 1. 统计频率
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// 2. 创建一个最小堆
	minHeap := &IHeap{}
	heap.Init(minHeap) // 初始化堆

	// 3. 遍历频率，维护一个大小为 k 的最小堆
	for num, count := range freqMap {
		heap.Push(minHeap, Pair{Number: num, Count: count})
		if minHeap.Len() > k {
			heap.Pop(minHeap) // 如果堆大小超过 k，就把最小的那个（堆顶）扔掉
		}
	}

	// 4. 堆里剩下的就是前 k 个高频元素，收集结果
	var result []int
	for minHeap.Len() > 0 {
		// Pop 的结果是 Pair 类型，我们需要它的 Number 字段
		result = append(result, heap.Pop(minHeap).(Pair).Number)
	}
	return result
}

// 另一种实现方式，使用 sort.Slice 进行自定义降序排序
func topKFrequent_Sort(nums []int, k int) []int {
	// 1. 统计频率
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// 2. 将 map 转换为 struct 切片
	type Pair struct {
		Number int
		Count  int
	}
	var pairs []Pair
	for num, count := range freqMap {
		pairs = append(pairs, Pair{Number: num, Count: count})
	}

	// 3. 使用 sort.Slice 进行自定义降序排序
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Count > pairs[j].Count
	})

	// 4. 取出前 k 个元素
	var result []int
	for i := 0; i < k; i++ {
		result = append(result, pairs[i].Number)
	}
	return result
}

~~~

#### 239.[滑动窗口最大值](https://leetcode.cn/problems/sliding-window-maximum/)

**思路**：核心思路是**避免重复计算**。我们不关心窗口里的所有元素，只关心那些“有潜力”成为最大值的“候选人”。

我们用一个特殊的数据结构——**单调队列**（用双端队列或切片实现）来维护这些候选人，并遵循以下三条规则：

1. **维护单调性**：当一个新元素准备入队时，我们从**队尾**把所有比它小的旧元素都“挤”出去，再让新元素入队。这样保证队列从头到尾是单调递减的。
2. **获取最大值**：因为队列是单调递减的，所以**队头永远是当前窗口的最大值**。
3. **处理过期元素**：当窗口滑动时，我们需要检查队头的元素是否已经“滑出”了窗口范围，如果是，就把它从**队头**移除。

通过这三步，我们在每次窗口滑动时，都能以近乎 O(1) 的时间找到最大值，最终总时间复杂度降到了 O(N)。

~~~go
func maxSlidingWindow(nums []int, k int) []int {
	// 1. 准备工作
	// 用一个切片来做双端队列，存放索引
	deque := []int{}
	result := []int{}
	// 2. 遍历数组
	for i := 0; i < len(nums); i++ {
		// -清理队尾
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= nums[i] {
			deque = deque[:len(deque)-1]
		}
		// -加入队尾
		deque = append(deque, i)
		// -清理队头
		if deque[0] <= i-k {
			deque = deque[1:]
		}
		// -加入结果
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}
	return result
}
~~~



#### 150.[逆波兰表达式求值](https://leetcode.cn/problems/evaluate-reverse-polish-notation/description/)

**思路**：

创建一个空栈（在Go里可以用切片实现）。

从左到右遍历表达式中的每一个元素（字符串）：

- 如果这个元素是一个**数字**，就把它转换成整型，然后**入栈**。
- 如果这个元素是一个**运算符** (`+`, `-`, `*`, `/`)：
  - 先从栈顶**出栈**一个数（比如 `num1`）。
  - 再从栈顶**出栈**第二个数（比如 `num2`）。
  - 执行 `num2 [运算符] num1` 的计算。
  - 把计算结果**入栈**。

遍历结束后，栈里会只剩下一个数字，它就是最终的答案。

~~~go
package main

import "strconv"

func evalRPN(tokens []string) int {
	// 我们的栈，用来存放整数
	stack := []int{}

	// 遍历每一个字符串 token
	for _, token := range tokens {
		// 尝试把 token 转换成数字
		num, err := strconv.Atoi(token)

		// 如果转换成功 (err == nil)，说明是数字
		if err == nil {
			// 数字直接入栈
			stack = append(stack, num)
		} else { // 如果转换失败，说明是运算符
			// 从栈顶弹出两个数字
			num1 := stack[len(stack)-1]
			num2 := stack[len(stack)-2]
			// 先把弹出的两个数从栈里移除
			stack = stack[:len(stack)-2]

			// 根据 token (运算符) 进行计算
			switch token {
			case "+":
				stack = append(stack, num2+num1)
			case "-":
				stack = append(stack, num2-num1)
			case "*":
				stack = append(stack, num2*num1)
			case "/":
				stack = append(stack, num2/num1)
			}
		}
	}

	// 循环结束后，栈里剩下的唯一一个数就是答案
	return stack[0]
}

~~~



#### 1047.[删除字符串中所有的相邻重复项](https://leetcode.cn/problems/remove-all-adjacent-duplicates-in-string/)

**思路：**

创建一个空栈。

从左到右遍历字符串中的每一个字符：

- 首先，检查栈是否为空。
  - 如果栈是空的（比如处理第一个字符时），直接将当前字符**入栈**。
  - 如果栈不是空的，就查看一下栈顶的元素。
- 将当前字符与栈顶元素进行比较：
  - 如果它们**相同**，就将栈顶元素**出栈**（实现“消除”效果）。
  - 如果它们**不同**，就将当前字符**入栈**。

遍历完整个字符串后，栈里剩下的字符，就是最终“消除”后留下的结果。我们只需要把它们从栈底到栈顶依次拼接起来即可。

**用list模拟栈**：

~~~go
func removeDuplicates(s string) string {
	stack := list.New()
	stack.PushBack(s[0])
	for i := 1; i < len(s); i++ {
		if stack.Len()!=0 && s[i] == stack.Back().Value {
			stack.Remove(stack.Back())
		} else {
			stack.PushBack(s[i])
		}
	}
	var result []byte
	for e := stack.Front(); e!= nil; e = e.Next() {
		result = append(result,e.Value.(byte))
	}
	return  string(result)
}
~~~

**用切片模拟栈**：

~~~go
func removeDuplicates(s string) string {
    var stack = []byte{}
    for i:=0;i<len(s);i++{
        if len(stack) > 0 && s[i] == stack[len(stack)-1]{
            stack = stack[:len(stack)-1]
        }else{
            stack = append(stack,s[i])
        }
    }
    return string(stack)
}
~~~

最好使用**切片**模拟，原因详见 [Go语言中栈的实现：Slice还是List？](https://xiaoli-lhx.github.io/go%E8%AF%AD%E8%A8%80/go%E5%9F%BA%E7%A1%80/go%E8%AF%AD%E8%A8%80%E6%A8%A1%E6%8B%9F%E6%A0%88/)

#### 20.[有效的括号](https://leetcode.cn/problems/valid-parentheses/description/)

**思路**：

创建一个空栈。

从左到右遍历字符串中的每一个字符：

- 如果字符是**左括号** (`(`, `{`, `[`)，就把它**推入栈**中。 .
- 如果字符是**右括号** (`)`, `}`, `]`)：
  - 检查栈是否为空。如果为空，说明这个右括号没有对应的左括号，字符串**无效**。
  - 如果不为空，就从栈顶**弹出一个元素**。检查这个弹出的左括号是否和当前的右括号匹配。如果不匹配，字符串**无效**。

遍历完整个字符串后，检查栈是否为空。如果不为空，说明有剩下的左括号，字符串**无效**。

如果以上所有检查都通过了，那么字符串就是**有效**的。

~~~go
func isValid(s string) bool {
	paris := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	list := list.New()
	for _, ch := range s {
		if ch == '(' || ch == '{' || ch == '[' {
			// push
			list.PushBack(ch)
		}
		if ch == ')' || ch == '}' || ch == ']' {
			if list.Len() == 0 {
				return false
			} else {
				// pop
				old := list.Back()
				list.Remove(old)
				if old.Value != paris[ch] {
					return false
				}
			}
		}
	}
	if list.Len() != 0 {
		return false
	}
	return true
}
~~~



#### 225.[用队列实现栈](https://leetcode.cn/problems/implement-stack-using-queues/description/)

方法一：Push简单，Pop复杂 (懒惰入栈)

- **思路**:
  - `Push` 操作时，新元素直接加入队尾。
  - `Pop` 操作时，为了拿到“最后”一个入队的元素，需要把队列中排在它前面的所有元素，都从队头取出，再重新放回队尾。这样操作后，原先的“最后一个”元素就跑到了队头，可以轻松取出。
- **操作**:
  - **入栈 (Push)**: 直接在队尾添加元素。
  - **出栈 (Pop)**: 将 `n-1` 个队头元素依次挪到队尾，然后取出此时的队头元素。

~~~go
package main

import "container/list"

type MyStack struct {
	queue *list.List
}

func Constructor() MyStack {
	return MyStack{
		queue: list.New(),
	}
}

func (this *MyStack) Push(x int) {
	this.queue.PushBack(x)
}

func (this *MyStack) Pop() int {
	n := this.queue.Len()
	for i := 0; i < n-1; i++ {
		n := this.queue.Front().Value.(int)
		this.queue.Remove(this.queue.Front())
		this.Push(n)
	}
	x := this.queue.Front().Value.(int)
	this.queue.Remove(this.queue.Front())
	return x
}

func (this *MyStack) Top() int {
	x := this.Pop()
	this.Push(x)
	return x
}

func (this *MyStack) Empty() bool {
	return this.queue.Len() == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

~~~

---

方法二：Push复杂，Pop简单 (积极入栈)

- **思路**:
  - 每次 `Push` 一个新元素后，不让它“默默”待在队尾，而是立刻通过一系列操作，把它移动到队头。
  - 这样可以保证队列的头部永远是“最新”的元素，也就是栈顶元素。
- **操作**:
  - **入栈 (Push)**: 新元素先正常加入队尾，然后立刻将它移动到队头。
  - **出栈 (Pop)**: 直接取出队头元素即可。
  - **查看栈顶 (Top)**: 直接查看队头元素即可。

~~~go
type MyStack struct {
    queue *list.List
}


func Constructor() MyStack {
    return MyStack{
        queue: list.New(),
    }
}

func (this *MyStack) Push(x int)  {
    i:=this.queue.PushBack(x)
    this.queue.MoveToFront(i)
}


func (this *MyStack) Pop() int {
    front:=this.queue.Front()
    this.queue.Remove(front)
    return front.Value.(int)
}


func (this *MyStack) Top() int {
    return this.queue.Front().Value.(int)
}


func (this *MyStack) Empty() bool {
    return this.queue.Len()==0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
~~~

#### 232. [用栈实现队列](https://leetcode.cn/problems/implement-queue-using-stacks/description/)

**思路**:利用两个栈，一个作为新元素的入口缓冲 (`inStack`)，另一个作为老元素的出口 (`outStack`)。通过在 `outStack` 为空时，将 `inStack` 的元素全部转移到 `outStack` 的方式，完成一次顺序的逆转，从而用 LIFO 的栈模拟出了 FIFO 的队列行为。

~~~go
package main

import "fmt"

type MyQueue struct {
	inStack  []int // 输入栈：负责接收所有新加入的元素
	outStack []int // 输出栈：负责提供队头元素用于 Pop 和 Peek
}

// Constructor 初始化队列
func Constructor() MyQueue {
	return MyQueue{
		inStack:  make([]int, 0),
		outStack: make([]int, 0),
	}
}

// Push 将元素 x 推到队列的末尾
func (this *MyQueue) Push(x int) {
	// 直接将元素压入输入栈
	this.inStack = append(this.inStack, x)
}

// transferIfNeeded 是一个辅助函数。
// 当输出栈为空时，它会将输入栈的所有元素转移到输出栈。
func (this *MyQueue) transferIfNeeded() {
	// 只有当输出栈为空时，才需要进行转移操作
	if len(this.outStack) == 0 {
		// 循环直到输入栈为空
		for len(this.inStack) > 0 {
			// 1. 从输入栈的栈顶弹出一个元素
			val := this.inStack[len(this.inStack)-1]
			this.inStack = this.inStack[:len(this.inStack)-1]

			// 2. 将该元素压入输出栈
			this.outStack = append(this.outStack, val)
		}
	}
}

// Pop 从队列的开头移除并返回元素
func (this *MyQueue) Pop() int {
	// 在操作前，确保输出栈有正确的队头元素
	this.transferIfNeeded()

	// 如果两个栈都为空，这里会 panic，实际应用中可以返回 error
	// 但根据题目通常的假设，不会对空队列调用 Pop

	// 1. 从输出栈的栈顶获取队头元素
	val := this.outStack[len(this.outStack)-1]
	// 2. 将其从输出栈中移除
	this.outStack = this.outStack[:len(this.outStack)-1]

	return val
}

// Peek 返回队列的第一个元素（不移除）
func (this *MyQueue) Peek() int {
	// 同样，在操作前，确保输出栈有正确的队头元素
	this.transferIfNeeded()

	// 返回输出栈的栈顶元素，即为队头
	return this.outStack[len(this.outStack)-1]
}

// Empty 如果队列为空，则返回 true
func (this *MyQueue) Empty() bool {
	// 当且仅当两个栈都为空时，队列才为空
	return len(this.inStack) == 0 && len(this.outStack) == 0
}

func main() {
	// 你可以在这里添加测试代码来验证实现
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	fmt.Println("Peek:", queue.Peek()) // 输出: Peek: 1
	fmt.Println("Pop:", queue.Pop())   // 输出: Pop: 1
	fmt.Println("Empty:", queue.Empty()) // 输出: Empty: false
	fmt.Println("Pop:", queue.Pop())   // 输出: Pop: 2
	fmt.Println("Empty:", queue.Empty()) // 输出: Empty: true
}

~~~

### 二叉树(Binary tree)

#### 1. 二叉树的递归遍历

> 1. 确定递归函数的参数和返回值
> 2. 确定递归函数的终止条件
> 3. 确定单层递归的逻辑

##### 144.[二叉树的前序遍历](https://leetcode.cn/problems/binary-tree-preorder-traversal/description/)

~~~go
func preorderTraversal(root *TreeNode) []int {
	var res []int
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		// 终止条件
		if node == nil {
			return
		}
		// 先序遍历
		res = append(res, node.Val)
		// 左子树
		traversal(node.Left)
		// 右子树
		traversal(node.Right)
	}
	traversal(root)
	return res
}
~~~

##### 145.[二叉树的后序遍历](https://leetcode.cn/problems/binary-tree-postorder-traversal/)

~~~go
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var res []int
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		// 终止条件
		if node == nil {
			return
		}
		// 后续遍历
		traversal(node.Left)
		traversal(node.Right)
		// 记录结果
		res = append(res, node.Val)
	}
	traversal(root)
	return res
}
~~~

##### 94.[二叉树的中序遍历](https://leetcode.cn/problems/binary-tree-inorder-traversal/)

~~~go
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		// 终止条件
		if node == nil {
			return
		}
		// 中序遍历
		traversal(node.Left)
		res = append(res, node.Val)
		traversal(node.Right)
	}
	traversal(root)
	return res
}
~~~

#### 2. 二叉树的迭代遍历

##### 144.[二叉树的前序遍历](https://leetcode.cn/problems/binary-tree-preorder-traversal/description/)

**思路：**

我们把这个过程想象成一个“**待办事项**”列表（这个列表就是我们的栈）。

1. **开始工作**：我们接到的第一个任务就是处理根节点。所以，我们先把**根节点**放入“待办事项”列表（入栈）。
2. **处理当前任务**：我们从列表里拿出任务来做。
   - 现在唯一的任务是根节点，我们把它拿出来（出栈）。
   - 按照前序遍历“**根** -> 左 -> 右”的顺序，我们**立刻**处理它，把它的值记录到结果里。“根”的部分就完成了。
3. **添加新任务**：处理完根节点后，它告诉我们接下来还有两个新任务：处理它的左子树和右子树。
   - 我们的“待办事项”列表是后进先出的。为了保证我们**先**做“处理左子树”这个任务，我们必须把它**后**放入列表。
   - 所以，我们先把**右子节点**放入待办列表（入栈），再把**左子节点**放入待办列表（入栈）。

你看，我们并不是一次性地把“右左中”都放进去。而是“**先放中 -> 取出中并处理 -> 再放右 -> 最后放左**”。

这样一来，下一个从栈顶被取出来的，自然就是我们后放进去的左子节点了，完美实现了“根 -> 左 -> 右”的顺序。

~~~go
func preorderTraversal(root *TreeNode) []int {
	var res []int
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}
		res = append(res, node.Val)
		stack = append(stack, node.Right)
		stack = append(stack, node.Left)
	}
	return res
}
~~~

##### 145.[二叉树的后序遍历](https://leetcode.cn/problems/binary-tree-postorder-traversal/)

**思路：**

先序遍历的入栈是入根，出根，入右，入左，出左，出右，从而保证了出栈顺序是根左右

后序遍历出栈顺序应该是左右根，为了保证出栈顺序的正确，我们调整一下先序遍历的入栈顺序，变成入根，出根，入左，入右，出右，出左，从而保证了出栈顺序是根右左，再将结果反转，即为左右根

~~~go
func postorderTraversal(root *TreeNode) []int {
	var res []int
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}
		res = append(res, node.Val)
		stack = append(stack, node.Left)
		stack = append(stack, node.Right)
	}
	reverse(res)
	return res
}
func reverse(arr []int) {
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}

~~~

##### 94.[二叉树的中序遍历](https://leetcode.cn/problems/binary-tree-inorder-traversal/)

**思路：**

中序遍历迭代法的三个核心步骤都想出来了：

1. **一路向左**：从一个节点出发，不断访问它的左子节点，并把路径上所有节点都推入栈中，直到左边走到底。
2. **处理节点**：当左边走到底时，从栈里弹出一个节点并记录它的值。
3. **转向右侧**：然后，把这个弹出的节点作为“根”，转向它的右子树，对右子树**重复以上所有过程**。

这个循环会一直进行，直到栈和当前节点都为空，说明整棵树都处理完了。

~~~go
// 迭代法
func inorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		node = node.Right
	}
	return res
}
~~~



## 三、设计类问题

### `LRU` 缓存

#### [146. `LRU` 缓存](https://leetcode.cn/problems/lru-cache/description/)

**思路:** 使用“哈希表 + 双向链表”。哈希表 `map[key]*list.Element` 实现 O(1) 的查找，双向链表维护节点的访问顺序。每次 `get` 或 `put` 操作后，将被访问的节点移动到链表头部。当容量满时，淘汰链表尾部的节点。

```
import "container/list"

type LRUCache struct {
    capacity int
    cache    map[int]*list.Element
    list     *list.List
}
type entry struct {
    key   int
    value int
}
func Constructor(capacity int) LRUCache {
    return LRUCache{
        capacity: capacity,
        cache:    make(map[int]*list.Element),
        list:     list.New(),
    }
}
func (this *LRUCache) Get(key int) int {
    if elem, ok := this.cache[key]; ok {
        this.list.MoveToFront(elem)
        return elem.Value.(*entry).value
    }
    return -1
}
func (this *LRUCache) Put(key, value int) {
    if elem, ok := this.cache[key]; ok {
        elem.Value.(*entry).value = value
        this.list.MoveToFront(elem)
        return
    }
    if this.list.Len() == this.capacity {
        last := this.list.Back()
        if last != nil {
            delete(this.cache, last.Value.(*entry).key)
            this.list.Remove(last)
        }
    }
    elem := this.list.PushFront(&entry{key, value})
    this.cache[key] = elem
}
```

### `LFU `缓存

#### [460. `LFU` 缓存](https://leetcode.cn/problems/lfu-cache/description/)

**思路:** 使用“两个哈希表 + 双向链表”。

1. `objectMap (map[key]*list.Element)`: 存储 key 到节点的映射。

2. `freqMap (map[freq]*list.List)`: 存储频率到该频率下所有节点的双向链表的映射，每个链表内部按 LRU 规则排序。

3. `minFreq`: 记录当前缓存中最低的访问频率。

   淘汰时，从 `minFreq `对应的链表中淘汰最久未使用的节点（链表尾部）。

~~~go
type entry struct {
	key   int
	value int
	freq  int
	elem  *list.Element // 在频率节点的 items 链表里的位置
}

type freqNode struct {
	freq  int
	items map[int]*entry
	list  *list.List // 维护 LRU
}

type LFUCache struct {
	capacity int
	minFreq  int
	cache    map[int]*entry
	freqMap  map[int]*list.Element // freq -> freqList 节点
	freqList *list.List
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		cache:    make(map[int]*entry),
		freqMap:  make(map[int]*list.Element),
		freqList: list.New(),
	}
}

func (c *LFUCache) Get(key int) int {
	if en, ok := c.cache[key]; ok {
		c.incrFreq(en)
		return en.value
	}
	return -1
}

func (c *LFUCache) Put(key, value int) {
	if c.capacity <= 0 {
		return
	}
	if en, ok := c.cache[key]; ok {
		en.value = value
		c.incrFreq(en)
		return
	}
	if len(c.cache) >= c.capacity {
		c.evict()
	}
	c.addNewEntry(key, value)
}

func (c *LFUCache) incrFreq(e *entry) {
	oldFreq := e.freq
	oldFreqElem := c.freqMap[oldFreq]
	oldFreqNode := oldFreqElem.Value.(*freqNode)

	// 从旧频率节点删除
	delete(oldFreqNode.items, e.key)
	oldFreqNode.list.Remove(e.elem)

	// 如果旧频率节点空了，删除 freqList 里的节点
	if len(oldFreqNode.items) == 0 {
		c.freqList.Remove(oldFreqElem)
		delete(c.freqMap, oldFreq)
		// 如果是最小频率，minFreq++
		if oldFreq == c.minFreq {
			c.minFreq++
		}
	}

	// 更新频率
	e.freq++
	newFreq := e.freq

	// 将节点放入新频率节点
	if newFreqElem, ok := c.freqMap[newFreq]; ok {
		newFreqNode := newFreqElem.Value.(*freqNode)
		e.elem = newFreqNode.list.PushBack(e.key)
		newFreqNode.items[e.key] = e
	} else {
		newFreqNode := &freqNode{
			freq:  newFreq,
			items: make(map[int]*entry),
			list:  list.New(),
		}
		e.elem = newFreqNode.list.PushBack(e.key)
		newFreqNode.items[e.key] = e

		// 找到插入位置，保持 freqList 递增
		var insertBefore *list.Element
		for elem := c.freqList.Front(); elem != nil; elem = elem.Next() {
			if elem.Value.(*freqNode).freq > newFreq {
				insertBefore = elem
				break
			}
		}
		var newElem *list.Element
		if insertBefore != nil {
			newElem = c.freqList.InsertBefore(newFreqNode, insertBefore)
		} else {
			newElem = c.freqList.PushBack(newFreqNode)
		}
		c.freqMap[newFreq] = newElem
	}
}

func (c *LFUCache) evict() {
	minFreqElem := c.freqMap[c.minFreq]
	minFreqNode := minFreqElem.Value.(*freqNode)

	oldestKeyElem := minFreqNode.list.Front()
	oldestKey := oldestKeyElem.Value.(int)

	delete(minFreqNode.items, oldestKey)
	minFreqNode.list.Remove(oldestKeyElem)
	delete(c.cache, oldestKey)

	if len(minFreqNode.items) == 0 {
		c.freqList.Remove(minFreqElem)
		delete(c.freqMap, c.minFreq)
		// 下一个 put 会重置 minFreq 为 1
	}
}

func (c *LFUCache) addNewEntry(key, value int) {
	newFreq := 1
	en := &entry{
		key:   key,
		value: value,
		freq:  newFreq,
	}
	if freqElem, ok := c.freqMap[newFreq]; ok {
		freqNode := freqElem.Value.(*freqNode)
		en.elem = freqNode.list.PushBack(key)
		freqNode.items[key] = en
	} else {
		freqNode := &freqNode{
			freq:  newFreq,
			items: make(map[int]*entry),
			list:  list.New(),
		}
		en.elem = freqNode.list.PushBack(key)
		freqNode.items[key] = en
		freqElem := c.freqList.PushFront(freqNode)
		c.freqMap[newFreq] = freqElem
	}
	c.cache[key] = en
	c.minFreq = newFreq
}
~~~

~~~go
package main

import "container/list"

// Object 一个key-value元素
type Object struct {
	key, val int
	freq     int // 频率
}

// LFUCache 一个LFU缓存
type LFUCache struct {
	capacity  int // 最大容量
	len       int // 长度
	minFreq   int // 最小频率
	objectMap map[int]*list.Element
	// key: nodeMap中所有元素出现的可能频次
	// val: NodeList频次相同头部的元素，操作时间离现在最近
	freqMap map[int]*list.List
}

// Constructor 构造函数
func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity:  capacity,
		len:       0,
		minFreq:   0,
		objectMap: make(map[int]*list.Element),
		freqMap:   make(map[int]*list.List),
	}
}

func (c *LFUCache) Get(key int) int {
	// 判断是否存在
	if e, ok := c.objectMap[key]; ok {
		ob := e.Value.(*Object)
		// 存在，更新频率
		c.updateFreq(e)

		return ob.val
	}
	return -1
}

func (c *LFUCache) Put(key int, value int) {
	if e, ok := c.objectMap[key]; ok {
		ob := e.Value.(*Object)
		ob.val = value
		// 存在，更新频率
		c.updateFreq(e)

	} else {
		// 不存在，新建
		ob := &Object{key: key, val: value, freq: 1}
		if c.len == c.capacity {
			// 缓存满了，需要淘汰
			c.evict()
			c.len--
		}
		// 添加
		c.insertMap(ob)
		c.minFreq = 1
		c.len++
	}
}

func (c *LFUCache) updateFreq(e *list.Element) {
	ob := e.Value.(*Object)
	// 从oldList中移除
	oldList := c.freqMap[ob.freq]
	oldList.Remove(e)
	// 特殊处理
	if ob.freq == c.minFreq && oldList.Len() == 0 {
		c.minFreq++
	}
	// 添加到新newList中
	ob.freq++
	c.insertMap(ob)
}

func (c *LFUCache) evict() {
	l := c.freqMap[c.minFreq]
	e := l.Back()
	ob := e.Value.(*Object)
	// 移除
	l.Remove(e)
	delete(c.objectMap, ob.key)
}

func (c *LFUCache) insertMap(ob *Object) {
	newList, ok := c.freqMap[ob.freq]
	if !ok {
		newList = list.New()
		c.freqMap[ob.freq] = newList
	}
	newElement := newList.PushFront(ob)
	c.objectMap[ob.key] = newElement
}
~~~


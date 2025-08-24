# `LeetCode` 刷题笔记 (重构版)

这是一份根据练习记录重构的刷题笔记，主要分为**核心算法思想**、**基础数据结构**和**设计类问题**三大模块，旨在提供一个更清晰、更系统化的复习框架。

## 目录

[TOC]

## 一、核心算法思想

### 1. 双指针 (Two Pointers)

双指针是一种通过维护两个指针在序列中同向或相向移动，来降低时间复杂度的技巧。

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

### 2. 滑动窗口 (Sliding Window)

滑动窗口是双指针的一种特例，用于解决子数组/子串问题。两个指针维护一个“窗口”，根据条件移动右指针扩大窗口，移动左指针缩小窗口。

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

### 3. 二分查找 (Binary Search)

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

### 4. 贪心算法 (Greedy Algorithm)

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

### 1. 数组 (Array)

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

### 2. 链表 (Linked List)

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

### 3. 字符串 (String)

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

### 4. 哈希表 (Hash Table)

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

## 三、设计类问题

### 1. `LRU` 缓存

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

### 2. `LFU `缓存

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


package hard

import (
	"container/heap"
	"fmt"
	"sort"
)

/*
	42.接雨水

	给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

	https://leetcode.cn/problems/trapping-rain-water/
*/

// []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
func Trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	leftMax := make([]int, n)
	rightMax := make([]int, n)
	leftMax[0] = height[0]
	rightMax[n-1] = height[n-1]
	// 从左边找最高的边，并一直维护
	// []int len: 12, cap: 12, [0,1,1,2,2,2,2,3,3,3,3,3]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	// 从右边开始找最高的边，并一直维护
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}
	ans := 0

	// 从左右两边找到最高的，减去墙体
	// []int len: 12, cap: 12, [3,3,3,3,3,3,3,3,2,2,2,1]
	for i := 0; i < n; i++ {
		ans += min(leftMax[i], rightMax[i]) - height[i]
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
   239. 滑动窗口最大值
   给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
   返回 滑动窗口中的最大值

   示例 1：

   输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
   输出：[3,3,5,5,6,7]
   解释：
   滑动窗口的位置                最大值
   ---------------               -----
   [1  3  -1] -3  5  3  6  7       3
    1 [3  -1  -3] 5  3  6  7       3
    1  3 [-1  -3  5] 3  6  7       5
    1  3  -1 [-3  5  3] 6  7       5
    1  3  -1  -3 [5  3  6] 7       6
    1  3  -1  -3  5 [3  6  7]      7
*/

var arr []int

type hp struct{ sort.IntSlice }

func (h *hp) Less(i, j int) bool   { return arr[h.IntSlice[i]] > arr[h.IntSlice[j]] }
func (h *hp) Push(val interface{}) { h.IntSlice = append(h.IntSlice, val.(int)) }
func (h *hp) Pop() interface{} {
	arr := h.IntSlice
	v := arr[len(arr)-1]
	h.IntSlice = h.IntSlice[:len(arr)-1]
	return v
}

func MaxSlidingWindow(nums []int, k int) []int {
	fmt.Println(nums)
	arr = nums
	q := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}
	heap.Init(q)

	n := len(nums)
	ans := make([]int, 1, n - k + 1)
	ans[0] = nums[q.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(q, i)
		for q.IntSlice[0] <= i - k {
			heap.Pop(q)
		}
		ans = append(ans, nums[q.IntSlice[0]])
	}

	return ans
}

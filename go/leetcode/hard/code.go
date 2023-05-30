package hard

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

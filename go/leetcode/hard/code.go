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
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(q, i)
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}
		ans = append(ans, nums[q.IntSlice[0]])
	}

	return ans
}

/*
	301. 删除无效的括号
	困难
	相关标签
	相关企业
	提示
	给你一个由若干括号和字母组成的字符串 s ，删除最小数量的无效括号，使得输入的字符串有效。

	返回所有可能的结果。答案可以按 任意顺序 返回。



	示例 1：

	输入：s = "()())()"
	输出：["(())()","()()()"]
	示例 2：

	输入：s = "(a)())()"
	输出：["(a())()","(a)()()"]
	示例 3：

	输入：s = ")("
	输出：[""]
*/

func isValid301(s string) bool {
	count := 0
	for _, c := range s {
		if c == '(' {
			count++
		} else if c == ')' {
			count--
			if count < 0 {
				return false
			}
		}
	}

	return count == 0
}

func dfs301(ans *[]string, s string, start, lCount, rCount, l, r int) {
	if l == 0 && r == 0 {
		if isValid301(s) {
			*ans = append(*ans, s)
		}
		return
	}

	for i := start; i < len(s); i++ {
		if s[i] == '(' {
			lCount++
		}

		if s[i] == ')' {
			rCount++
		}

		if i > start && s[i] == s[i-1] {
			continue
		}

		if s[i] == '(' && l > 0 {
			dfs301(ans, s[:i]+s[i+1:], i, lCount-1, rCount, l-1, r)
		}
		if s[i] == ')' && r > 0 {
			dfs301(ans, s[:i]+s[i+1:], i, lCount, rCount-1, l, r-1)
		}

		if rCount > lCount {
			break
		}
	}
}

func RemoveInvalidParentheses(s string) []string {
	lRemove, rRemove := 0, 0
	for _, c := range s {
		if c == '(' {
			lRemove++
		} else if c == ')' && lRemove == 0 {
			rRemove++
		} else if c == ')' && lRemove > 0 {
			lRemove--
		}
	}

	ans := make([]string, 0)
	dfs301(&ans, s, 0, 0, 0, lRemove, rRemove)

	return ans
}

/*
    https://leetcode.cn/problems/burst-balloons/description/?envType=problem-list-v2&envId=2cktkvj
	代码
	312. 戳气球
	已解答
	困难
	相关标签
	相关企业
	有 n 个气球，编号为0 到 n - 1，每个气球上都标有一个数字，这些数字存在数组 nums 中。

	现在要求你戳破所有的气球。戳破第 i 个气球，你可以获得 nums[i - 1] * nums[i] * nums[i + 1] 枚硬币。 这里的 i - 1 和 i + 1 代表和 i 相邻的两个气球的序号。如果 i - 1或 i + 1 超出了数组的边界，那么就当它是一个数字为 1 的气球。

	求所能获得硬币的最大数量。



	示例 1：
	输入：nums = [3,1,5,8]
	输出：167
	解释：
	nums = [3,1,5,8] --> [3,5,8] --> [3,8] --> [8] --> []
	coins =  3*1*5    +   3*5*8   +  1*3*8  + 1*8*1 = 167
	示例 2：

	输入：nums = [1,5]
	输出：10


	提示：

	n == nums.length
	1 <= n <= 300
	0 <= nums[i] <= 100
*/

func MaxCoins(nums []int) int {
	nums = append(nums, 1)
	nums = append([]int{1}, nums...)

	n := len(nums)
	dp := make([][]int, n)
	sli := make([]int, n*n)
	for i := 0; i < n; i++ {
		dp[i] = sli[:n]
		sli = sli[n:]
	}

	// 这里的转移逻辑为，假设最后一个不打破的气球为k
	// 那么结果为 coins[i+1,k-1] + coins[i] * coins[k] * coins[j] + coins[k+1,j-1]
	// 将i+1到k-1的气球全部打破，积分记为c[i+1,k-1]
	// 将k+1到j-1的气球全部打破，积分记为c[k+1,j-1]
	// 那么最后剩的气球为c[i] * c[k] * c[j]
	// 所以转移方程为：
	// c[i][j] = max(c[i][j], c[i+1,k-1] + c[i] * c[k] * coins[j] + c[k+1,j-1])
	/*
		    nums = [1,3,1,5,8,1]
			dp[0][2] = 3
			dp[1][3] = 15
			dp[2][4] = 40
			dp[3][5] = 40
			dp[0][4] = 30
	*/
	// 长度从3开始，只要nums长度大于1，这里满足循环条件，本题nums长度的区间范围是大于等于1的
	for l := 3; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			j := i + l - 1
			// c[i+1] -- c[j-1]
			for k := i + 1; k <= j-1; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+nums[i]*nums[k]*nums[j]+dp[k][j])
			}
		}
	}

	return dp[0][n-1]
}

/*
	https://leetcode.cn/problems/first-missing-positive/description/?envType=study-plan-v2&envId=top-100-liked
	41. 缺失的第一个正数
	给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。

	请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。


	示例 1：

	输入：nums = [1,2,0]
	输出：3
	解释：范围 [1,2] 中的数字都在数组中。
	示例 2：

	输入：nums = [3,4,-1,1]
	输出：2
	解释：1 在数组中，但 2 没有。
	示例 3：

	输入：nums = [7,8,9,11,12]
	输出：1
	解释：最小的正数 1 没有出现。

*/

func FirstMissingPositive(nums []int) int {
	l := len(nums)
	for i := range nums {
		for nums[i] > 0 && nums[i] <= l && nums[i] != i+1 {
			t := nums[i] - 1
			if nums[t] == t+1 {
				break
			}
			nums[i], nums[t] = nums[t], nums[i]
		}
	}

	for i := 0; i <= len(nums); i++ {
		v := nums[i]
		if v != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

// func SolveNQueens(n int) [][]string {
// 	ans := make([][]string, 0)
// 	board := make([][]byte, 0)
// 	row := strings.Repeat(".", n)
// 	for i := 0; i < n; i++ {
// 		board = append(board, []byte(row))
// 	}
// 	var backtrack func(int)
// 	backtrack = func(i int) {
// 		if i == n {
// 			ans = append(ans, slices.Clone(board))
// 			return
// 		}

// 		for j := 0; j < n; j++ {
// 			if !isValid(board, i, j) {
// 				continue
// 			}

// 			board[i][j] = 'Q'
// 			backtrack(i + 1)
// 			board[i][j] = '.'
// 		}
// 	}
// 	backtrack(0)
// 	return ans
// }

// func isValid(board [][]byte, i, j int) bool {
// 	n := len(board)
// 	for row := 0; row < n; row++ {
// 		if board[row][j] == 'Q' {
// 			return false
// 		}
// 	}

// 	row := i - 1
// 	col := j - 1
// 	for row >= 0 && col >= 0 {
// 		if board[row][col] == 'Q' {
// 			return false
// 		}
// 		row--
// 		col--
// 	}

// 	row = i - 1
// 	col = j + 1
// 	for row >= 0 && col < n {
// 		if board[row][col] == 'Q' {
// 			return false
// 		}
// 		row--
// 		col++
// 	}
// 	return true
// }

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m == 0 && n == 0 {
		return 0.0
	}
	oddNum := false
	if (m+n)%2 == 1 {
		oddNum = true
	}

	if m == 0 || n == 0 {
		nums := nums1
		if m == 0 {
			nums = nums2
		}
		if oddNum {
			return float64(nums[len(nums)/2])
		} else {
			return float64(nums[len(nums)/2]+nums[(len(nums)-1)/2]) / 2
		}
	}

	k1, k2 := (m+n)/2, (m+n)/2
	if !oddNum {
		k2 = k1 + 1
	}

	p1, p2 := 0, 0
	for (p1+p2) < k2 && p1 < m-1 && p2 < n-1 {
		if nums1[p1] > nums2[p2] {
			p2++
		} else {
			p1++
		}
	}

	var ans float64
	if !oddNum {
		ans = float64((nums1[p1] + nums2[p2])) / 2
	} else {
		if nums1[p1] > nums2[p2] {
			ans = float64(nums1[p1])
		} else {
			ans = float64(nums2[p2])
		}
	}
	return ans
}

/*
32. 最长有效括号
https://leetcode.cn/problems/longest-valid-parentheses/description/?envType=study-plan-v2&envId=top-100-liked
*/
func LongestValidParentheses(s string) int {
	ans := 0
	stack := make([]int, 0)
	stack = append(stack, -1)
	for i, c := range s {
		if c == '(' {
			stack = append(stack, i)
		} else {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			if len(stack) <= 0 {
				stack = append(stack, i)
			} else {
				ans = max(ans, i-stack[len(stack)-1])
			}
		}
	}
	return ans
}

package mid

import (
	"fmt"
	"math"
	"strings"
)

/*

	https://leetcode.cn/problems/word-break/
	139
	给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s 。

	注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。

	来源：力扣（LeetCode）
	链接：https://leetcode.cn/problems/word-break
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

func WordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)

	for _, word := range wordDict {
		wordDictSet[word] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			// s[j:i]表示单词是否存在
			// 存在且标记为
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

/*
	https://leetcode.cn/problems/course-schedule/
	207 课程表

	你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。

	在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。

	例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
	请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。

	来源：力扣（LeetCode）
	链接：https://leetcode.cn/problems/course-schedule
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 图的深度优先搜索
func CanFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges   = make([][]int, numCourses)
		visited = make([]int, numCourses)
		result  []int
		valid   = true
		dfs     func(u int)
	)

	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[u] = 2
		result = append(result, u)
	}

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	return valid
}

/*
	208
	Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼写检查。

	请你实现 Trie 类：

	Trie() 初始化前缀树对象。
	void insert(String word) 向前缀树中插入字符串 word 。
	boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
	boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。

	来源：力扣（LeetCode）
	链接：https://leetcode.cn/problems/implement-trie-prefix-tree
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, ch := range word {
		// 定位到数组中的位置
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (this *Trie) SearchPrefix(word string) *Trie {
	node := this
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

func (this *Trie) Search(word string) bool {
	node := this.SearchPrefix(word)
	return node != nil && node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.SearchPrefix(prefix) != nil
}

/*
  给你 k 种面值的硬币，面值分别为 c1, c2 ... ck，每种硬币的
  数量无限，再给一个总金额 amount，问你最少需要几枚硬币凑出这个金额，
 如果不可能凑出，算法返回 -1
*/

var dict = map[int]int{}

func CoinsChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	if amount < 0 {
		return -1
	}

	res, ok := dict[amount]
	if ok {
		return res
	}

	res = math.MaxInt
	for _, coin := range coins {
		subProblem := CoinsChange(coins, amount-coin)
		if subProblem == -1 {
			continue
		}
		res = Min(res, 1+subProblem)
	}

	if res != math.MaxInt && res != -1 {
		dict[amount] = res
	}

	return res
}

func Min(l, r int) int {
	if l < r {
		return l
	}
	return r
}

func ProductExceptSelf(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)

	fmt.Println(nums)

	ans[0] = 1
	for i := 1; i < n; i++ {
		ans[i] = nums[i-1] * ans[i-1]
		fmt.Println(ans)
	}

	R := 1
	for i := n - 1; i >= 0; i-- {
		ans[i] = ans[i] * R
		R *= nums[i]
		fmt.Println(ans)
		fmt.Println(R)
	}

	return ans
}

func FindDisappearedNumbers(nums []int) []int {
	n := len(nums)
	ans := make([]int, 0)
	for _, num := range nums {
		num := (num - 1) % n
		nums[num] += n
	}

	for i, num := range nums {
		if num <= n {
			ans = append(ans, i+1)
		}
	}

	return ans
}

/*

	https://leetcode.cn/problems/perfect-squares/description/?envType=featured-list&envId=2cktkvj?envType=featured-list&envId=2cktkvj
	279
	给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
	完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。

*/

/*
    每当遇到一个平方数的时候：i = j * j，会将当前数的结果归一
	遍历前数，如果有子平方的和，则会累计加1
*/

func NumSquares(n int) int {
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minn := math.MaxInt
		for j := 1; j*j <= i; j++ {
			minn = Min(minn, f[i-j*j]+1)
		}
		f[i] = minn
	}
	return f[n]
}

// 394. 字符串解码
// https://leetcode.cn/problems/decode-string/description/?envType=problem-list-v2&envId=2cktkvj
// 中等
// 给定一个经过编码的字符串，返回它解码后的字符串。

// 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

// 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

// 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
func DecodeString(s string) string {
	type pair struct {
		num int
		sb  *strings.Builder
	}

	var sb = &strings.Builder{}
	var stack []pair
	var num int

	for _, ch := range s {
		switch {
		case '0' <= ch && ch <= '9':
			num = num*10 + int(ch-'0')
		// 支持嵌套结构
		case ch == '[':
			stack = append(stack, pair{num: num, sb: sb})
			num = 0
			sb = &strings.Builder{}
		case ch == ']':
			p := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 把当前的加入到前面的 sb 上面
			p.sb.Grow(p.num * sb.Len())
			// for j := 0; j < p.num; j++ {
			// 	p.sb.WriteString(sb.String())
			// }
			decodeStr := strings.Repeat(sb.String(), p.num)
			p.sb.WriteString(decodeStr)
			// sb 换成前面的
			sb = p.sb
		default:
			sb.WriteRune(ch)
		}
	}
	return sb.String()
}

// func GroupAnagrams(strs []string) [][]string {
// 	m := make(map[string][]string, 0)
// 	for _, s := range strs {
// 		t := []byte(s)
// 		slices.Sort(t)
// 		sortS := t
// 		m[string(sortS)] = append(m[string(sortS)], s)
// 	}

// 	ans := make([][]string, 0)
// 	for _, l := range m {
// 		ans = append(ans, l)
// 	}

// 	return ans
// }

/*
	    https://leetcode.cn/problems/3sum/description/?envType=study-plan-v2&envId=top-100-liked
		15. 三数之和
		已解答
		中等
		相关标签
		相关企业
		提示
		给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。

		注意：答案中不可以包含重复的三元组。


		示例 1：

		输入：nums = [-1,0,1,2,-1,-4]
		输出：[[-1,-1,2],[-1,0,1]]
		解释：
		nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
		nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
		nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
		不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
		注意，输出的顺序和三元组的顺序并不重要。
		示例 2：

		输入：nums = [0,1,1]
		输出：[]
		解释：唯一可能的三元组和不为 0 。
		示例 3：

		输入：nums = [0,0,0]
		输出：[[0,0,0]]
		解释：唯一可能的三元组和为 0 。
*/
// func ThreeSum(nums []int) [][]int {
// 	slices.Sort(nums)

// 	ans := make([][]int, 0)
// 	n := len(nums)
// 	for i := 0; i < n; i++ {
// 		// 经过排序后，如果最小值都大于0，那么与之后的值加和不可能等于0
// 		if nums[i] > 0 {
// 			return ans
// 		}

// 		// 对第一个元素做去重逻辑
// 		// 去重前值，保证每一个元素在首次遇到时，可以被包含在计算范围内
// 		// 后续计算的两个元素，只需要保证是当前值的负值即可
// 		// 如果列表值为[-1,-1,-1,2]
// 		// nums[i] == nums[i+1]会导致上述类型的计算被去重
// 		if i > 0 && nums[i] == nums[i-1] {
// 			continue
// 		}

// 		l := i + 1
// 		r := n - 1
// 		for l < r {
// 			// 列表已经排序

// 			if nums[i]+nums[l]+nums[r] > 0 { // 数字向减小方向移动
// 				r--
// 			} else if nums[i]+nums[l]+nums[r] < 0 { // 数字向增大方向移动
// 				l++
// 			} else {
// 				ans = append(ans, []int{nums[i], nums[l], nums[r]})
// 				// 元素2去重
// 				for l < r && nums[r] == nums[r-1] {
// 					r--
// 				}
// 				// 元素3去重
// 				for l < r && nums[l] == nums[l+1] {
// 					l++
// 				}
// 				// 上面的去重逻辑不包含当前被计算过的值，去重只到边界，因此需要再挪动一次索引
// 				r--
// 				l++
// 			}
// 		}
// 	}

// 	return ans
// }

func LengthOfLongestSubstring(s string) int {
	ans := 0
	l := 0
	hash := make(map[rune]bool)
	for r, c := range s {
		for hash[c] {
			delete(hash, rune(s[l]))
			l++
		}
		hash[c] = true
		ans = max(ans, r-l+1)
	}
	return ans
}

/*
    https://leetcode.cn/problems/find-all-anagrams-in-a-string/description/?envType=study-plan-v2&envId=top-100-liked
	438. 找到字符串中所有字母异位词
  	给定两个字符串 s 和 p，找到 s 中所有 p 的
	异位词
	的子串，返回这些子串的起始索引。不考虑答案输出的顺序。



	示例 1:

	输入: s = "cbaebabacd", p = "abc"
	输出: [0,6]
	解释:
	起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
	起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
	示例 2:

	输入: s = "abab", p = "ab"
	输出: [0,1,2]
	解释:
	起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
	起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
	起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。


	提示:

	1 <= s.length, p.length <= 3 * 104
	s 和 p 仅包含小写字母
*/

func FindAnagrams(s string, p string) []int {
	ans := make([]int, 0)
	cnt := [128]int{}
	for _, c := range p {
		cnt[c]++
	}

	left := 0
	window := [128]int{}
	for right, c := range s {
		window[c]++
		// 窗口满了
		for window[c] > cnt[c] {
			window[s[left]]--
			left++
		}

		if right-left+1 == len(p) {
			ans = append(ans, left)
		}
	}

	return ans
}

/*
	https://leetcode.cn/problems/subarray-sum-equals-k/description/?envType=study-plan-v2&envId=top-100-liked
	560. 和为 K 的子数组
	已解答
	中等
	相关标签
	相关企业
	提示
	给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。

	子数组是数组中元素的连续非空序列。



	示例 1：

	输入：nums = [1,1,1], k = 2
	输出：2
	示例 2：

	输入：nums = [1,2,3], k = 3
	输出：2


	提示：

	1 <= nums.length <= 2 * 104
	-1000 <= nums[i] <= 1000
	-107 <= k <= 107
*/

func SubarraySum(nums []int, k int) int {
	ans := 0

	hash := map[int]int{0: 1}
	//  处理0 - 0 = 0
	sum := 0
	for _, num := range nums {
		sum += num
		// sum 为 0 到 i的前缀和
		// sum[x,y]为从x到y的前缀和
		// sum - sum[x,y] = k，即y+1到i部分的值为k
		// sum - k == sum[x,y]
		// 序列中可以连续出现0，每出现一次，前缀和不变，但子序列要增加
		ans += hash[sum-k]
		hash[sum]++
	}

	return ans
}

/*
	https://leetcode.cn/problems/maximum-subarray/description/?envType=study-plan-v2&envId=top-100-liked
	代码
	测试用例
	测试结果
	测试结果
	53. 最大子数组和
	已解答
	中等
	相关标签
	相关企业
	给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

	子数组
	是数组中的一个连续部分。



	示例 1：

	输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
	输出：6
	解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
	示例 2：

	输入：nums = [1]
	输出：1
	示例 3：

	输入：nums = [5,4,-1,7,8]
	输出：23


	提示：

	1 <= nums.length <= 105
	-104 <= nums[i] <= 104
*/

func MaxSubArray(nums []int) int {
	ans := math.MinInt

	sum := 0
	for _, n := range nums {
		sum += n
		if sum > ans {
			ans = sum
		}

		if sum < 0 {
			sum = 0
		}
	}
	return ans
}

/*
	https://leetcode.cn/problems/set-matrix-zeroes/description/?envType=study-plan-v2&envId=top-100-liked
	73. 矩阵置零
	提示
	给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。

	示例 1：

	输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
	输出：[[1,0,1],[0,0,0],[1,0,1]]
	示例 2：

	输入：matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
	输出：[[0,0,0,0],[0,4,5,0],[0,3,1,0]]

	提示：

	m == matrix.length
	n == matrix[0].length
	1 <= m, n <= 200
	-231 <= matrix[i][j] <= 231 - 1
*/

func SetZeroes(matrix [][]int) {
	// 为了不占用额外空间，在matrix上进行标记是否置零
	/*
		思路：
		将matrix的首行首列记录是否含有0
		然后从第二行和第二列开始统计是否含有0，将结果记录在首行首列中
		根据首行首列的记录进行置零
		根据统计首行首列的结果对首行首列进行操作
	*/

	firstRow := false
	firstCol := false

	// 记录行首
	for i := range matrix[0] {
		if matrix[0][i] == 0 {
			firstRow = true
			break
		}
	}

	// 记录列首
	for i := range matrix {
		if matrix[i][0] == 0 {
			firstCol = true
			break
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	rowLen := len(matrix)
	colLen := len(matrix[0])
	for i := 1; i < rowLen; i++ {
		for j := 1; j < colLen; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if firstRow {
		for i := range matrix[0] {
			matrix[0][i] = 0
		}
	}

	if firstCol {
		for i := range matrix {
			matrix[i][0] = 0
		}
	}
}

func SearchMatrix(matrix [][]int, target int) bool {
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1

	for left <= right && top <= bottom {
		if matrix[bottom][left] > target {
			bottom--
		} else if matrix[bottom][left] < target {
			left++
		} else {
			return true
		}
	}

	return false
}

func Partition(s string) [][]string {
	n := len(s)
	track := make([]string, 0)
	ans := make([][]string, 0)
	var dfs func(int, int)
	dfs = func(i, start int) {
		if i == n {
			ans = append(ans, append([]string(nil), track...))
			return
		}

		if i < n-1 {
			dfs(i+1, start)
		}

		if isPalindrome(s, start, i) {
			track = append(track, s[start:i+1])
			dfs(i+1, i+1)
			track = track[:len(track)-1]
		}
	}
	dfs(0, 0)
	return ans
}

func isPalindrome(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func MaxProfit(prices []int) int {
	n := len(prices)
	dp_i_0, dp_i_1 := 0, math.MinInt

	for i := 0; i < n; i++ {
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = max(dp_i_1, -prices[i])
	}
	return dp_i_0
}

var robMem []int

func Rob(nums []int) int {
	robMem = make([]int, len(nums)+2)
	for i := range robMem {
		robMem[i] = -1
	}
	return robDp(nums, 0)
}

func robDp(nums []int, start int) int {
	if len(nums) >= start {
		return 0
	}

	if robMem[start] != -1 {
		return robMem[start]
	}

	robMem[start] = max(robDp(nums, start+1), nums[start]+robDp(nums, start+2))
	return robMem[start]
}

func MaximalRectangle(matrix [][]byte) int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return 0
    }

    maxArea := 0
    rows, cols := len(matrix), len(matrix[0])
    height := make([]int, cols)
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if matrix[i][j] == '1' {
                height[j]++
            } else {
                height[j] = 0
            }
        }

        curr := MaxRectangle(height)
        if curr > maxArea {
            maxArea = curr
        }
    }
    return maxArea
}

/*
["1","0","1","0","0"],
["1","0","1","1","1"],
["1","1","1","1","1"],
["1","0","0","1","0"]
转换为
["3","1","3","2","2"]
*/

func MaxRectangle(height []int) int {
    n := len(height)

    stack := []int{-1}
    maxArea := 0
    for i := 0; i < n; i++ {
        for len(stack) > 1 && height[i] < height[stack[len(stack)-1]] {
            t := stack[len(stack)-1]
            fmt.Println(t)
            stack = stack[:len(stack)-1]
            w := i - stack[len(stack)-1] - 1
            area := height[t] * w
            if area > maxArea {
                maxArea = area
            }
        }
        stack = append(stack, i)
    }

    for len(stack) > 1 {
        t := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        w := n - stack[len(stack)-1] - 1
        area := height[t] * w
        if area > maxArea {
            maxArea = area
            fmt.Println(area)
        }
    }
    return maxArea
}
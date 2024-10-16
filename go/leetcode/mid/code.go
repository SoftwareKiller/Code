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
		case ch == '[':
			stack = append(stack, pair{num: num, sb: sb})
			num = 0
			sb = &strings.Builder{}
		case ch == ']':
			p := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 把当前的加入到前面的 sb 上面
			p.sb.Grow(p.num * sb.Len())
			for j := 0; j < p.num; j++ {
				p.sb.WriteString(sb.String())
			}
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

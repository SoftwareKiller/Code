package mid

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

	dp := make([]bool, len(s) + 1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j:=0; j < i; j++ {
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

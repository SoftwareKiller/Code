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
        edges = make([][]int, numCourses)
        visited = make([]int, numCourses)
        result []int
        valid = true
        dfs func(u int)
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

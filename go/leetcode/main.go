package main

import "leetcode/mid"

func main() {
	// // 42
	// num := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	// ans := hard.Trap(num)
	// fmt.Println("ans:", ans)

	// // 139
	// s := "catsandog"
	// wordDict := []string{"cats", "dog", "sand", "and", "cat"}
	// if mid.WordBreak(s, wordDict) {
	// 	fmt.Println("catsandog, Success")
	// }

	// s = "applepenapple"
	// wordDict = []string{"apple", "pen"}
	// if mid.WordBreak(s, wordDict) {
	// 	fmt.Println("applepenapple, Success")
	// }

	// // 207
	// numCourses := 3
	// prerequisites := [][]int{{1, 0}, {2, 1}}
	// mid.CanFinish(numCourses, prerequisites)

	// // 208
	// trie := mid.Constructor()
	// trie.Insert("apple")
	// fmt.Println(trie.Search("apple"))   // 返回 True
	// fmt.Println(trie.Search("app"))     // 返回 False
	// fmt.Println(trie.StartsWith("app")) // 返回 True
	// trie.Insert("app")
	// fmt.Println(trie.Search("app")) // 返回 True

	// fmt.Println(mid.CoinsChange([]int{1, 3, 5}, 11))

	// // 238
	// nums := []int{1, 5, 3, 3, 2}
	// fmt.Println(mid.ProductExceptSelf(nums))

	// nums = []int{3, 4, 5}
	// fmt.Println(mid.FindDisappearedNumbers(nums))

	// // 239
	// nums = []int{5, 3, -1, -3, 5, 3, 6, 7}
	// fmt.Println(hard.MaxSlidingWindow(nums, 3))

	// // 279
	// fmt.Println(mid.NumSquares(12))

	//394
	// str := "3[a2[c]]"
	// fmt.Println(mid.DecodeString(str))

	// 301
	// str := "()())()"
	// str := ")d))"
	// str := "(r(()()("
	// fmt.Println(hard.RemoveInvalidParentheses(str))

	// 312
	// nums := []int{3, 1, 5, 8}
	// fmt.Println(hard.MaxCoins(nums))
	// strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	// fmt.Println(mid.GroupAnagrams(strs))
	// nums := []int{-1, 0, 1, 2, -1, -4}
	// fmt.Println(mid.ThreeSum(nums))
	// s := "abcabcbb"
	// fmt.Println(mid.LengthOfLongestSubstring(s))
	// s := "cbaebabacd"
	// fmt.Println(mid.FindAnagrams(s, "abc"))

	// 560
	// nums := []int{1, -1, 0}
	// fmt.Println(mid.SubarraySum(nums, 0))

	// 56
	// fmt.Println(mid.MaxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))

	// 41
	// fmt.Println(hard.FirstMissingPositive([]int{4,1,8,5,0,3,2}))

	// 240
	// matrix := [][]int{
	// 	{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30},
	// }
	// fmt.Println(mid.SearchMatrix(matrix, 5))
	// 131
	// fmt.Println(mid.Partition("aab"))
	// 04
	//fmt.Println(hard.FindMedianSortedArrays([]int{2,2,4,4}, []int{2,2,2,4,4}))
	// 121
	// fmt.Println(mid.MaxProfit([]int{7, 1, 5, 3, 6, 4}))
	// 32
	// fmt.Println(hard.LongestValidParentheses("()(()"))
	// mid.MaximalRectangle([][]byte{
	// 	{'1', '0', '1', '0', '0'},
	// 	{'1', '0', '1', '1', '1'},
	// 	{'1', '1', '1', '1', '1'}})
	mid.MaxRectangle([]int{3, 4, 5, 2, 6})
}

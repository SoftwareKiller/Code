package main

import (
	"fmt"
	"leetcode/hard"
)

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

	// 394
	// str := "3[a]2[bc]"
	// fmt.Println(mid.DecodeString(str))

	// 301
	// str := "()())()"
	// str := ")d))"
	// str := "(r(()()("
	// fmt.Println(hard.RemoveInvalidParentheses(str))

	// 312
	nums := []int{3, 1, 5, 8}
	fmt.Println(hard.MaxCoins(nums))
}

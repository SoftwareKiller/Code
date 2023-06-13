package main

import (
	"fmt"
	"leetcode/hard"
	"leetcode/mid"
)

func main() {
	// 42
	num := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	ans := hard.Trap(num)
	fmt.Println("ans:", ans)

	// 139
	s := "catsandog"
	wordDict := []string{"cats", "dog", "sand", "and", "cat"}
	if mid.WordBreak(s, wordDict) {
		fmt.Println("catsandog, Success")
	}

	s = "applepenapple"
	wordDict = []string{"apple", "pen"}
	if mid.WordBreak(s, wordDict) {
		fmt.Println("applepenapple, Success")
	}

	// 207
	numCourses := 3
	prerequisites := [][]int{{1, 0}, {2, 1}}
	mid.CanFinish(numCourses, prerequisites)
}

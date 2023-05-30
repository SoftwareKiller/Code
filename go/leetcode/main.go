package main

import (
	"fmt"
	"leetcode/hard"
)

func main() {
	num := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	ans := hard.Trap(num)
	fmt.Println("ans:", ans)
}

package service

import (
	"fmt"
	"sync"
	"time"
)

const LEN = 64 * 1024 * 1024

func Cal() {
	arr := make([]int, LEN)

	t := time.Now().UnixMilli()
	for i := 0; i < LEN; i += 2 {
		arr[i] *= i
	}
	t2 := time.Now().UnixMilli()
	fmt.Println(t2 - t)
}

func Cal1() {
	arr := make([]int, LEN)

	t := time.Now().UnixMilli()
	for i := 0; i < LEN; i += 512 {
		arr[i] *= i
	}
	t2 := time.Now().UnixMilli()
	fmt.Println(t2 - t)
}

func Cal2() {
	size := []int32{1, 16, 512, 1024}

	for _, step := range size {
		for k := 1; k <= 18; k++ {
			t := time.Now().UnixMicro()
			n := int32(k) * step
			mem := make([]int32, n)
			for i := 0; i < 10000000; i++ {
				for j := int32(0); j < n; j += step {
					mem[j] += j
				}
			}
			t1 := time.Now().UnixMicro()
			fmt.Println("step:", step, "count:", k, "cost:", t1-t)
		}
	}
}

func Cal3() {
	row := 1024 * 48
	col := 48 * 1024
	arr := [1024 * 48][48 * 1024]int{}

	t := time.Now().UnixMilli()
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			arr[r][c] *= row * col
		}
	}

	t1 := time.Now().UnixMilli()
	for c := 0; c < col; c++ {
		for r := 0; r < row; r++ {
			arr[r][c] *= row * col
		}
	}
	t2 := time.Now().UnixMilli()
	fmt.Println(t1-t, t2-t1)
}

func Cal4() {
	t := 6
	restult := [6]int{}
	total := 16 * 1024 * 1024
	arr := make([]int, total)

	t1 := time.Now().UnixMicro()
	wait := sync.WaitGroup{}
	f := func(id int) {
		wait.Add(1)
		restult[id] = 0
		chunk := total/t + 1
		start := id * chunk
		end := min(start+chunk, total)
		c := 0
		for i := start; i < end; i++ {
			if arr[i]%2 != 0 {
				c++
				// restult[id]++
			}
		}
		restult[id] = c
		wait.Done()
	}
	for i := 0; i < t; i++ {
		go f(i)
	}
	wait.Wait()
	t2 := time.Now().UnixMicro()
	fmt.Println(t2 - t1)
}

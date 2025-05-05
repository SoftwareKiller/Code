package main

import "sync"

type mapMgr struct {
	m  map[int]int
	mu sync.RWMutex
}

var MapMgr = mapMgr{
	m: make(map[int]int),
}

func NewMapMgr() *mapMgr {
	return &MapMgr
}

func (m *mapMgr) Set(k, v int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.m[k] = v
}

func main() {
	// // BenchmarkConcurrentOpMap()
	// // BenchmarkSyncMap()
	// ch := make(chan int)
	// ch <- 42      // 发送数据
	// println(<-ch) // 接收数据
	m := NewMapMgr()
	m.Set(1, 1)
}

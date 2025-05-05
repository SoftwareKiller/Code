package main

import (
	"sync"
	"time"
)

type Map struct {
	M  map[int]int
	mu sync.RWMutex
}

func (m *Map) Get(key int) int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.M[key]
}

func (m *Map) Set(key, value int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.M[key] = value
}

func BenchmarkConcurrentOpMap() {
	m := &Map{M: make(map[int]int)}
	start := time.Now().UnixMicro()
	wg := sync.WaitGroup{}
	for i := 0; i <= 1; i++ {
		wg.Add(1)
		go func(i int) {
			for j := i * 1000; j < i*1000+1000; j++ {
				m.Set(j, j)
			}
			wg.Done()
		}(i)
	}
	// for i := 0; i < 1000000; i++ {
	// 	m.Set(i, i)
	// }

	for i := 1; i <= 12; i++ {
		wg.Add(1)
		go func(i int) {
			for j := i * 1000; j < i*1000+1000; j++ {
				m.Get(j)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	println(time.Now().UnixMicro() - start)
}

func BenchmarkSyncMap() {
	m := sync.Map{}
	start := time.Now().UnixMicro()
	wg := sync.WaitGroup{}
	for i := 0; i <= 12; i++ {
		wg.Add(1)
		go func(i int) {
			for j := i * 1000; j < i*1000+1000; j++ {
				m.Store(j, j)
			}
			wg.Done()
		}(i)
	}

	// for i := 0; i < 1000000; i++ {
	// 	m.Store(i, i)
	// }

	for i := 1; i <= 1; i++ {
		wg.Add(1)
		go func(i int) {
			for j := i * 1000; j < i*1000+1000; j++ {
				m.Load(j)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	println(time.Now().UnixMicro() - start)
}

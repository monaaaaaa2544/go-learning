package lock

import (
	"sync"
	"testing"
	"time"
)

// 互斥锁
func TestLock(t *testing.T) {
	var mu sync.Mutex

	count := 0
	for i := 0; i < 300; i++ {
		go func() {
			defer func() {
				mu.Unlock()
			}()
			mu.Lock()
			count++
		}()
	}

	time.Sleep(time.Second)
	t.Log(count)
}

// WaitGroup
func TestWaitGroup(t *testing.T) {
	var mu sync.Mutex
	var wg sync.WaitGroup

	count := 0
	for i := 0; i < 300; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mu.Unlock()
			}()
			mu.Lock()
			count++
			wg.Done()
		}()
	}

	wg.Wait()
	t.Log(count)
}

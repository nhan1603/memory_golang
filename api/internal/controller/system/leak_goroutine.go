package system

import (
	"fmt"
	"runtime"
	"time"
)

func leakGoroutine() {
	go func() {
		for {
			_ = make([]byte, 1024*1024*10) // Allocate 10 MB
			time.Sleep(100 * time.Millisecond)
		}
	}()
}

func (i impl) MonitorGoroutine() {
	runtime.GC() // Run garbage collector to clean up unused memory

	for i := 0; i < 10; i++ {
		leakGoroutine()
		time.Sleep(500 * time.Millisecond)
	}

	// Monitor memory stats periodically
	for {
		time.Sleep(5 * time.Second)
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		fmt.Printf("Alloc = %v MiB, TotalAlloc = %v MiB, Sys = %v MiB, NumGC = %v\n",
			m.Alloc/1024/1024, m.TotalAlloc/1024/1024, m.Sys/1024/1024, m.NumGC)
	}
}
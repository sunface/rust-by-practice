package talent

import (
	"bytes"
	"fmt"
	"log"
	"runtime"
	"strconv"
	"time"
)

//测量一段代码执行时间
func TraceCode() func() {
	start := time.Now()
	return func() {
		t := time.Now().Sub(start).Nanoseconds()
		fmt.Printf("运行耗时:%d(纳秒)\n", t)
	}

}

//打印当前堆栈
func PrintStack(all bool) {
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, all)

	log.Println("[FATAL] catch a panic,stack is: ", string(buf[:n]))
}

func GetStack(all bool) string {
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, all)
	return string(buf[:n])
}

// 获取常用runtime统计信息
func RuntimeStats(gc bool, heapObj bool, goroutineNum bool) []int64 {
	s := &runtime.MemStats{}
	runtime.ReadMemStats(s)

	stats := make([]int64, 5)
	if gc {
		// 上一次gc耗时
		t := s.PauseNs[(s.NumGC+255)%256]
		stats[0] = int64(t)

		// gc总次数
		num := s.NumGC
		stats[1] = int64(num)

		// 下一次gc触发时，heapalloc的大小
		ng := s.NextGC
		stats[2] = int64(ng)
	}

	if heapObj {
		ho := s.HeapObjects
		stats[3] = int64(ho)
	}

	if goroutineNum {
		ng := runtime.NumGoroutine()
		stats[4] = int64(ng)
	}

	return stats
}

func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

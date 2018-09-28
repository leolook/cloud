package goid

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	gid := Get()
	fmt.Printf("gid:%d\n", gid)
	for i := 0; i < 10; i++ {
		go func() {
			gid := Get()
			fmt.Printf("gid:%d\n", gid)
		}()
	}
}

func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gid := Get()
		fmt.Printf("gid:%d\n", gid)
	}
}

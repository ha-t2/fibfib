package fib_test

import (
	"fmt"
	"github.com/ha-t2/fibfib/fib"
	"testing"
)

func TestFib(t *testing.T) {
	if got := fib.Fib(8); got != 21 {
		t.Errorf("Fib(%d) expect %d, but got %d", 5, 21, got)
	}
}

func ExampleFib() {
	fmt.Println(fib.Fib(10))
	// Output: 55
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib.Fib(20)
	}
}

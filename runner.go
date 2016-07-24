package main

import (
	"fmt"
	"testing"
)

func functionToBeTested() int {
	var sum int
	for i := 1; i < 20000; i++ {
		sum += i
	}
	return sum
}

func main() {
	br := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			functionToBeTested()
		}
	})
	fmt.Println(br)
}

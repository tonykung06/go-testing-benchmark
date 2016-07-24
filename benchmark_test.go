package main

import "testing"

func BenchmarkExample(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			functionToBeTested()
		}
	})
}

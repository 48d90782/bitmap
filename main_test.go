package main

import "testing"

// Without inlining BenchmarkSimpleBitmapIndex-12    	  120832	      9747 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSimpleBitmapIndex(b *testing.B) {
	_, _, terrace, reservations, _, expensive := initData()

	resBitmap := make([]uint64, bitmapLength)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		andnot(terrace, expensive, resBitmap)
		and(reservations, resBitmap, resBitmap)
	}
}

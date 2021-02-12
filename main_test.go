package main

import "testing"

func BenchmarkSimpleBitmapIndex(b *testing.B) {
	_, _, terrace, reservations, _, expensive := initData()

	resBitmap := make([]byte, bitmapLength)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		andnot(terrace, expensive, resBitmap)
		and(reservations, resBitmap, resBitmap)
	}
}

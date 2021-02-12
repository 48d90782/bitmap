package main

import "math/rand"

const restaurants = 65536
const bitmapLength = restaurants / (8 * 8)

var (
	nearMetro      = make([]uint64, bitmapLength)
	privateParking = make([]uint64, bitmapLength)
	terrace        = make([]uint64, bitmapLength)
	reservations   = make([]uint64, bitmapLength)
	veganFriendly  = make([]uint64, bitmapLength)
	expensive      = make([]uint64, bitmapLength)
)

func fill(r *rand.Rand, b []uint64, probability float32) {
	for i := 0; i < len(b); i++ {
		for j := uint(0); j < 8; j++ {
			if r.Float32() < probability {
				b[i] |= 1 << j
			}
		}
	}
}

func indexes(a []uint64) []int {
	var res []int
	for i := 0; i < len(a); i++ {
		for j := 7; j > 0; j-- {
			if a[i]&(1<<uint(j)) > 0 {
				res = append(res, (7-j)+(i*8))
			}
		}
	}

	return res
}

func and(a []uint64, b []uint64, res []uint64) {
	i := 0
	if len(a) != len(b) || len(b) != len(res) {
		return
	}

loop:
	if i < len(a) {
		res[i] = a[i] & b[i]
		i++
		goto loop
	}
}

func andnot(a []uint64, b []uint64, res []uint64) {
	i := 0
	if len(a) != len(b) || len(b) != len(res) {
		return
	}

loop:
	if i < len(a) {
		res[i] = a[i] & ^ b[i]
		i++
		goto loop
	}
}

func initData() ([]uint64, []uint64, []uint64, []uint64, []uint64, []uint64) {
	r := rand.New(rand.NewSource(10))
	fill(r, nearMetro, 0.5)
	fill(r, privateParking, 0.01)
	fill(r, terrace, 0.05)
	fill(r, reservations, 0.95)
	fill(r, veganFriendly, 0.2)
	fill(r, expensive, 0.1)
	return nearMetro, privateParking, terrace, reservations, veganFriendly, expensive
}

func main() {

}

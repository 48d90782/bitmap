package main

import "math/rand"

const restaurants = 65536
const bitmapLength = restaurants / 8

var (
	nearMetro      = make([]byte, bitmapLength)
	privateParking = make([]byte, bitmapLength)
	terrace        = make([]byte, bitmapLength)
	reservations   = make([]byte, bitmapLength)
	veganFriendly  = make([]byte, bitmapLength)
	expensive      = make([]byte, bitmapLength)
)

func fill(r *rand.Rand, b []byte, probability float32) {
	for i := 0; i < len(b); i++ {
		for j := uint(0); j < 8; j++ {
			if r.Float32() < probability {
				b[i] |= 1 << j
			}
		}
	}
}

func indexes(a []byte) []int {
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

func and(a []byte, b []byte, res []byte) {
	for i := 0; i < len(a); i++ {
		res[i] = a[i] & b[i]
	}
}

func andnot(a []byte, b []byte, res []byte) {
	for i := 0; i < len(a); i++ {
		res[i] = a[i] & ^ b[i]
	}
}

func initData() ([]byte, []byte, []byte, []byte, []byte, []byte) {
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

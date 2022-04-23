package vptree

import "math/bits"

// distance returns the popcnt between h1 and h2
func distance(h1 uint64, h2 uint64) uint64 {
	return uint64(bits.OnesCount64(h1 ^ h2))
}

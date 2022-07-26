package utils

import "math/rand"

func RandUInt32(min, max uint32) uint32 {
	return uint32(rand.Intn(int(max-min))) + min
}

func RandInt(min, max int) int {
	return rand.Intn(max-min) + min
}

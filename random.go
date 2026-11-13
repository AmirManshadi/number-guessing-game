package main

import "math/rand"

func GetRandomInt(low, high uint8) uint8 {
	return uint8(rand.Int31n(int32(high)) + int32(low))
}

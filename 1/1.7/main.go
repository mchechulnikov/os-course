package main

import (
	"fmt"
	"math"
)

// x86 real mode address mapping: (segment*16 + offset) mod (2^20)
// where segment and offset — 16-bit values
func main() {
	lowOrder20BitsModValue := int(math.Pow(2, 20))
	result := 0
	for seg := 0; seg <= 65535; seg++ {
		for off := 0; off <= 65535; off++ {
			if ((seg*16 + off) % lowOrder20BitsModValue) == 0x7c00 {
				result++
			}
		}
	}
	fmt.Printf("%d", result)
}
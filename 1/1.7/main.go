package main

import (
	"fmt"
	"math"
)

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
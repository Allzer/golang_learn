package main

import (
	"fmt"
	"math"
)

func main() {
	var user_height = 1.8
	var user_kg float64 = 100
	var imt = user_kg / math.Pow(user_height, 2)
	fmt.Println((imt))
}
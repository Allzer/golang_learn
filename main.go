package main

import (
	"fmt"
	"math"
)

func main() {
	// var user_height float64 = 1.8
	// var user_kg float64 = 83

	imt := output_res(1.8, 83)
	fmt.Println(imt)
}

func output_res(user_height float64, user_kg float64) float64 {
	imt := user_kg / math.Pow(user_height/100, 2)
	return imt
}
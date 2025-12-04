package main

import (
	"fmt"
	"math"
)

func main() {
	output_res(1.8, 100)
}

func output_res(user_height float64, user_kg float64) {// указать тип/ы данных того, что будем возварщать. Чтобы указать несколько значений типы нужно указать в скобках {
    imt := user_kg / math.Pow(user_height/100, 2)
    fmt.Println(imt)
}
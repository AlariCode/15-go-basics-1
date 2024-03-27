package main

import (
	"fmt"
	"math"
)

func main() {
	var userHeight = 1.8
	var userKg float64 = 100
	var IMT = userKg / math.Pow(userHeight, 2)
	fmt.Print(IMT)
}

package main

import (
	"fmt"
	"math"
)

func main() {
	userHeight := 1.8
	userKg := 100.0
	IMT := userKg / math.Pow(userHeight, 2)
	fmt.Print(IMT)
}

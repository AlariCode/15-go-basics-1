package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	var userHeight float64
	var userKg float64
	fmt.Println("__ Калькулятор индекса массы тела __")
	fmt.Print("Введите свой рост в сантиментрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	fmt.Printf("Ваш индекс массы тела: %.0f", IMT)
}

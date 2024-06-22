package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount, years float64 = 1000, 10
	expectedReturnRate := 5.5 // thi way we can avoid type conversion

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fmt.Println("Future value is", futureValue)
}

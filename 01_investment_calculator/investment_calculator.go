package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	var investmentAmount, years float64
	expectedReturnRate := 5.5 // this way we can avoid type conversion

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount) // & means we are passing the address of the variable (pointer)

	fmt.Print("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter the years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("Future value is", futureValue)
	fmt.Println("Future Real value is", futureRealValue)
}

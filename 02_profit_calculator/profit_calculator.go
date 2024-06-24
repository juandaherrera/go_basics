package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	tax := ebt * taxRate / 100
	eat := ebt - tax
	ratio := ebt / revenue * 100

	fmt.Println("Earnings before tax is", ebt)
	fmt.Println("Tax is", tax)
	fmt.Println("Earnings after tax (profit) is", eat)
	fmt.Println("Profit margin is", ratio)
}

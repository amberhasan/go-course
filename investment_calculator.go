package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64 //default value of 0.0
	var years float64
	var expectedReturnRate float64 = 5.5

	outputText("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	outputText("Enter the number of years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future value is %.2f", futureValue)
	formattedFRV := fmt.Sprintf("Future real value is %.2f", futureRealValue)

	fmt.Print(formattedFV, formattedFRV)
}

func outputText(text string) {
	fmt.Print(text)
}
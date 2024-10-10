package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main() {
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateValues(investmentAmount, expectedReturnRate, years)
	futureValueString := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	futureRealValueString := fmt.Sprintf("Future Real Value: %.2f\n", futureRealValue)

	fmt.Print(futureValueString, futureRealValueString)
}

func calculateValues(investmentAmount float64, expectedReturnRate float64, years float64) (float64, float64) {
	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	var futureRealValue = futureValue / math.Pow(1+inflationRate/100, years)

	return futureValue, futureRealValue
}

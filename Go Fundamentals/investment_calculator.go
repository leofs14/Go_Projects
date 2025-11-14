package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount, years float64 = 1000, 10 //boa prática.
	expectedReturnRate := 5.5 //quando nao precisa especificar o tipo faça isso.
	//var years float64 = 10

	fmt.Scan()

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	inflationAdjustedValue := futureValue / math.Pow(1 + inflationRate / 100, years)

	fmt.Println(futureValue)
	fmt.Println(inflationAdjustedValue)
}

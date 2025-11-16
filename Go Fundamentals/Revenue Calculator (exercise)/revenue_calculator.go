package main

import "fmt"

// Global Variables:
var revenue float64
const taxRate= 27.5

func main() {
	// Data Input:
	inputRevenue("Insert your monthly revenue: ")

	// Calculations:
	earningBeforeTaxes, earningAfterTaxes, ratio := calculateValues(revenue)

	// Data Output:
	formattedEBT := fmt.Sprintf("Earning Before Taxes: %.1f\n", earningBeforeTaxes)
	formattedEAT := fmt.Sprintf("Earning After Taxes: %.1f\n", earningAfterTaxes)
	formattedRatio := fmt.Sprintf("Ratio: %.2f\n", ratio)

	fmt.Print(formattedEBT, formattedEAT, formattedRatio)
}

//Function managing Data Input:
func inputRevenue (text string){
	fmt.Print(text)
	fmt.Scan(&revenue)
}

//Function to Execute Calculations:
func calculateValues(revenue float64) (ebt, eat, rt float64){
	ebt = revenue

	eat = revenue - (revenue * (taxRate / 100))

	rt = ebt / eat

	return //ebt, eat and rt in that exact order.
}
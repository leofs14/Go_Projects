package main

import (
	"fmt"
	"os"
)

// Global Variables:
const taxRate= 27.5
const valueFile = "values.txt"

// File Writing Functions
func writeAllToFile(ebt, eat, rt float64) {
    text := fmt.Sprintf("EBT: %.2f\nEAT: %.2f\nRT: %.2f\n", ebt, eat, rt)
    os.WriteFile(valueFile, []byte(text), 0644)
}

func main() {
	// Data Input:
	var revenue float64
	fmt.Print("Insert your monthly revenue: ")
	fmt.Scan(&revenue)

	if revenue <= 0 {
		fmt.Println("Negative or null revenue inserted. Please insert a valid one.")
		return // Used to exit the app
	}

	// Calculations:
	earningBeforeTaxes, earningAfterTaxes, ratio := calculateValues(revenue)
	writeAllToFile(earningBeforeTaxes, earningAfterTaxes, ratio)

	// Data Output:
	formattedEBT := fmt.Sprintf("Earning Before Taxes: %.1f\n", earningBeforeTaxes)
	formattedEAT := fmt.Sprintf("Earning After Taxes: %.1f\n", earningAfterTaxes)
	formattedRatio := fmt.Sprintf("Ratio: %.2f\n", ratio)

	fmt.Print(formattedEBT, formattedEAT, formattedRatio)
}

//Function to Execute Calculations:
func calculateValues(revenue float64) (ebt, eat, rt float64){
	ebt = revenue

	eat = revenue - (revenue * (taxRate / 100))

	rt = ebt / eat

	return //ebt, eat and rt in that exact order.
}
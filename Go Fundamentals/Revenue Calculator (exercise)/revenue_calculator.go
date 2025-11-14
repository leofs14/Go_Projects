package main

import "fmt"

func main() {
	var revenue float64
	const taxRate= 27.5

	fmt.Print("Insert your Monthly Revenue: ")
	fmt.Scan(&revenue)

	earningBeforeTaxes := revenue
	
	taxes := revenue * (taxRate / 100)

	earningAfterTaxes := revenue - taxes

	ratio := earningBeforeTaxes / earningAfterTaxes

	fmt.Println(earningBeforeTaxes)
	fmt.Println(earningAfterTaxes)
	fmt.Println(ratio)
}
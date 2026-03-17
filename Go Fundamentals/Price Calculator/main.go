package main

import (
	"fmt"

	cmdmanager "example.com/pricecalc/Cmdmanager"
	prices "example.com/pricecalc/Prices"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	for _, taxRate := range taxRates {
		//fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Could not process the job")
			fmt.Println(err)
		}
	}

}
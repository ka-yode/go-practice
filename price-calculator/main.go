package main

import (
	"fmt"

	"github.com/kayode-dev/go-practice/price-calculator/filemanager"
	"github.com/kayode-dev/go-practice/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", (taxRate*100)))
		//cmd := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("Couldn't process this job")
			fmt.Println(err)
		}

	}

}

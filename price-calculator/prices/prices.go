package prices

import (
	"fmt"

	conversion "github.com/kayode-dev/go-practice/price-calculator/conversions"
	"github.com/kayode-dev/go-practice/price-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	IOManager         filemanager.FileManager
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

// constructor function to create and return a struct populated with values
func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: fm,
		TaxRate:   taxRate,
	}
}

/*
method attached to the TaxIncludedpriceStruct because it is a receiver function
denoted by the (job TaxIncludedPiceJob) before the function name
when we have a synta like this we can access the function anywhere an instance of the struct has been created
*/
func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	fmt.Println(result)
	job.TaxIncludedPrices = result
	job.IOManager.WriteJson(job)
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := job.IOManager.ReadFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println(err)
		return
	}
	job.InputPrices = prices
}

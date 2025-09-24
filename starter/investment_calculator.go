package main

import (
	"fmt"
	"math"
)

func main()  {
	const inflationRAte =2.5

	var investMentAmount, years, expectedReturnRate float64 

	fmt.Println("please enter your investment amount");
	fmt.Scan(&investMentAmount)
	fmt.Println("how many years is this investment going to run for?");
	fmt.Scan(&years)
	fmt.Println("what is the expected return rate?");
	fmt.Scan(&expectedReturnRate)

	var futureValue =investMentAmount* math.Pow(1 + expectedReturnRate/100, years)
	var futureRealValue = investMentAmount / math.Pow(1+ inflationRAte/100, years) 
fmt.Println("Your future investment is ", futureValue)
fmt.Println("Your future real investment value is ", futureRealValue)
}
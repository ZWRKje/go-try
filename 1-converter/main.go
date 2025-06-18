package main

import (
	"fmt"
	"slices"
)

var currencys = []string{"USD", "EUR", "RUB"}

const USDTOEUR = 0.88
const USDTORUB = 79.86
const EURTORUB = (1 / USDTOEUR) * USDTORUB

func main() {
	var convertMap = map[string]float64{
		"USD:EUR": USDTOEUR,
		"USD:RUB": USDTORUB,
		"EUR:USD": (1 / USDTOEUR),
		"EUR:RUB": EURTORUB,
		"RUB:USD": (1 / USDTORUB),
		"RUB:EUR": (1 / EURTORUB),
	}
	baseCurrency, convertCurrency, money := getConvertedParams()

	res := convert(&convertMap, money, baseCurrency, convertCurrency)

	fmt.Printf("%.2f %s => %.2f %s", money, baseCurrency, res, convertCurrency)
}

func getConvertedParams() (string, string, float64) {
	var money float64

	baseCurrency := "nil"
	convertCurrency := "nil"

	supMessage := fmt.Sprintf("Supported currency: %v", currencys)

	baseCurrency = inputBaseCurrency(supMessage)
	convertCurrency = inputTargetCurrency(supMessage, baseCurrency)
	money = inputAmount()

	return baseCurrency, convertCurrency, money
}

func inputBaseCurrency(supMessage string) string {
	var baseCurrency string
	for {
		fmt.Printf("Input base currency. %s\n", supMessage)
		_, err := fmt.Scanf("%s", &baseCurrency)

		if err != nil {
			fmt.Println(err)
			continue
		}

		if !slices.Contains(currencys, baseCurrency) {
			fmt.Println("Unsupported currency")
			continue
		}
		return baseCurrency
	}
}

func inputTargetCurrency(supMessage string, baseCurrency string) string {
	var convertCurrency string
	for {
		fmt.Printf("Input convert currency. %s\n", supMessage)
		_, err := fmt.Scanf("%s", &convertCurrency)

		if err != nil {
			fmt.Println(err)
			continue
		}

		if !slices.Contains(currencys, convertCurrency) || convertCurrency == baseCurrency {
			fmt.Println("Unsupported currency")
			continue
		}
		return convertCurrency
	}

}

func inputAmount() float64 {
	var money float64

	for {
		fmt.Println("Input amount of money: ")
		_, err := fmt.Scanf("%f", &money)

		if err != nil {
			fmt.Println(err)
			continue
		}

		if money <= 0 {
			fmt.Println("Incorrect money amount")
			continue
		}

		return money
	}
}

func convert(convertMap *map[string]float64, money float64, baseCurrency string, convertCurrency string) float64 {
	key := fmt.Sprintf("%s:%s", baseCurrency, convertCurrency)
	return money * ((*convertMap)[key])
}

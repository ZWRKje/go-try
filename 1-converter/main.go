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

	baseCurrency, convertCurrency, money := getConvertedParams()

	res := convert(money, baseCurrency, convertCurrency)

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

func convert(money float64, baseCurrency string, convertCurrency string) float64 {
	switch {
	case baseCurrency == "USD" && convertCurrency == "RUB":
		return money * USDTORUB
	case baseCurrency == "USD" && convertCurrency == "EUR":
		return money * USDTOEUR
	case baseCurrency == "EUR" && convertCurrency == "USD":
		return money * (1 / USDTOEUR)
	case baseCurrency == "EUR" && convertCurrency == "RUB":
		return money * EURTORUB
	case baseCurrency == "RUB" && convertCurrency == "USD":
		return money * (1 / USDTORUB)
	case baseCurrency == "RUB" && convertCurrency == "EUR":
		return money * (1 / EURTORUB)
	default:
		return money
	}
}

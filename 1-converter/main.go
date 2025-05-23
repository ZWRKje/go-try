package main

import "fmt"

func main() {
	const USDTOEUR = 0.88
	const USDTORUB = 79.86
	const EURTORUB = (1 / USDTOEUR) * USDTORUB

	getInput()
}

func getInput() (string, string, float64) {
	var baseCurrency string
	var convertCurrency string
	var money float64

	fmt.Scanf("Input base currency: %s", baseCurrency)
	fmt.Scanf("Input convert currency: %s", convertCurrency)
	fmt.Scanf("Input amount of money: %f", money)

	return baseCurrency, convertCurrency, money
}

func convert(money float64, convertCurrency string, baseCurrency string) {

}

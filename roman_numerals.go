package main

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

// declaring rules about roman numerals as data,
// rather than hiding them inside a switch / case
// inside of the algorithm.
// switch statements should normally be a red flag that we might
// be capturing a concept or data inside imperative code, when
// it might be better captured in a class structure instead.
var AllRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {

	var result strings.Builder

	for _, numeral := range AllRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

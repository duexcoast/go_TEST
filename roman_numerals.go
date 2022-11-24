package main

import "strings"

type RomanNumeral struct {
	Value int
	Symbol string
}

var AllRomanNumerals = []RomanNumeral {
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman (arabic int) string {

	var result strings.Builder

	for _, numeral := range AllRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}
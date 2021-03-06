package main

import (
	"fmt"
	"strconv"

	"github.com/bdeleonardis1/eventtest/events"
)

type parity int
const (
	even parity = iota
	odd
)

func main() {
	fmt.Print("Enter a number 1-5: ")
	var textNum string
	_, _ = fmt.Scanln(&textNum)
	par := getParity(textNum)

	fmt.Printf("%v is an %v number\n", textNum, parityToString(par))
}

func getParity(textNum string) parity {
	// 1 is the most common number checked so we have an optimized check here.
	if textNum == "1" {
		events.EmitEvent(events.NewEvent("1Optimization"))
		return odd
	}

	if len(textNum) == 1 {
		par, _ := optimizedSingleDigitParity(textNum)
		return par
	}

	num := convertToNumber(textNum)
	if len(textNum) == 2 && num < 0 {
		par, _ := optimizedNegativeSingleDigit(num)
		return par
	}

	events.EmitEvent(events.NewEvent("Modding"))
	if num % 2 == 0 {
		return even
	}
	events.EmitEvent(events.NewEvent("TheVeryEnd"))
	return odd
}

// optimizedSingleDigitParity returns the parity of all single digit numbers.
// If the number has more than one digit it returns an error.
func optimizedSingleDigitParity(textNum string) (parity, error) {
	events.EmitEvent(events.NewEvent("OptimizedSingleDigit"))
	switch textNum {
	case "1", "3", "5", "7", "9":
		return odd, nil
	case "0", "2", "4", "6", "8":
		return even, nil
	default:
		return 0, fmt.Errorf("%v is not a single digit number", textNum)
	}
}

func convertToNumber(textNum string) int {
	events.EmitEvent(events.NewEvent("convertToNumber"))
	num, _ := strconv.Atoi(textNum)
	return num
}

func optimizedNegativeSingleDigit(num int) (parity, error) {
	events.EmitEvent(events.NewEvent("OptimizedNegativeSingleDigit"))
	switch num {
	case -1, -3, -5, -7, -9:
		return odd, nil
	case -2, -4, -6, -8:
		return even, nil
	default:
		return 0, fmt.Errorf("%v is not a single digit negative number", num)
	}
}

func parityToString(par parity) string {
	if par == even {
		return "even"
	}
	return "odd"
}



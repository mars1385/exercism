package allyourbase

import (
	"errors"
	"math"
	"slices"
)

const (
	LowInputBase  string = "input base must be >= 2"
	LowOutputBase string = "output base must be >= 2"
	InvalidDigit  string = "all digits must satisfy 0 <= d < input base"
)

const lowestBase int = 2

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {

	var result []int

	if inputBase < lowestBase {
		return nil, errors.New(LowInputBase)
	}

	if outputBase < lowestBase {
		return nil, errors.New(LowOutputBase)
	}

	var temp int
	for i, d := range inputDigits {
		if d < 0 || d >= inputBase {
			return nil, errors.New(InvalidDigit)
		}

		temp += d * int(math.Pow(float64(inputBase), float64(len(inputDigits)-1-i)))
	}

	if temp == 0 {
		return []int{0}, errors.New(LowOutputBase)
	}

	for temp > 0 {
		result = append(result, temp%outputBase)

		temp = temp / outputBase

	}
	slices.Reverse(result)
	return result, nil
}

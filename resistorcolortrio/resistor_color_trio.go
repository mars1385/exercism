package resistorcolortrio

import (
	"math"
	"strconv"
)

func convertColorToValue(color string) int {

	switch color {
	case "brown":
		return 1
	case "red":
		return 2
	case "orange":
		return 3
	case "yellow":
		return 4
	case "green":
		return 5
	case "blue":
		return 6
	case "violet":
		return 7
	case "grey":
		return 8
	case "white":
		return 9
	}
	return 0
}

func Label(colors []string) string {

	result := ""

	for index, color := range colors {

		if index == 2 {
			res, _ := strconv.Atoi(result)

			result = strconv.Itoa(res * int(math.Pow(10, float64(convertColorToValue(color)))))

		}
		if index <= 1 {
			result += strconv.Itoa(convertColorToValue(color))
		}
	}

	res, _ := strconv.Atoi(result)

	valueType := " ohms"

	if res != 0 && res%1000 == 0 {
		valueType = " kiloohms"
		result = strconv.Itoa(res / 1000)

	}

	if res != 0 && res%1000000 == 0 {
		valueType = " megaohms"
		result = strconv.Itoa(res / 1000000)

	}

	if res != 0 && res%1000000000 == 0 {
		valueType = " gigaohms"
		result = strconv.Itoa(res / 1000000000)

	}

	return result + valueType
}

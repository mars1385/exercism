package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

func findAndMerge(list FreqMap, text string) {

	res := Frequency(text)

	for key, val := range res {

		list[key] += val

	}

}
func ConcurrentFrequency(texts []string) FreqMap {

	var result = FreqMap{}

	for _, text := range texts {

		findAndMerge(result, text)

	}

	return result

}

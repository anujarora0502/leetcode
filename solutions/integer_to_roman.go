package solutions

import "math"

func IntToRoman(num int) string {
	place := 0
	result := ""
	for num != 0 {
		result = convertToRoman(num%10, place) + result
		num /= 10
		place++
	}
	return result
}

func convertToRoman(num int, place int) string {
	result := ""
	roman := map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
	}

	if num == 1 || num == 5 {
		result += roman[int(float64(num)*math.Pow(float64(10), float64(place)))]
	} else if num == 4 {
		result += roman[int(math.Pow(float64(10), float64(place)))] + roman[int(float64(5)*math.Pow(float64(10), float64(place)))]
	} else if num == 9 {
		result += roman[int(math.Pow(float64(10), float64(place)))] + roman[int(math.Pow(float64(10), float64(place+1)))]
	} else if num > 1 && num < 4 {
		for i := 0; i < num; i++ {
			result += roman[int(math.Pow(float64(10), float64(place)))]
		}
	} else if num > 5 && num < 9 {
		result += roman[int(float64(5)*math.Pow(float64(10), float64(place)))]
		for i := 5; i < num; i++ {
			result += roman[int(math.Pow(float64(10), float64(place)))]
		}
	}

	return result
}

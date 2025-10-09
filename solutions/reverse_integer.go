package solutions

import (
	"fmt"
	"math"
)

func ReverseInteger(x int32) int32 {
  n := int32(x)
	var isItPositive bool
	if n >= 0 {
		isItPositive = true
	}

  if !isItPositive {
      n = n * -1
  }

	multiplier := int32(math.Pow(10, float64(calcLength(n)) - 1))
	result := int32(0);
	
	for n != 0 {
		result += (n%10)*multiplier
		if result < 0 || result/multiplier != n%10{
			return 0;
		}
		n /= 10
		multiplier /= 10
		fmt.Println(result)
	} 
	if isItPositive {
		if result >= 0 {
			return int32(result)
		}else {
			return 0
		}
	}else {
		if result >= 0 {
			return -1 * int32(result)
		}else {
			return 0
		}
	}
}

func calcLength (n int32) int32{
	length := int32(0)
	for n != 0 {
		length += 1
		n = n/10
	}
	return length
}
package solutions

func NumWaterBottles2(numBottles int, numExchange int) int {
	totalNumberOfBottles := numBottles

	for numBottles >= numExchange {
        totalNumberOfBottles++
        numBottles -= numExchange
		numBottles++
		numExchange++
	}

	return totalNumberOfBottles
}
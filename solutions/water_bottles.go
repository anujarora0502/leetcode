package solutions

func NumWaterBottles(numBottles int, numExchange int) int {
	totalNumberOfBottles := numBottles

	for numBottles >= numExchange {
        totalNumberOfBottles += numBottles/numExchange
        numBottles = (numBottles/numExchange) + (numBottles%numExchange)
	}

	return totalNumberOfBottles
}
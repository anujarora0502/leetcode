package solutions

import "fmt"

func CanCompleteCircuit(gas []int, cost []int) int {
	for i := 0; i < len(gas); i++ {
		fmt.Println("i ---Start----", i)
		fuel := gas[i]
		var j int
		if i == len(gas)-1 {
			j = 0
		} else {
			j = i + 1
		}

		if fuel-cost[i] < 0 {
			continue
		} else {
			fuel -= cost[i]
		}

		for {
			fmt.Println("Fuel", fuel)
			fuel += gas[j]
			if fuel-cost[j] < 0 {
				break
			} else if j == i {
				return i
			}
			fuel -= cost[j]
			if j == len(gas)-1 {
				j = 0
			} else {
				j++
			}
		}
	}
	return -1
}

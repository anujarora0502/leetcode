package solutions

func Insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0)
	index := 0
    flag := false

	for index < len(intervals) {
		if intervals[index][0] <= newInterval[0] && newInterval[0] <= intervals[index][1] {
			start := intervals[index][0]
			end := 0
			for index < len(intervals) {
                if intervals[index][0] > newInterval[1] {
					if index == 0 {
						end = max(newInterval[1], intervals[index][1])
					} else if newInterval[1] > intervals[index-1][1] {
						end = newInterval[1]
					} else {
						end = intervals[index-1][1]
					}
					break
				}
				index++
			}
			arr := make([]int, 2)
			arr[0] = start
            if end == 0 {
                arr[1] = max(intervals[index-1][1], newInterval[1])
            }else {
                arr[1] = end
            }
			result = append(result, arr)
			result = append(result, intervals[index:]...)
            flag = true
			return result
		} else if newInterval[0] < intervals[index][0] {
			start := newInterval[0]
			end := 0
			for index < len(intervals) {
				if intervals[index][0] > newInterval[1] {
					if index == 0 {
						end = newInterval[1]
					} else if newInterval[1] > intervals[index-1][1] {
						end = newInterval[1]
					} else {
						end = intervals[index-1][1]
					}
					break
				}
				index++
			}
			arr := make([]int, 2)
			arr[0] = start
            if end == 0 && index != 0{
                arr[1] = max(intervals[index-1][1], newInterval[1])
            }else {
                arr[1] = end
            }
			result = append(result, arr)
			result = append(result, intervals[index:]...)
            flag = true
			return result
		} else {
			result = append(result, intervals[index])
		}
		index++
	}

    if !flag {
        result = append(result, newInterval)
    }

	return result
}
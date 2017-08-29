package simplecsv

// countIndexesInSlices counts how many times an integrer shows up in one or more slices
func countIndexesInSlices(indexes [][]int) map[int]int {
	valuesMap := map[int]int{}
	var valueExists bool
	for _, toCount := range indexes {
		for _, v := range toCount {
			_, valueExists = valuesMap[v]
			if valueExists != true {
				valuesMap[v] = 1
			} else {
				valuesMap[v]++
			}
		}
	}
	return valuesMap
}

// showsMoreThanTimes returns an array of ints they show more than the min time
func showsMoreThanTimes(valuesMap map[int]int, min int) []int {
	result := []int{}
	for k, v := range valuesMap {
		if v >= min {
			result = append(result, k)
		}
	}
	return result
}

// contains checks if a number exits in a slice
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// OrIndex operates an OR operator in the indexes
func OrIndex(indexes ...[]int) []int {
	minTimes := 1
	base := [][]int{}
	for _, index := range indexes {
		base = append(base, index)
	}
	valuesMap := countIndexesInSlices(base)
	result := showsMoreThanTimes(valuesMap, minTimes)
	return result
}

// AndIndex operates an AND operator in the indexes
func AndIndex(indexes ...[]int) []int {
	minTimes := len(indexes)
	base := [][]int{}
	for _, index := range indexes {
		base = append(base, index)
	}
	valuesMap := countIndexesInSlices(base)
	result := showsMoreThanTimes(valuesMap, minTimes)
	return result
}

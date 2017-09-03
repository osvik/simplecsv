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

// NotIndex negates the index between a min value and a max value
// The min value tipically is either 0 or 1 (without and with headers)
// The max value tipically is the csv length - 1 (number of rows -1).
// A table with a header and 2 rows (length = 3) can have indexes with values of 1 and 2
func NotIndex(index []int, min, max int) []int {

	counter := make(map[int]bool)
	for i := min; i <= max; i++ {
		counter[i] = false
	}

	for _, v := range index {
		counter[v] = true
	}

	negativeIndex := []int{}
	for i := min; i <= max; i++ {
		if counter[i] == false {
			negativeIndex = append(negativeIndex, i)
		}
	}

	return negativeIndex
}

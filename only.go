package simplecsv

// OnlyThisRows removes all rows that are not in the index and sorts the csv by the index order
// All rows must exist or it fails
func (s SimpleCsv) OnlyThisRows(rowsIndex []int, header bool) (SimpleCsv, bool) {
	ok := true

	lenghtS := len(s)
	for _, v := range rowsIndex {
		if v < 0 || v >= lenghtS {
			ok = false
			return s, ok
		}
	}

	newCsv := SimpleCsv{}
	if header == true {
		headers := s.GetHeaders()
		newCsv = append(newCsv, headers)
	}

	var rowToAdd []string
	for _, g := range rowsIndex {
		rowToAdd, ok = s.GetRow(g)
		if ok == false {
			return newCsv, ok
		}
		newCsv = append(newCsv, rowToAdd)
	}
	return newCsv, ok
}

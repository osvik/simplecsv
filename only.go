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
		rowToAdd, _ = s.GetRow(g)
		newCsv = append(newCsv, rowToAdd)
	}
	return newCsv, ok
}

// OnlyThisFields returns a simplecsv with the fields
func (s SimpleCsv) OnlyThisFields(fields []string) (SimpleCsv, bool) {

	var row []string
	var fieldToAdd string
	lenghS := len(s)

	newCsv := CreateEmpyCsv(fields)

	for i := 1; i < lenghS; i++ {
		for _, n := range fields {
			fieldToAdd, _ = s.GetCellByField(n, i)
			row = append(row, fieldToAdd)
		}
		newCsv, _ = newCsv.AddRow(row)
		row = []string{}

	}

	return newCsv, true
}

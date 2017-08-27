package simplecsv

// AddEmptyColumn ads an empty column at the end of the csv
func (s SimpleCsv) AddEmptyColumn(columnName string) (SimpleCsv, bool) {
	var added bool

	s[0] = append(s[0], columnName)

	for i := 1; i < len(s); i++ {
		s[i] = append(s[i], "")
	}
	added = true
	return s, added
}

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

// RemoveColumn removes the column in position X
func (s SimpleCsv) RemoveColumn(columnPosition int) (SimpleCsv, bool) {
	var removed bool

	if columnPosition >= 0 && columnPosition < len(s[0]) {
		s[0] = append(s[0][:columnPosition], s[0][columnPosition+1:]...)
		for i := 1; i < len(s); i++ {
			s[i] = append(s[i][:columnPosition], s[i][columnPosition+1:]...)
		}
		removed = true
	} else {
		removed = false
	}

	return s, removed
}

// RemoveColumnByName removes column by name
func (s SimpleCsv) RemoveColumnByName(columnName string) (SimpleCsv, bool) {
	var removed bool
	columnPosition := s.GetHeaderPosition(columnName)
	s, removed = s.RemoveColumn(columnPosition)
	return s, removed
}

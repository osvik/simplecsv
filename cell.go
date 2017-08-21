package simplecsv

// GetCell Returns the value of a cell by position.
// If the cell does not exist, it returns the second value as false.
func (s SimpleCsv) GetCell(column int, row int) (string, bool) {
	if column >= 0 && column < len(s[0]) && row >= 0 && row < len(s) {
		return s[row][column], true
	}
	return "", false

}

// GetCellByField returns the value of a cell by column name.
// If the cell does not exist, it returns the second value as false.
func (s SimpleCsv) GetCellByField(columnName string, row int) (string, bool) {
	column := s.GetHeaderPosition(columnName)
	if column >= 0 && row < len(s) {
		return s[row][column], true
	}
	return "", false
}

// SetCell changes the value of a cell
func (s SimpleCsv) SetCell(column int, row int, value string) (SimpleCsv, bool) {
	if column >= 0 && column < len(s[0]) && row >= 0 && row < len(s) {
		s[row][column] = value
		return s, true
	}
	return s, false
}

// SetCellByField changes the value of a cell with a specific column
func (s SimpleCsv) SetCellByField(columnName string, row int, value string) (SimpleCsv, bool) {
	var changed bool
	if s.GetHeaderPosition(columnName) == -1 {
		return s, false
	}
	columnPosition := s.GetHeaderPosition(columnName)
	s, changed = s.SetCell(columnPosition, row, value)
	return s, changed
}

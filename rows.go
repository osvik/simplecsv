package simplecsv

// GetNumberRows returns the number of rows, including the header row
func (s SimpleCsv) GetNumberRows() int {
	return len(s)
}

// GetRow returns the row rowNumber
// If rowNumber does not exist, it returns an empty slice and false
func (s SimpleCsv) GetRow(rowNumber int) ([]string, bool) {
	if rowNumber >= 0 && rowNumber < len(s) {
		return s[rowNumber], true
	}
	return []string{}, false
}

// GetRowAsMap returns the row as a map
// If the row does not exist, returns second value as false
func (s SimpleCsv) GetRowAsMap(rowNumber int) (map[string]string, bool) {
	RowAsMap := make(map[string]string)
	if rowNumber >= 0 && rowNumber < len(s) {
		headers := s.GetHeaders()

		for k, v := range headers {
			RowAsMap[v] = s[rowNumber][k]
		}
		return RowAsMap, true
	}
	RowAsMap = map[string]string{"": ""}
	return RowAsMap, false
}

// AddRow ads a row at the end of the csv
func (s SimpleCsv) AddRow(rowValue []string) (SimpleCsv, bool) {
	var added bool
	if len(rowValue) == len(s[0]) {
		s = append(s, rowValue)
		added = true
	} else {
		added = false
	}
	return s, added
}

// AddRowFromMap adds map values to a row.
// Ignores keys with unexistiging headers
// Fills blank where the key does not exists
func (s SimpleCsv) AddRowFromMap(rowValue map[string]string) (SimpleCsv, bool) {
	headers := s.GetHeaders()
	var added bool
	var newRow []string
	for _, v := range headers {
		value, _ := rowValue[v]
		newRow = append(newRow, value)
	}
	s, added = s.AddRow(newRow)
	return s, added
}

// SetRow updates a row in the csv
func (s SimpleCsv) SetRow(rowNumber int, rowValue []string) (SimpleCsv, bool) {
	if len(rowValue) != len(s[0]) || rowNumber >= len(s) || rowNumber < 0 {
		return s, false
	}
	s[rowNumber] = rowValue
	return s, true
}

// SetRowFromMap replaces a row by the maps value
// Ignores keys with unexistiging headers
// Fills blank where the key does not exists
func (s SimpleCsv) SetRowFromMap(rowNumber int, rowValue map[string]string) (SimpleCsv, bool) {
	headers := s.GetHeaders()
	var newRow []string
	var modified bool
	for _, v := range headers {
		value, _ := rowValue[v]
		newRow = append(newRow, value)
	}
	s, modified = s.SetRow(rowNumber, newRow)
	return s, modified
}

// DeleteRow deletes the row
func (s SimpleCsv) DeleteRow(rowNumber int) (SimpleCsv, bool) {
	if rowNumber >= 0 && rowNumber < len(s) {
		s = append(s[:rowNumber], s[rowNumber+1:]...)
		return s, true
	}
	return s, false
}

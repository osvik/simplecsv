package simplecsv

// GetHeaders returns the headers as a slice
func (s SimpleCsv) GetHeaders() []string {
	return s[0]
}

// GetHeader returns the header name in a position
// Returns false as the second value if it does not exist
func (s SimpleCsv) GetHeader(columnPosition int) (string, bool) {
	if columnPosition >= 0 && columnPosition < len(s[0]) {
		return s[0][columnPosition], true
	}
	return "", false
}

// GetHeaderPosition returns the header position,
// returns -1 if it does not exist
func (s SimpleCsv) GetHeaderPosition(columnName string) int {
	headers := s.GetHeaders()
	for k, v := range headers {
		if v == columnName {
			return k
		}
	}
	return -1
}

// RenameHeader renames the header
// Returns the second value as false if it did't found the oldHeader
func (s SimpleCsv) RenameHeader(oldHeader string, newHeader string) (SimpleCsv, bool) {
	headerPosition := s.GetHeaderPosition(oldHeader)
	if headerPosition >= 0 && headerPosition < len(s) {
		s[0][headerPosition] = newHeader
		return s, true
	}
	return s, false

}

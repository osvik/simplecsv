package simplecsv

import (
	"encoding/csv"
	"os"
)

// SimpleCsv is the type for simple csv
type SimpleCsv [][]string

// CreateEmpyCsv creates an empty CSV with the headers passed as a slice
func CreateEmpyCsv(columnNames []string) SimpleCsv {
	a := make([][]string, 1)
	a[0] = make([]string, len(columnNames))
	a[0] = columnNames
	return a
}

// ReadCsvFile reads a file and returns a [][]string slice
func ReadCsvFile(filename string) (SimpleCsv, bool) {
	ok := true
	file, err := os.Open(filename)
	if err != nil {
		ok = false
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ','
	// lineCount := 0

	allRecords, err := reader.ReadAll()
	if err != nil {
		ok = false
	}
	return allRecords, ok
}

// WriteCsvFile writes the slice to a file
func (s SimpleCsv) WriteCsvFile(filename string) bool {
	ok := true
	file, err := os.Create(filename)
	if err != nil {
		ok = false
	}
	w := csv.NewWriter(file)
	w.WriteAll(s)
	if err := w.Error(); err != nil {
		ok = false
	}
	w.Flush()
	return ok
}

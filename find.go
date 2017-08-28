package simplecsv

import (
	"regexp"
	"strings"
)

// FindInColumn returns a slice with the rownumbers where the "word" is in the  columnPosition
// If the column is not valid it returns an empty slice and a second false value
func (s SimpleCsv) FindInColumn(columnPosition int, word string) ([]int, bool) {
	var validColumn bool
	results := []int{}

	if columnPosition < 0 || columnPosition >= len(s[0]) {
		validColumn = false
		return results, validColumn
	}

	validColumn = true
	numberOfRows := len(s)
	for i := 0; i < numberOfRows; i++ {
		if strings.ToLower(s[i][columnPosition]) == strings.ToLower(word) {
			results = append(results, i)
		}
	}

	return results, validColumn

}

// FindInField returns a slice with the rownumbers where the "word" is in the column name
// If the column is not valid it returns an empty slice and a second false value
func (s SimpleCsv) FindInField(columnName string, word string) ([]int, bool) {

	columnPosition := s.GetHeaderPosition(columnName)

	var validColumn bool
	results := []int{}

	if columnPosition == -1 {
		validColumn = false
		return results, validColumn
	}

	validColumn = true
	numberOfRows := len(s)
	for i := 1; i < numberOfRows; i++ {
		if strings.ToLower(s[i][columnPosition]) == strings.ToLower(word) {
			results = append(results, i)
		}
	}

	return results, validColumn

}

// MatchInColumn returns a slice with the rownumbers where the regular expression applies in the  columnPosition
// If the column or regular expression are not valid it returns an empty slice and a second false value
func (s SimpleCsv) MatchInColumn(columnPosition int, regularexpression string) ([]int, bool) {
	var ok bool
	results := []int{}

	r, u := regexp.Compile(regularexpression)

	if u != nil {
		ok = false
		return results, ok
	}
	if columnPosition < 0 || columnPosition >= len(s[0]) {
		ok = false
		return results, ok
	}

	ok = true
	numberOfRows := len(s)
	for i := 0; i < numberOfRows; i++ {
		if r.MatchString(s[i][columnPosition]) {
			results = append(results, i)
		}
	}

	return results, ok

}

// MatchInField returns a slice with the rownumbers where the regular expression applies in the column name
func (s SimpleCsv) MatchInField(columnName string, regularexpression string) ([]int, bool) {

	columnPosition := s.GetHeaderPosition(columnName)
	var ok bool
	results := []int{}

	r, u := regexp.Compile(regularexpression)

	if columnPosition == -1 {
		ok = false
		return results, ok
	}

	if u != nil {
		ok = false
		return results, ok
	}

	ok = true
	numberOfRows := len(s)
	for i := 1; i < numberOfRows; i++ {
		if r.MatchString(s[i][columnPosition]) {
			results = append(results, i)
		}
	}
	return results, ok

}

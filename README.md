# Simplecsv

The **Go** package **simplecsv** is a simple mini-library to handle **csv** files. I'm building it to help me writing small command line scripts. Maybe it's useful to someone else as well.

**Some notes:**

* all **read** methods return the value in the csv and a second true/false value that is true if the value exists
* all **write** methods that change the csv return the changed csv and a true/false value if the operation was sucessful
* all cells are **strings**

## Examples

### Csv file

Simplecsv works with comma separated csv files.

#### Read

Reads file and parses as a SimpleCsv object. `fileRead` is false if there's an error reading the file or parsing the CSV.

```go
var x simplecsv.SimpleCsv
var fileRead bool
x, fileRead = simplecsv.ReadCsvFile("my1file.csv")
```

#### Create

Create empty file and define cssv headers:

```go
var u simplecsv.SimpleCsv
u = simplecsv.CreateEmpyCsv([]string{"Age", "Gender", "ID"})
```

#### Write

Write the SimpleCsv object to my2file.csv. If there's an error, `wasWritten` is false.

```go
wasWritten := u.WriteCsvFile("my2file.csv")
```

### Headers

The cells of the first row are considered headers.

#### Get

Get all headers:

```go
headers := x.GetHeaders()
```

Get header at position one (second position as it starts from 0):

```go
headerName, headerExists := x.GetHeader(1)
```

Get header position: (it returns `-1` if the header does not exist)

```go
position := x.GetHeaderPosition("Gender")
```

#### Rename

Rename header: (old header, new header)

```go
x, headerExists := x.RenameHeader("ID", "IDnumber")
```

`headerExists` is false if the old header does not exist.

### Rows

#### Get

Get number of rows:

```go
numberOfRows := x.GetNumberRows()
```

Get second row:

```go
row, rowExists := x.GetRow(1)
```

Get second row as a map:

```go
row, rowExists := x.GetRowAsMap(1)
```

#### Add

Add a **slice** to a row. The slice must have the same size as the CSV number of columns. If not wasSucessful is false.

```go
var wasSucessful bool
x, wasSucessful = x.AddRow([]string{"24", "M", "2986732"})
```

Add row from **map**: (If the map keys don't exist as columns, the value will be discarded. If a key does not exist, it will create empty cells.)

```go
mymap := make(map[string]string)
mymap["Age"] = "62"
mymap["Gender"] = "F"
mymap["ID"] = "6463246"
var wasAdded bool 
x, wasAdded = x.AddRowFromMap(mymap)
```

#### Set

Set second row (1) from a **slice**. The length of the slice must be the same as the number of columns and the row must already exist. If there’s an error `wasSet` is false.

```go
var wasSet bool
x, wasSet = x.SetRow(1, []string{"45", "F", "8356138"})
```

Set second row from **map**: If the map keys don't exist as columns, the value will be discarded. If a key does not exist, it will create empty cells.)

```go
mymap2 := make(map[string]string)
mymap2["Age"] = "62"
mymap2["Gender"] = "F"
mymap2["ID"] = "6463246"
var wasSet bool
x, wasSet = x.SetRowFromMap(1, mymap2)
```

#### Update

Unlike `SetRowFromMap`, `UpdateRowCellsFromMap` does not erase the cells value just because the column names are not keys in the  map. It updates the cells that have the column name in the map and maintains the value of all the others.

To update the age in row 1:

```go
mymap3 := make(map[string]string)
mymap3["Age"] = "63"
var wasUpdated bool
x, wasUpdated = x.UpdateRowCellsFromMap(1, mymap3)
```

#### Delete

Delete second row: (If the row number is invalid, `wasDeleted` is false)

```go
var wasDeleted bool
x, wasDeleted = x.DeleteRow(1)
```

### Column

#### Add

To add a column at the end of the CSV:

```go
var wasSucessful bool
x, wasSucessful = x.AddEmptyColumn("NewColumn")
```

#### Remove

To remove a column at position 1 (second column, because it's zero based):

```go
var wasRemoved bool
x, wasRemoved = x.RemoveColumn(1)
```

To remove a column by name:

```go
var wasRemoved bool
x, wasRemoved = x.RemoveColumnByName("Gender")
```

### Cells

#### Get

Get value of the cell in the second column, second row:

```go
cellValue, cellExists := x.GetCell(1, 1)
```

Get the value of the cell in the column *Age*, second row:

```go
cellValue, cellExists := x.GetCellByField("Age", 1)
```

#### Set

Changes the value of the cell in the first column (0) and the second row (1) to "27":

```go
var wasChanged bool 
x, wasChanged = x.SetCell(0, 1, "27")
```

The same, using the column name instead of the column position:

```go 
var wasChanged bool 
x, wasChanged = x.SetCellByField("Age", 1, "27")
```

### Find

#### Find in column

Find the word "27" in the first column (column 0):

```go 
rowsWithWord, validColumn := x.FindInColumn(0, "27")
```

It returns a slice of rownumbers (int) where you can find the word in the column position. Please note that in simplecsv all cells are strings.

If it doesn't find the value, it returns an empty slice.

In `FindInColumn` and in `FindInField` case does not matter. Foo = foo = FOO.

#### Find in field

The same as `FindInColumn` but using a column/field name instead of position. Please note that `FindInField`, unlike `FindInColumn` never includes the header in the search result.

```go 
rowsWithWord, validFieldName := x.FindInField("Age", "27")
```

If the field name does not exist, the second value returned (`validFieldName`) is false.

#### Match in column

Find where results match a regular expression in the third column (column 2):

```go 
rowsWithWord, areParamsOk := x.MatchInColumn(2, "p([a-z]+)ch$")
```

Use ^ and $ in the regular expression to match exact results.

#### Match in field

Same as with MatchInColumn, but with a field (column name). Find where results match a regular expression in the third column (column "ID"):

```go 
rowsWithWord, areParamsOk := x.MatchInField("ID", "p([a-z]+)ch$")
```

Use ^ and $ in the regular expression to match exact results.

Please note that `MatchInField`, unlike `MatchInColumn` never includes the header in the search result.

### Boolean functions

Use boolean functions AND, OR and NOT to combine indexes and produce other complex indexes that point to rows in the csv. Very useful to produce search results. 

Indexes are slices of row numbers (integrers between 0 or 1 and the length of the csv -1).

#### OR 2 indexes

```go 
var w []int
w = simplecsv.OrIndex(a,b)
w = simplecsv.OrIndex(a,b,c,d,e)
```

Note: The `OrIndex` function can be used in 2 or more operands.

#### AND 2 indexes

```go 
var p []int
p = simplecsv.AndIndex(a,b)
p = simplecsv.AndIndex(a,b,c,d,e)
```

Note: The `AndIndex` function can be used in 2 or more operands.


#### NOT an index

The code bellow returns the negative of the index `g`, between row 1 and row 4. If `g` is an index with the values `{1, 2}` the negative of `g` is `{3, 4}`. Because 3 and 4 are the integrers between 1 and 4 that are not in `g`.

```go 
var g []int
min := 1
max := 4
p = simplecsv.NotIndex(g, min, max)
```

Note: For csvs with headers the min value is usually 1 and for csvs without headers the min value is usually 0.

### Only 

Only are 2 functions to simplify and sort a csv.

#### Only this rows

It removes rows that are not in the index and reorders a CSV by the index order. If header is true, it starts by the csv header.

```go 
newIndex := []int{1,3}
header := true
x, _ = x.OnlyThisRows(newIndex, header)
```

#### Only this fields

Removes fields that are not in the list of fields, reorders the CSV by the list of fields and adds fields that do not exist as blank fields.

```go 
fieldsList := []string{"Age","ID"}
x, _ = x.OnlyThisFields(fieldsList)
```

## To do

* Function to sort an index, incorporate that function to all methods that return and index.

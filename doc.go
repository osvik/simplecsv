// © 2017 Osvaldo Gago
/*

The Go package simplecsv is a simple mini-library to handle csv files. I'm building it to help me writing small command line scripts. Maybe it's useful to someone else as well.

Some notes:

- all read methods return the value in the csv and a second true/false value that is true if the value exists

- all write methods that change the csv return the changed csv and a true/false value if the operation was sucessful

- all cells are strings

- Simplecsv works with comma separated csv files.

* CSV FILE *

READ

Reads file and parses as a SimpleCsv object. `fileRead` is false if there's an error reading the file or parsing the CSV.

    var x simplecsv.SimpleCsv
    x, fileRead = simplecsv.ReadCsvFile("my1file.csv")

CREATE

Create empty file and define cssv headers:

    var u simplecsv.SimpleCsv
    u = simplecsv.CreateEmpyCsv([]string{"Age", "Gender", "ID"})

WRITE

Write the SimpleCsv object to my2file.csv. If there's an error, `wasWritten` is false.

    wasWritten := u.WriteCsvFile("my2file.csv")

* HEADERS *

The cells of the first row are considered headers.

GET

Get all headers:

    headers := x.GetHeaders()


Get header at position one (second position as it starts from 0):

    headerName, headerExists := x.GetHeader(1)

Get header position: (it returns `-1` if the header does not exist)

    position := x.GetHeaderPosition("Gender")

RENAME

Rename header: (old header, new header)

    x, headerExists := x.RenameHeader("ID", "IDnumber")

`headerExists` is false if the old header does not exist.    

* ROWS *

GET

Get number of rows:

    numberOfRows := x.GetNumberRows()

Get second row:

    row, rowExists := x.GetRow(1)

Get second row as a map:

    row, rowExists := x.GetRowAsMap(1)

ADD

Add a slice to a row. The slice must have the same size as the CSV number of columns. If not wasSucessful is false.

    x, wasSucessful = x.AddRow([]string{"24", "M", "2986732"})

Add row from map: (If the map keys don't exist as columns, the value will be discarded. If a key does not exist, it will create empty cells.)

    mymap := make(map[string]string)
    mymap["Age"] = "62"
    mymap["Gender"] = "F"
    mymap["ID"] = "6463246"
    x, wasAdded = x.AddRowFromMap(mymap)

SET

Set second row (1) from a slice. The length of the slice must be the same as the number of columns and the row must already exist. If there’s an error `wasSet` is false.

    x, wasSet = x.SetRow(1, []string{"45", "F", "8356138"})

Set second row from map: If the map keys don't exist as columns, the value will be discarded. If a key does not exist, it will create empty cells.)

    mymap2 := make(map[string]string)
    mymap2["Age"] = "62"
    mymap2["Gender"] = "F"
    mymap2["ID"] = "6463246"
    x, wasAdded = x.SetRowFromMap(1, mymap2)

DELETE

Delete second row: (If the row number is invalid, `wasDeleted` is false)

    x, wasDeleted = x.DeleteRow(1)

* CELLS *



.
    
*/
package simplecsv

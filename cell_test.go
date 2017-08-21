package simplecsv

import (
	"reflect"
	"testing"
)

func TestSimpleCsv_GetCell(t *testing.T) {
	type args struct {
		column int
		row    int
	}
	tests := []struct {
		name  string
		s     SimpleCsv
		args  args
		want  string
		want1 bool
	}{
		{name: "GetCell1",
			want:  "foo",
			want1: true,
			args:  args{row: 1, column: 1},
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			}},
		{name: "GetCell2",
			want:  "",
			want1: false,
			args:  args{row: 4, column: 1},
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			}},
		{name: "GetCell3",
			want:  "",
			want1: false,
			args:  args{row: 4, column: 0},
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetCell(tt.args.column, tt.args.row)
			if got != tt.want {
				t.Errorf("SimpleCsv.GetCell() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SimpleCsv.GetCell() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSimpleCsv_GetCellByField(t *testing.T) {
	type args struct {
		columnName string
		row        int
	}
	tests := []struct {
		name  string
		s     SimpleCsv
		args  args
		want  string
		want1 bool
	}{
		{name: "GetCellByField1",
			want:  "foo",
			want1: true,
			args:  args{row: 1, columnName: "Ban"},
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			}},
		{name: "GetCellByField2",
			want:  "",
			want1: false,
			args:  args{row: 1, columnName: "Book"},
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			}},
		{name: "GetCellByField3",
			want:  "",
			want1: false,
			args:  args{row: 10, columnName: "Ban"},
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetCellByField(tt.args.columnName, tt.args.row)
			if got != tt.want {
				t.Errorf("SimpleCsv.GetCellByField() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SimpleCsv.GetCellByField() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSimpleCsv_SetCell(t *testing.T) {
	type args struct {
		column int
		row    int
		value  string
	}
	tests := []struct {
		name  string
		s     SimpleCsv
		args  args
		want  SimpleCsv
		want1 bool
	}{
		{name: "SetCell1",
			want: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "bar", "soo"},
			},
			want1: true,
			args:  args{row: 1, column: 1, value: "bar"},
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			}},
		{name: "SetCell2",
			want: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			},
			want1: false,
			args:  args{row: 3, column: 1, value: "bar"},
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			}},
		{name: "SetCell3",
			want: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			},
			want1: false,
			args:  args{row: 1, column: 3, value: "bar"},
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.SetCell(tt.args.column, tt.args.row, tt.args.value)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimpleCsv.SetCell() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SimpleCsv.SetCell() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSimpleCsv_SetCellByField(t *testing.T) {
	type args struct {
		columnName string
		row        int
		value      string
	}
	tests := []struct {
		name  string
		s     SimpleCsv
		args  args
		want  SimpleCsv
		want1 bool
	}{
		{name: "SetCellByField1",
			want: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "bar", "soo"},
			},
			want1: true,
			args:  args{row: 1, columnName: "Ban", value: "bar"},
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			}},
		{name: "SetCellByField2",
			want: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			},
			want1: false,
			args:  args{row: 1, columnName: "Gyu", value: "bar"},
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.SetCellByField(tt.args.columnName, tt.args.row, tt.args.value)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimpleCsv.SetCellByField() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SimpleCsv.SetCellByField() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

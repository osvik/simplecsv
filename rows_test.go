package simplecsv

import (
	"reflect"
	"testing"
)

func TestSimpleCsv_GetNumberRows(t *testing.T) {
	tests := []struct {
		name string
		s    SimpleCsv
		want int
	}{
		{name: "GetNumberRows",
			want: 2,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetNumberRows(); got != tt.want {
				t.Errorf("SimpleCsv.GetNumberRows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimpleCsv_GetRow(t *testing.T) {
	type args struct {
		rowNumber int
	}
	tests := []struct {
		name  string
		s     SimpleCsv
		args  args
		want  []string
		want1 bool
	}{
		{name: "GetRow1",
			want:  []string{"Moo", "foo", "soo"},
			want1: true,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			},
			args: args{
				rowNumber: 1,
			},
		},
		{name: "GetRow2",
			want:  []string{},
			want1: false,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			},
			args: args{
				rowNumber: 3,
			},
		},
		{name: "GetRow3",
			want:  []string{},
			want1: false,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			},
			args: args{
				rowNumber: -1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetRow(tt.args.rowNumber)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimpleCsv.GetRow() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SimpleCsv.GetRow() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSimpleCsv_GetRowAsMap(t *testing.T) {
	type args struct {
		rowNumber int
	}
	tests := []struct {
		name  string
		s     SimpleCsv
		args  args
		want  map[string]string
		want1 bool
	}{
		{name: "GetRowAsMap1",
			want: map[string]string{"Ann": "Moo",
				"Ban":  "foo",
				"Cann": "soo"},
			want1: true,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			},
			args: args{
				rowNumber: 1,
			},
		},
		{name: "GetRowAsMap2",
			want:  map[string]string{"": ""},
			want1: false,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			},
			args: args{
				rowNumber: 10,
			},
		},
		{name: "GetRowAsMap3",
			want:  map[string]string{"": ""},
			want1: false,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			},
			args: args{
				rowNumber: -1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetRowAsMap(tt.args.rowNumber)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimpleCsv.GetRowAsMap() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SimpleCsv.GetRowAsMap() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSimpleCsv_AddRow(t *testing.T) {
	type args struct {
		rowValue []string
	}
	tests := []struct {
		name  string
		s     SimpleCsv
		args  args
		want  SimpleCsv
		want1 bool
	}{
		{name: "AddRow1",
			want: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
				{"Ah", "Eh", "Ih"},
			},
			want1: true,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			},
			args: args{
				rowValue: []string{"Ah", "Eh", "Ih"},
			},
		},
		{name: "AddRow2",
			want: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			},
			want1: false,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			},
			args: args{
				rowValue: []string{"Ah", "Eh", "Ih", "oh"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.AddRow(tt.args.rowValue)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimpleCsv.AddRow() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SimpleCsv.AddRow() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSimpleCsv_AddRowFromMap(t *testing.T) {
	type args struct {
		rowValue map[string]string
	}
	tests := []struct {
		name  string
		s     SimpleCsv
		args  args
		want  SimpleCsv
		want1 bool
	}{
		{name: "AddRowFromMap1",
			want: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
				{"Ah", "Eh", "Ih"},
			},
			want1: true,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			},
			args: args{
				rowValue: map[string]string{"Ann": "Ah",
					"Ban":  "Eh",
					"Cann": "Ih"},
			},
		},
		{name: "AddRowFromMap2",
			want: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
				{"Ah", "", ""},
			},
			want1: true,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			},
			args: args{
				rowValue: map[string]string{"Ann": "Ah"},
			},
		},
		{name: "AddRowFromMap3",
			want: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
				{"Ah", "Eh", "Ih"},
			},
			want1: true,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			},
			args: args{
				rowValue: map[string]string{"Ann": "Ah",
					"Ban":  "Eh",
					"Cann": "Ih",
					"Qwe":  "Sss"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.AddRowFromMap(tt.args.rowValue)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimpleCsv.AddRowFromMap() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SimpleCsv.AddRowFromMap() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSimpleCsv_SetRow(t *testing.T) {
	type args struct {
		rowNumber int
		rowValue  []string
	}
	tests := []struct {
		name  string
		s     SimpleCsv
		args  args
		want  SimpleCsv
		want1 bool
	}{
		{name: "SetRow",
			want: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Ah", "Eh", "Ih"},
			},
			want1: true,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			},
			args: args{
				rowValue:  []string{"Ah", "Eh", "Ih"},
				rowNumber: 1,
			},
		},
		{name: "SetRow2",
			want: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			},
			want1: false,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			},
			args: args{
				rowValue:  []string{"Ah", "Eh", "Ih"},
				rowNumber: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.SetRow(tt.args.rowNumber, tt.args.rowValue)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimpleCsv.SetRow() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SimpleCsv.SetRow() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSimpleCsv_SetRowFromMap(t *testing.T) {
	type args struct {
		rowNumber int
		rowValue  map[string]string
	}
	tests := []struct {
		name  string
		s     SimpleCsv
		args  args
		want  SimpleCsv
		want1 bool
	}{
		{name: "SetRowFromMap",
			want: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Ah", "Eh", "Ih"},
			},
			want1: true,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			},
			args: args{
				rowValue: map[string]string{"Ann": "Ah",
					"Ban":  "Eh",
					"Cann": "Ih"},
				rowNumber: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.SetRowFromMap(tt.args.rowNumber, tt.args.rowValue)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimpleCsv.SetRowFromMap() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SimpleCsv.SetRowFromMap() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSimpleCsv_DeleteRow(t *testing.T) {
	type args struct {
		rowNumber int
	}
	tests := []struct {
		name  string
		s     SimpleCsv
		args  args
		want  SimpleCsv
		want1 bool
	}{
		{name: "DeleteRow1",
			want: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Ah", "Eh", "Ih"},
			},
			want1: true,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
				{"Ah", "Eh", "Ih"},
			},
			args: args{
				rowNumber: 1,
			},
		},
		{name: "DeleteRow2",
			want: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
				{"Ah", "Eh", "Ih"},
			},
			want1: false,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
				{"Ah", "Eh", "Ih"},
			},
			args: args{
				rowNumber: 10,
			},
		},
		{name: "DeleteRow2",
			want: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
				{"Ah", "Eh", "Ih"},
			},
			want1: false,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
				{"Ah", "Eh", "Ih"},
			},
			args: args{
				rowNumber: -1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.DeleteRow(tt.args.rowNumber)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimpleCsv.DeleteRow() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SimpleCsv.DeleteRow() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

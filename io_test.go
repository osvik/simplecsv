package simplecsv

import (
	"reflect"
	"testing"
)

func TestCreateEmpyCsv(t *testing.T) {
	type args struct {
		columnNames []string
	}
	tests := []struct {
		name string
		args args
		want SimpleCsv
	}{
		{name: "CreateEmpyCsv",
			want: SimpleCsv{{"Ann", "Ban", "Cann"}},
			args: args{columnNames: []string{"Ann", "Ban", "Cann"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateEmpyCsv(tt.args.columnNames); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateEmpyCsv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadCsvFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name  string
		args  args
		want  SimpleCsv
		want1 bool
	}{
		{name: "ReadCsvFile1", args: args{filename: "testdata/small.csv"},
			want: SimpleCsv{
				{"Nome", "Email", "Telefono"},
				{"Osvaldo", "osvaldo@test.com", "111"},
				{"Isabel", "isabel@test.com", "222"},
				{"Carlos", "carlos@test.com", "333"},
				{"Marisa", "marisa@test.com", "444"},
			},
			want1: true,
		},
		{name: "ReadCsvFile2", args: args{filename: "doesntexist.csv"},
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := ReadCsvFile(tt.args.filename)
			if got1 != tt.want1 {
				t.Errorf("ReadCsvFile() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSimpleCsv_WriteCsvFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		s    SimpleCsv
		args args
		want bool
	}{
		{name: "WriteCsvFile1",
			args: args{filename: "testdata/small.csv"},
			want: true,
			s: SimpleCsv{
				{"Nome", "Email", "Telefono"},
				{"Osvaldo", "osvaldo@test.com", "111"},
				{"Isabel", "isabel@test.com", "222"},
				{"Carlos", "carlos@test.com", "333"},
				{"Marisa", "marisa@test.com", "444"},
			},
		},
		{name: "WriteCsvFile2",
			args: args{filename: ""},
			want: false,
			s:    SimpleCsv{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.WriteCsvFile(tt.args.filename); got != tt.want {
				t.Errorf("SimpleCsv.WriteCsvFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

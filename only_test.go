package simplecsv

import (
	"reflect"
	"testing"
)

func TestSimpleCsv_OnlyThisRows(t *testing.T) {
	type args struct {
		rowsIndex []int
		header    bool
	}
	tests := []struct {
		name  string
		s     SimpleCsv
		args  args
		want  SimpleCsv
		want1 bool
	}{
		{name: "OnlyThisRows1",
			args: args{
				header:    true,
				rowsIndex: []int{1, 3},
			},
			want1: true,
			want: SimpleCsv{
				{"A", "B", "C"},
				{"A1", "B1", "C1"},
				{"A3", "B3", "C3"},
			},
			s: SimpleCsv{
				{"A", "B", "C"},
				{"A1", "B1", "C1"},
				{"A2", "B2", "C2"},
				{"A3", "B3", "C3"},
			},
		},
		{name: "OnlyThisRows2",
			args: args{
				header:    true,
				rowsIndex: []int{1, 4},
			},
			want1: false,
			want: SimpleCsv{
				{"A", "B", "C"},
				{"A1", "B1", "C1"},
				{"A2", "B2", "C2"},
				{"A3", "B3", "C3"},
			},
			s: SimpleCsv{
				{"A", "B", "C"},
				{"A1", "B1", "C1"},
				{"A2", "B2", "C2"},
				{"A3", "B3", "C3"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.OnlyThisRows(tt.args.rowsIndex, tt.args.header)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimpleCsv.OnlyThisRows() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SimpleCsv.OnlyThisRows() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

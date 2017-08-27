package simplecsv

import (
	"reflect"
	"testing"
)

func TestSimpleCsv_FindInColumn(t *testing.T) {
	type args struct {
		columnPosition int
		word           string
	}
	tests := []struct {
		name  string
		s     SimpleCsv
		args  args
		want  []int
		want1 bool
	}{
		{name: "FindInColumn1",
			want:  []int{2, 3},
			want1: true,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
				{"Net", "bar", "Kap"},
				{"Net", "Bar", "Kap"},
			},
			args: args{
				columnPosition: 1,
				word:           "BAR",
			},
		},
		{name: "FindInColumn2",
			want:  []int{},
			want1: true,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
				{"Net", "bar", "Kap"},
				{"Net", "Bar", "Kap"},
			},
			args: args{
				columnPosition: 1,
				word:           "ZEN",
			},
		},
		{name: "FindInColumn3",
			want:  []int{},
			want1: false,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
				{"Net", "bar", "Kap"},
				{"Net", "Bar", "Kap"},
			},
			args: args{
				columnPosition: 3,
				word:           "BAR",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.FindInColumn(tt.args.columnPosition, tt.args.word)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimpleCsv.FindInColumn() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SimpleCsv.FindInColumn() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

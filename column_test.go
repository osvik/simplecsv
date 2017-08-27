package simplecsv

import (
	"reflect"
	"testing"
)

func TestSimpleCsv_AddEmptyColumn(t *testing.T) {
	type args struct {
		columnName string
	}
	tests := []struct {
		name  string
		s     SimpleCsv
		args  args
		want  SimpleCsv
		want1 bool
	}{
		{name: "AddEmptyColumn",
			want: SimpleCsv{
				{"Ann", "Ban", "Cann", "New"},
				{"Moo", "foo", "soo", ""},
				{"Ah", "Eh", "Ih", ""},
			},
			want1: true,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
				{"Ah", "Eh", "Ih"},
			},
			args: args{
				columnName: "New",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.AddEmptyColumn(tt.args.columnName)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimpleCsv.AddEmptyColumn() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SimpleCsv.AddEmptyColumn() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

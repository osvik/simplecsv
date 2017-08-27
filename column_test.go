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

func TestSimpleCsv_RemoveColumn(t *testing.T) {
	type args struct {
		columnPosition int
	}
	tests := []struct {
		name  string
		s     SimpleCsv
		args  args
		want  SimpleCsv
		want1 bool
	}{
		{name: "RemoveColumn1",
			want: SimpleCsv{
				{"Ann", "Cann"},
				{"Moo", "soo"},
				{"Ah", "Ih"},
			},
			want1: true,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
				{"Ah", "Eh", "Ih"},
			},
			args: args{
				columnPosition: 1,
			},
		},
		{name: "RemoveColumn2",
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
				columnPosition: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RemoveColumn(tt.args.columnPosition)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimpleCsv.RemoveColumn() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SimpleCsv.RemoveColumn() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSimpleCsv_RemoveColumnByName(t *testing.T) {
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
		{name: "RemoveColumnByName1",
			want: SimpleCsv{
				{"Ann", "Cann"},
				{"Moo", "soo"},
				{"Ah", "Ih"},
			},
			want1: true,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
				{"Ah", "Eh", "Ih"},
			},
			args: args{
				columnName: "Ban",
			},
		},
		{name: "RemoveColumnByName2",
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
				columnName: "DontExist",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RemoveColumnByName(tt.args.columnName)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimpleCsv.RemoveColumnByName() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SimpleCsv.RemoveColumnByName() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

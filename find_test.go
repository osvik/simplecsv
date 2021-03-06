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

func TestSimpleCsv_FindInField(t *testing.T) {
	type args struct {
		columnName string
		word       string
	}
	tests := []struct {
		name  string
		s     SimpleCsv
		args  args
		want  []int
		want1 bool
	}{
		{name: "FindInField1",
			want:  []int{2, 3},
			want1: true,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
				{"Net", "bar", "Kap"},
				{"Net", "Bar", "Kap"},
			},
			args: args{
				columnName: "Ban",
				word:       "BAR",
			},
		},
		{name: "FindInField2",
			want:  []int{},
			want1: true,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
				{"Net", "bar", "Kap"},
				{"Net", "Bar", "Kap"},
			},
			args: args{
				columnName: "Ban",
				word:       "ZEN",
			},
		},
		{name: "FindInField3",
			want:  []int{},
			want1: false,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
				{"Net", "bar", "Kap"},
				{"Net", "Bar", "Kap"},
			},
			args: args{
				columnName: "Nek",
				word:       "BAR",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.FindInField(tt.args.columnName, tt.args.word)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimpleCsv.FindInField() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SimpleCsv.FindInField() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSimpleCsv_MatchInColumn(t *testing.T) {
	type args struct {
		columnPosition    int
		regularexpression string
	}
	tests := []struct {
		name  string
		s     SimpleCsv
		args  args
		want  []int
		want1 bool
	}{
		{name: "MatchInColumn1",
			want:  []int{2, 3, 4},
			want1: true,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "peaching", "soo"},
				{"Net", "peach", "Kap"},
				{"Net", "punch", "Kap"},
				{"Net", "unpunch", "Kap"},
			},
			args: args{
				columnPosition:    1,
				regularexpression: "p([a-z]+)ch$",
			},
		},
		{name: "MatchInColumn2",
			want:  []int{},
			want1: true,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
				{"Net", "peach", "Kap"},
				{"Net", "punch", "Kap"},
			},
			args: args{
				columnPosition:    1,
				regularexpression: "z([a-z]+)ch",
			},
		},
		{name: "MatchInColumn3",
			want:  []int{},
			want1: false,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
				{"Net", "bar", "Kap"},
				{"Net", "Bar", "Kap"},
			},
			args: args{
				columnPosition:    3,
				regularexpression: "p([a-z]+)ch$",
			},
		},
		{name: "MatchInColumn4",
			want:  []int{},
			want1: false,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
				{"Net", "bar", "Kap"},
				{"Net", "Bar", "Kap"},
			},
			args: args{
				columnPosition:    1,
				regularexpression: "[",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.MatchInColumn(tt.args.columnPosition, tt.args.regularexpression)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimpleCsv.MatchInColumn() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SimpleCsv.MatchInColumn() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSimpleCsv_MatchInField(t *testing.T) {
	type args struct {
		columnName        string
		regularexpression string
	}
	tests := []struct {
		name  string
		s     SimpleCsv
		args  args
		want  []int
		want1 bool
	}{
		{name: "MatchInField1",
			want:  []int{2, 3, 4},
			want1: true,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "peaching", "soo"},
				{"Net", "peach", "Kap"},
				{"Net", "punch", "Kap"},
				{"Net", "unpunch", "Kap"},
			},
			args: args{
				columnName:        "Ban",
				regularexpression: "p([a-z]+)ch$",
			},
		},
		{name: "MatchInField2",
			want:  []int{},
			want1: true,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
				{"Net", "peach", "Kap"},
				{"Net", "punch", "Kap"},
			},
			args: args{
				columnName:        "Ban",
				regularexpression: "z([a-z]+)ch",
			},
		},
		{name: "MatchInField3",
			want:  []int{},
			want1: false,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
				{"Net", "bar", "Kap"},
				{"Net", "Bar", "Kap"},
			},
			args: args{
				columnName:        "Nez",
				regularexpression: "p([a-z]+)ch$",
			},
		},
		{name: "MatchInField4",
			want:  []int{},
			want1: false,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
				{"Net", "bar", "Kap"},
				{"Net", "Bar", "Kap"},
			},
			args: args{
				columnName:        "Ban",
				regularexpression: "[",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.MatchInField(tt.args.columnName, tt.args.regularexpression)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimpleCsv.MatchInField() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SimpleCsv.MatchInField() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

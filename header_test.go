package simplecsv

import (
	"reflect"
	"testing"
)

func TestSimpleCsv_GetHeaders(t *testing.T) {
	tests := []struct {
		name string
		s    SimpleCsv
		want []string
	}{
		{name: "GetHeaders",
			want: []string{"Ann", "Ban", "Cann"},
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetHeaders(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimpleCsv.GetHeaders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimpleCsv_GetHeader(t *testing.T) {
	type args struct {
		columnPosition int
	}
	tests := []struct {
		name  string
		s     SimpleCsv
		args  args
		want  string
		want1 bool
	}{
		{name: "GetHeader1",
			want:  "Ban",
			want1: true,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			},
			args: args{
				columnPosition: 1,
			},
		},
		{name: "GetHeader2",
			want:  "",
			want1: false,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			},
			args: args{
				columnPosition: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetHeader(tt.args.columnPosition)
			if got != tt.want {
				t.Errorf("SimpleCsv.GetHeader() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SimpleCsv.GetHeader() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSimpleCsv_GetHeaderPosition(t *testing.T) {
	type args struct {
		columnName string
	}
	tests := []struct {
		name string
		s    SimpleCsv
		args args
		want int
	}{
		{name: "GetHeaderPosition1",
			want: 1,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			},
			args: args{
				columnName: "Ban",
			},
		},
		{name: "GetHeaderPosition2",
			want: -1,
			s: SimpleCsv{
				{"Ann", "Ban", "Cann"},
				{"Moo", "foo", "soo"},
			},
			args: args{
				columnName: "Moo",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetHeaderPosition(tt.args.columnName); got != tt.want {
				t.Errorf("SimpleCsv.GetHeaderPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimpleCsv_RenameHeader(t *testing.T) {
	type args struct {
		oldHeader string
		newHeader string
	}
	tests := []struct {
		name  string
		s     SimpleCsv
		args  args
		want  SimpleCsv
		want1 bool
	}{
		{name: "RenameHeader1", want: SimpleCsv{
			{"Ann", "Xoo", "Cann"},
			{"Moo", "foo", "soo"},
		}, want1: true, s: SimpleCsv{
			{"Ann", "Ban", "Cann"},
			{"Moo", "foo", "soo"},
		}, args: args{
			oldHeader: "Ban",
			newHeader: "Xoo",
		}},
		{name: "RenameHeader2", want: SimpleCsv{
			{"Ann", "Ban", "Cann"},
			{"Moo", "foo", "soo"},
		}, want1: false, s: SimpleCsv{
			{"Ann", "Ban", "Cann"},
			{"Moo", "foo", "soo"},
		}, args: args{
			oldHeader: "KKK",
			newHeader: "Xoo",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.RenameHeader(tt.args.oldHeader, tt.args.newHeader)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimpleCsv.RenameHeader() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SimpleCsv.RenameHeader() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

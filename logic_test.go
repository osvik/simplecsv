package simplecsv

import (
	"reflect"
	"testing"
)

// compareUnorderedSlices compares if 2 slices contain the same elements in any order
func compareUnorderedSlices(x, y interface{}) bool {
	nx := x.([]int)
	ny := y.([]int)

	for _, v := range nx {
		if !contains(ny, v) {
			return false
		}
	}

	for _, k := range ny {
		if !contains(nx, k) {
			return false
		}
	}
	return true
}

func Test_countIndexesInSlices(t *testing.T) {
	type args struct {
		indexes [][]int
	}
	tests := []struct {
		name string
		args args
		want map[int]int
	}{
		{name: "A",
			want: map[int]int{1: 1, 2: 1, 3: 1, 4: 1, 5: 1, 6: 1, 7: 1, 8: 1},
			args: args{indexes: [][]int{
				{2, 4, 6, 8},
				{1, 3, 5, 7},
			}},
		},
		{name: "B",
			want: map[int]int{1: 2, 2: 1, 3: 1, 4: 1, 5: 1, 6: 1, 7: 1, 8: 3},
			args: args{indexes: [][]int{
				{2, 4, 6, 8},
				{1, 3, 5, 7, 8},
				{1, 8},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countIndexesInSlices(tt.args.indexes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countIndexesInSlices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_showsMoreThanTimes(t *testing.T) {
	type args struct {
		valuesMap map[int]int
		min       int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "ShowMoreThanTimes1",
			want: []int{2, 4},
			args: args{
				valuesMap: map[int]int{1: 1, 2: 2, 3: 1, 4: 3},
				min:       2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := showsMoreThanTimes(tt.args.valuesMap, tt.args.min); !compareUnorderedSlices(got, tt.want) {
				t.Errorf("showsMoreThanTimes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrIndex(t *testing.T) {
	type args struct {
		indexes [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "OrIndex1",
			want: []int{2, 4, 6, 1, 3, 5},
			args: args{
				indexes: [][]int{
					{2, 4, 6},
					{1, 3, 5},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OrIndex(tt.args.indexes...); !compareUnorderedSlices(got, tt.want) {
				t.Errorf("OrIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAndIndex(t *testing.T) {
	type args struct {
		indexes [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "OrIndex1",
			want: []int{6},
			args: args{
				indexes: [][]int{
					{2, 4, 6},
					{1, 6, 5},
					{1, 6, 3},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AndIndex(tt.args.indexes...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AndIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_contains(t *testing.T) {
	type args struct {
		s []int
		e int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Contains1",
			want: true,
			args: args{
				s: []int{1, 2, 3, 4},
				e: 4,
			},
		},
		{name: "Contains2",
			want: false,
			args: args{
				s: []int{1, 2, 3, 4},
				e: 9,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotIndex(t *testing.T) {
	type args struct {
		index []int
		min   int
		max   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "NotInIndex1",
			want: []int{2, 4, 6},
			args: args{
				index: []int{1, 3, 5, 7},
				min:   1,
				max:   7,
			},
		},
		{name: "NotInIndex2",
			want: []int{2, 4, 6, 0},
			args: args{
				index: []int{1, 3, 5, 7},
				min:   0,
				max:   7,
			},
		},
		{name: "NotInIndex3",
			want: []int{},
			args: args{
				index: []int{1, 2, 3},
				min:   1,
				max:   3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NotIndex(tt.args.index, tt.args.min, tt.args.max); !compareUnorderedSlices(got, tt.want) {
				t.Errorf("NotIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

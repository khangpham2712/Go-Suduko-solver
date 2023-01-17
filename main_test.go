package main

import (
	"reflect"
	"testing"
)

func TestSolveSudoku(t *testing.T) {
	input := [][]int{
		{7, 8, 0, 4, 0, 0, 1, 2, 0},
		{6, 0, 0, 0, 7, 5, 0, 0, 9},
		{0, 0, 0, 6, 0, 1, 0, 7, 8},
		{0, 0, 7, 0, 4, 0, 2, 6, 0},
		{0, 0, 1, 0, 5, 0, 9, 3, 0},
		{9, 0, 4, 0, 6, 0, 0, 0, 5},
		{0, 7, 0, 3, 0, 0, 0, 1, 2},
		{1, 2, 0, 0, 0, 7, 4, 0, 0},
		{0, 4, 9, 2, 0, 6, 0, 0, 7},
	}
	expected := [][]int{
		{7, 8, 5, 4, 3, 9, 1, 2, 6},
		{6, 1, 2, 8, 7, 5, 3, 4, 9},
		{4, 9, 3, 6, 2, 1, 5, 7, 8},
		{8, 5, 7, 9, 4, 3, 2, 6, 1},
		{2, 6, 1, 7, 5, 8, 9, 3, 4},
		{9, 3, 4, 1, 6, 2, 7, 8, 5},
		{5, 7, 8, 3, 9, 4, 6, 1, 2},
		{1, 2, 6, 5, 8, 7, 4, 9, 3},
		{3, 4, 9, 2, 1, 6, 8, 5, 7},
	}
	type args struct {
		board [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{name: "Test 1",
			args: args{
				board: input,
			},
			want: expected},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveSudoku(tt.args.board); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SolveSudoku() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "fmt"

type Point struct {
	X int
	Y int
}

func IsValidNum(board [][]int, i int, j int, num int) bool {
	for k := 0; k < 9; k++ {
		if (k != j && board[i][k] == board[i][j]) || (k != i && board[k][j] == board[i][j]) {
			return false
		}
	}
	row, col := i/3, j/3
	for m := row * 3; m <= row*3+2; m++ {
		for n := col * 3; n <= col*3+2; n++ {
			if (m != i || n != j) && board[m][n] == board[i][j] {
				return false
			}
		}
	}
	return true
}

func SolveSudoku(board [][]int) [][]int {
	// Write your code here.
	points := []Point{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				points = append(points, Point{X: i, Y: j})
			}
		}
	}
	done := false
	Traverse(board, 0, len(points), points, &done)
	return board
}

func Traverse(board [][]int, i int, length int, points []Point, done *bool) {
	for num := 1; num <= 9; num++ {
		board[points[i].X][points[i].Y] = num
		if IsValidNum(board, points[i].X, points[i].Y, num) {
			if i == length-1 {
				*done = true
				return
			}
			Traverse(board, i+1, length, points, done)
			if *done {
				return
			}
		}
		board[points[i].X][points[i].Y] = 0
	}
}

func main() {
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
	actual := SolveSudoku(input)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if expected[i][j] != actual[i][j] {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")
}

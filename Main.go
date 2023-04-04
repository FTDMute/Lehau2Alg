package main

import "fmt"

type Solution struct {
	N              int
	M              int
	printFlag      bool
	findOne        bool
	K              int
	SolutionsCount int
	Solutions      [][][]int
}

func NewSolution(N, M int, printFlag, findOne bool) Solution {
	var sol Solution = Solution{N: N, M: M, printFlag: printFlag, findOne: findOne, K: 0, SolutionsCount: 0, Solutions: make([][][]int, 0)}
	return sol
}

func PrintBoard(board *[][]int) {
	fmt.Println()
	for _, row := range *board {
		for _, el := range row {
			fmt.Print(el, "\t")
		}
		fmt.Println()
	}
}

func (sol *Solution) IsPossible(x, y int, board *[][]int) bool {
	if x+1 < sol.M && y < sol.N && (*board)[y][x+1] == 1 {
		return false
	}
	if x+1 < sol.M && y+1 < sol.N && (*board)[y+1][x+1] == 1 {
		return false
	}
	if x < sol.M && y+1 < sol.N && (*board)[y+1][x] == 1 {
		return false
	}
	if x-1 > -1 && y-1 > -1 && (*board)[y-1][x-1] == 1 {
		return false
	}
	if x < sol.M && y-1 > -1 && (*board)[y-1][x] == 1 {
		return false
	}
	if x+1 < sol.M && y-1 > -1 && (*board)[y-1][x+1] == 1 {
		return false
	}
	if x-1 > -1 && y < sol.N && (*board)[y][x-1] == 1 {
		return false
	}
	if x-1 > -1 && y+1 < sol.N && (*board)[y+1][x-1] == 1 {
		return false
	}
	return true
}

func (sol *Solution) FindSolutions(board [][]int, posed, xtoc, ytoc int) {
	if xtoc == sol.M && ytoc == sol.N-1 {
		sol.SolutionsCount += 1
		sol.Solutions = append(sol.Solutions, board)
		if (*sol).printFlag {
			PrintBoard(&board)
		}
		return
	}
	for y := ytoc; y < (*sol).N; y++ {
		for x := xtoc; x < (*sol).M; x++ {
			if board[y][x] == 0 && (*sol).IsPossible(x, y, &board) {
				board[y][x] = 1
				if posed+1 > (*sol).K {
					sol.K = posed + 1
				}
				if (*sol).findOne {
					return
				}
				(*sol).FindSolutions(board, posed+1, x+1, y)
				board[y][x] = 0
			}
		}
		xtoc = 0
	}
}

func main() {
	var n, m int
	fmt.Print("Введите размеры доски (X Y): ")
	fmt.Scan(&m, &n)
	solution := NewSolution(n, m, true, false)
	board := make([][]int, 0)
	for i := 0; i < n; i++ {
		board = append(board, make([]int, m))
		for j := 0; j < m; j++ {
			board[i][j] = 0
		}
	}
	solution.FindSolutions(board, 0, 0, 0)
	fmt.Println("K:", solution.K, "SolutionsCount:", solution.SolutionsCount)
}

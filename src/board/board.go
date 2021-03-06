package board

import (
	"fmt"
)

type Board struct {
	BoardSize BoardSize
	Board     [][]bool
}

type BoardSize struct {
	Rows    int `json:"rows"`
	Columns int `json:"columns"`
}

func (b *Board) Reset() {
	for i := 0; i < b.BoardSize.Columns; i++ {
		for j := 0; j < b.BoardSize.Rows; j++ {
			b.Board[i][j] = false
		}
	}
}

func (b Board) Print() {

	for i := 0; i < b.BoardSize.Columns; i++ {
		for j := 0; j < b.BoardSize.Rows; j++ {
			if b.Board[i][j] == true {
				fmt.Printf("\033[1;31m#\033[0m")
			} else {
				fmt.Printf("\033[1;32m#\033[0m")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func CreateEmptyBoard(colLen, rowLen int) [][]bool {
	board := make([][]bool, colLen)
	rows := make([]bool, colLen*rowLen)

	for i := 0; i < colLen; i++ {
		board[i] = rows[i*rowLen : (i+1)*rowLen]
	}
	for i := 0; i < colLen; i++ {
		for j := 0; j < rowLen; j++ {
			board[i][j] = false
		}
	}
	return board
}

func (b Board) IsOutofBounds(col, row int) (bool, error) {
	if col < 0 || row < 0 || col > b.BoardSize.Columns-1 || row > b.BoardSize.Rows-1 {
		return true, fmt.Errorf("index out of range [%d][%d] of [%d][%d]", col, row, b.BoardSize.Columns-1, b.BoardSize.Rows-1)
	}
	return false, nil
}

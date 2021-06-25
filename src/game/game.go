package game

import (
	"encoding/csv"
	"fmt"
	"gameoflife/board"
	"log"
	"os"
	"strconv"
	"time"
)

type Game struct {
	Board             board.Board
	neighboursOptions []Option
	ChangedCells      []CellData
}

type CellData struct {
	X     int  `json:"x"`
	Y     int  `json:"y"`
	State bool `json:"state"`
}

type Option struct {
	Column int `json:"columns"`
	Row    int `json:"rows"`
}

func NewGame(colLen, rowLen int) *Game {
	game := Game{}
	game.InitGame(colLen, rowLen)
	return &game
}

func (g *Game) InitGame(colLen, rowLen int) {

	g.Board.BoardSize.Columns = colLen
	g.Board.BoardSize.Rows = rowLen
	g.Board.Board = board.CreateEmptyBoard(g.Board.BoardSize.Columns, g.Board.BoardSize.Rows)

	g.neighboursOptions = []Option{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
		{-1, -1},
		{1, 1},
		{1, -1},
		{-1, 1},
	}
}

func (g *Game) InitGameFromCSV(filePath string) {
	records, err := readCsv(filePath)
	if err != nil {
		log.Fatal(err)
	}
	g.InitGame(len(records), len(records[0]))

	for i := 0; i < g.Board.BoardSize.Columns; i++ {
		for j := 0; j < g.Board.BoardSize.Rows; j++ {
			g.Board.Board[i][j], err = strconv.ParseBool(records[i][j])
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
		}
	}
}

func readCsv(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return [][]string{}, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return records, nil
}

func (g Game) PrintBoard() {
	g.Board.Print()
}

func (g *Game) Set(col int, row int, value bool) error {
	if outOfBounds, err := g.Board.IsOutofBounds(col, row); !outOfBounds {
		g.Board.Board[col][row] = value
		return nil
	} else {
		return err
	}
}

func (g Game) Run() {
	for {
		g.PrintBoard()
		g.Tick()
		time.Sleep(600 * time.Millisecond)
	}
}

func (g *Game) Tick() {
	tempBoard := board.CreateEmptyBoard(g.Board.BoardSize.Columns, g.Board.BoardSize.Rows)
	g.ChangedCells = make([]CellData, 0)

	for i := 0; i < g.Board.BoardSize.Columns; i++ {
		for j := 0; j < g.Board.BoardSize.Rows; j++ {
			if g.isAlive(i, j) {
				tempBoard[i][j] = true
				if tempBoard[i][j] != g.Board.Board[i][j] {
					g.ChangedCells = append(g.ChangedCells, CellData{
						X:     i,
						Y:     j,
						State: true,
					})
				}
			} else {
				tempBoard[i][j] = false
				if tempBoard[i][j] != g.Board.Board[i][j] {
					g.ChangedCells = append(g.ChangedCells, CellData{
						X:     i,
						Y:     j,
						State: false,
					})
				}
			}
		}

	}
	g.Board.Board = tempBoard
}

func (g Game) isAlive(col, row int) bool {
	liveNeighbourCount := 0
	for _, option := range g.neighboursOptions {
		if outOfBounds, _ := g.Board.IsOutofBounds(col+option.Column, row+option.Row); !outOfBounds {
			if g.Board.Board[col+option.Column][row+option.Row] == true {
				liveNeighbourCount++
			}
		}
	}
	if liveNeighbourCount == 3 || (g.Board.Board[col][row] == true && liveNeighbourCount == 2) {
		return true
	} else {
		return false
	}
}

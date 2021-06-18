package game

import (
	"../board"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

const alive = 1
const dead = 0

type Game struct {
	Board             board.Board
	neighboursOptions []neighboursOption
	ChangedCells      []CellData
}

type CellData struct {
	X     int `json:"x"`
	Y     int `json:"y"`
	State int `json:"state"`
}

type neighboursOption struct {
	columnDif int
	rowDif    int
}

func (g *Game) InitGame(colLen, rowLen int) {

	g.Board.BoardSize.Columns = colLen
	g.Board.BoardSize.Rows = rowLen
	g.Board.Board = board.CreateEmptyBoard(g.Board.BoardSize.Columns, g.Board.BoardSize.Rows)

	g.neighboursOptions = []neighboursOption{
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
			g.Board.Board[i][j], err = strconv.Atoi(records[i][j])
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

func (g *Game) Set(col, row, value int) {

	if outOfBounds, err := g.Board.IsOutofBounds(col, row); !outOfBounds {
		g.Board.Board[col][row] = value
	} else {
		log.Fatal(err)
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
				tempBoard[i][j] = alive
				if tempBoard[i][j] != g.Board.Board[i][j] {
					g.ChangedCells = append(g.ChangedCells, CellData{
						X:     i,
						Y:     j,
						State: alive,
					})
				}
			} else {
				tempBoard[i][j] = dead
				if tempBoard[i][j] != g.Board.Board[i][j] {
					g.ChangedCells = append(g.ChangedCells, CellData{
						X:     i,
						Y:     j,
						State: dead,
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
		if outOfBounds, _ := g.Board.IsOutofBounds(col+option.columnDif, row+option.rowDif); !outOfBounds {
			if g.Board.Board[col+option.columnDif][row+option.rowDif] == alive {
				liveNeighbourCount++
			}
		}
	}
	if liveNeighbourCount == 3 || (g.Board.Board[col][row] == alive && liveNeighbourCount == 2) {
		return true
	} else {
		return false
	}
}

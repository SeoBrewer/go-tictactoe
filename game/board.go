package game

import "fmt"

type Board struct {
	cells [3][3]string
}

func NewBoard() *Board {
	board := &Board{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board.cells[i][j] = " "
		}
	}
	return board
}

func (b *Board) Print() {
	fmt.Println("\n  0   1   2")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", i)
		for j := 0; j < 3; j++ {
			fmt.Printf(" %s ", b.cells[i][j])
			if j < 2 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < 2 {
			fmt.Println("  ---|---|---")
		}
	}
	fmt.Println()
}

func (b *Board) MakeMove(row, col int, player string) bool {
	if row < 0 || row > 2 || col < 0 || col > 2 || b.cells[row][col] != " " {
		return false
	}
	b.cells[row][col] = player
	return true
}

func (b *Board) CheckWinner() string {
	// Verificar filas
	for i := 0; i < 3; i++ {
		if b.cells[i][0] != " " && b.cells[i][0] == b.cells[i][1] && b.cells[i][1] == b.cells[i][2] {
			return b.cells[i][0]
		}
	}

	// Verificar columnas
	for j := 0; j < 3; j++ {
		if b.cells[0][j] != " " && b.cells[0][j] == b.cells[1][j] && b.cells[1][j] == b.cells[2][j] {
			return b.cells[0][j]
		}
	}

	// Verificar diagonales
	if b.cells[0][0] != " " && b.cells[0][0] == b.cells[1][1] && b.cells[1][1] == b.cells[2][2] {
		return b.cells[0][0]
	}

	if b.cells[0][2] != " " && b.cells[0][2] == b.cells[1][1] && b.cells[1][1] == b.cells[2][0] {
		return b.cells[0][2]
	}

	return ""
}

func (b *Board) IsFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.cells[i][j] == " " {
				return false
			}
		}
	}
	return true
}

func (b *Board) GetEmptySpots() [][2]int {
	var emptySpots [][2]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.cells[i][j] == " " {
				emptySpots = append(emptySpots, [2]int{i, j})
			}
		}
	}
	return emptySpots
}

func (b *Board) GetCell(row, col int) string {
	return b.cells[row][col]
}

func (b *Board) SetCell(row, col int, value string) {
	b.cells[row][col] = value
}
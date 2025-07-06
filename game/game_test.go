package game

import "testing"

func TestNewBoard(t *testing.T) {
	board := NewBoard()

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board.cells[i][j] != " " {
				t.Errorf("Expected empty cell at [%d][%d], got %s", i, j, board.cells[i][j])
			}
		}
	}
}

func TestMakeMove(t *testing.T) {
	board := NewBoard()

	// Movimiento válido
	if !board.MakeMove(1, 1, "X") {
		t.Error("Expected valid move to succeed")
	}

	if board.GetCell(1, 1) != "X" {
		t.Error("Expected X at position [1][1]")
	}

	// Movimiento inválido (posición ocupada)
	if board.MakeMove(1, 1, "O") {
		t.Error("Expected invalid move to fail")
	}

	// Movimiento inválido (fuera de rango)
	if board.MakeMove(5, 5, "X") {
		t.Error("Expected out of bounds move to fail")
	}
}

func TestCheckWinnerRow(t *testing.T) {
	board := NewBoard()

	// Test fila ganadora
	board.MakeMove(0, 0, "X")
	board.MakeMove(0, 1, "X")
	board.MakeMove(0, 2, "X")

	winner := board.CheckWinner()
	if winner != "X" {
		t.Errorf("Expected X to win, got %s", winner)
	}
}

func TestCheckWinnerColumn(t *testing.T) {
	board := NewBoard()

	// Test columna ganadora
	board.MakeMove(0, 0, "O")
	board.MakeMove(1, 0, "O")
	board.MakeMove(2, 0, "O")

	winner := board.CheckWinner()
	if winner != "O" {
		t.Errorf("Expected O to win, got %s", winner)
	}
}

func TestCheckWinnerDiagonal(t *testing.T) {
	board := NewBoard()

	// Test diagonal ganadora
	board.MakeMove(0, 0, "X")
	board.MakeMove(1, 1, "X")
	board.MakeMove(2, 2, "X")

	winner := board.CheckWinner()
	if winner != "X" {
		t.Errorf("Expected X to win, got %s", winner)
	}
}

func TestCheckWinnerAntiDiagonal(t *testing.T) {
	board := NewBoard()

	// Test diagonal inversa ganadora
	board.MakeMove(0, 2, "O")
	board.MakeMove(1, 1, "O")
	board.MakeMove(2, 0, "O")

	winner := board.CheckWinner()
	if winner != "O" {
		t.Errorf("Expected O to win, got %s", winner)
	}
}

func TestIsFull(t *testing.T) {
	board := NewBoard()

	// Tablero vacío
	if board.IsFull() {
		t.Error("Expected empty board to not be full")
	}

	// Llenar tablero
	symbols := []string{"X", "O", "X", "O", "X", "O", "X", "O", "X"}
	index := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board.MakeMove(i, j, symbols[index])
			index++
		}
	}

	if !board.IsFull() {
		t.Error("Expected full board to be full")
	}
}

func TestGetEmptySpots(t *testing.T) {
	board := NewBoard()

	// Tablero vacío debería tener 9 espacios
	emptySpots := board.GetEmptySpots()
	if len(emptySpots) != 9 {
		t.Errorf("Expected 9 empty spots, got %d", len(emptySpots))
	}

	// Hacer un movimiento
	board.MakeMove(1, 1, "X")
	emptySpots = board.GetEmptySpots()
	if len(emptySpots) != 8 {
		t.Errorf("Expected 8 empty spots after one move, got %d", len(emptySpots))
	}
}

func TestAIFindWinningMove(t *testing.T) {
	board := NewBoard()
	ai := NewAI("O")

	// Configurar tablero para que O pueda ganar
	board.MakeMove(0, 0, "O")
	board.MakeMove(0, 1, "O")
	// board.MakeMove(0, 2, "O") sería la jugada ganadora

	row, col, canWin := ai.findWinningMove(board, "O")
	if !canWin {
		t.Error("Expected AI to find winning move")
	}

	if row != 0 || col != 2 {
		t.Errorf("Expected winning move at (0,2), got (%d,%d)", row, col)
	}
}
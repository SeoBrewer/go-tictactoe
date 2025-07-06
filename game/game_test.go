package game

import "testing"

func TestNewBoard(t *testing.T) {
	board := NewBoard()

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board.GetCell(i, j) != " " {
				t.Errorf("Expected empty cell at [%d][%d], got %s", i, j, board.GetCell(i, j))
			}
		}
	}
}

func TestMakeMove(t *testing.T) {
	board := NewBoard()

	// Valid move
	if !board.MakeMove(1, 1, "X") {
		t.Error("Expected valid move to succeed")
	}

	if board.GetCell(1, 1) != "X" {
		t.Error("Expected X at position [1][1]")
	}

	// Invalid move (position occupied)
	if board.MakeMove(1, 1, "O") {
		t.Error("Expected invalid move to fail")
	}

	// Invalid move (out of bounds)
	if board.MakeMove(5, 5, "X") {
		t.Error("Expected out of bounds move to fail")
	}

	if board.MakeMove(-1, -1, "X") {
		t.Error("Expected negative coordinates to fail")
	}
}

func TestSetCell(t *testing.T) {
	board := NewBoard()

	// Test SetCell
	board.SetCell(2, 2, "X")
	if board.GetCell(2, 2) != "X" {
		t.Error("Expected SetCell to work")
	}

	// Overwrite with SetCell
	board.SetCell(2, 2, "O")
	if board.GetCell(2, 2) != "O" {
		t.Error("Expected SetCell to overwrite")
	}

	// SetCell with empty space
	board.SetCell(2, 2, " ")
	if board.GetCell(2, 2) != " " {
		t.Error("Expected SetCell to clear cell")
	}
}

func TestCheckWinnerRow(t *testing.T) {
	board := NewBoard()

	// Test winning row
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

	// Test winning column
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

	// Test winning diagonal
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

	// Test winning anti-diagonal
	board.MakeMove(0, 2, "O")
	board.MakeMove(1, 1, "O")
	board.MakeMove(2, 0, "O")

	winner := board.CheckWinner()
	if winner != "O" {
		t.Errorf("Expected O to win, got %s", winner)
	}
}

func TestNoWinner(t *testing.T) {
	board := NewBoard()

	// Empty board
	winner := board.CheckWinner()
	if winner != "" {
		t.Errorf("Expected no winner on empty board, got %s", winner)
	}

	// Some moves, no winner yet
	board.MakeMove(0, 0, "X")
	board.MakeMove(1, 1, "O")
	board.MakeMove(2, 2, "X")

	winner = board.CheckWinner()
	if winner != "" {
		t.Errorf("Expected no winner yet, got %s", winner)
	}
}

func TestIsFull(t *testing.T) {
	board := NewBoard()

	// Empty board
	if board.IsFull() {
		t.Error("Expected empty board to not be full")
	}

	// Fill board with no winner (tie)
	board.MakeMove(0, 0, "X") // X
	board.MakeMove(0, 1, "O") // O
	board.MakeMove(0, 2, "X") // X
	board.MakeMove(1, 0, "O") // O
	board.MakeMove(1, 1, "X") // X
	board.MakeMove(1, 2, "O") // O
	board.MakeMove(2, 0, "O") // O
	board.MakeMove(2, 1, "X") // X
	board.MakeMove(2, 2, "O") // O

	if !board.IsFull() {
		t.Error("Expected full board to be full")
	}

	// Check that there is no winner (real tie)
	winner := board.CheckWinner()
	if winner != "" {
		t.Errorf("Expected no winner in tie game, got %s", winner)
	}
}

func TestGetEmptySpots(t *testing.T) {
	board := NewBoard()

	// Empty board should have 9 spots
	emptySpots := board.GetEmptySpots()
	if len(emptySpots) != 9 {
		t.Errorf("Expected 9 empty spots, got %d", len(emptySpots))
	}

	// Make a move
	board.MakeMove(1, 1, "X")
	emptySpots = board.GetEmptySpots()
	if len(emptySpots) != 8 {
		t.Errorf("Expected 8 empty spots after one move, got %d", len(emptySpots))
	}

	// Check that coordinates are correct
	found := false
	for _, spot := range emptySpots {
		if spot[0] == 1 && spot[1] == 1 {
			found = true
			break
		}
	}
	if found {
		t.Error("Expected occupied spot to not be in empty spots list")
	}
}

func TestAIFindWinningMove(t *testing.T) {
	board := NewBoard()
	ai := NewAI("O")

	// Set up board so O can win in a row
	board.MakeMove(0, 0, "O")
	board.MakeMove(0, 1, "O")
	// board.MakeMove(0, 2, "O") would be the winning move

	row, col, canWin := ai.findWinningMove(board, "O")
	if !canWin {
		t.Error("Expected AI to find winning move")
	}

	if row != 0 || col != 2 {
		t.Errorf("Expected winning move at (0,2), got (%d,%d)", row, col)
	}

	// Check that the board did not change after simulation
	if board.GetCell(0, 2) != " " {
		t.Error("Expected board to be unchanged after findWinningMove")
	}
}

func TestAIFindBlockingMove(t *testing.T) {
	board := NewBoard()
	ai := NewAI("O")

	// Set up board so X can win (O must block)
	board.MakeMove(1, 0, "X")
	board.MakeMove(1, 1, "X")
	// X can win at (1,2), O must block

	row, col, mustBlock := ai.findWinningMove(board, "X")
	if !mustBlock {
		t.Error("Expected AI to find blocking move")
	}

	if row != 1 || col != 2 {
		t.Errorf("Expected blocking move at (1,2), got (%d,%d)", row, col)
	}
}

func TestPlayerCreation(t *testing.T) {
	player := NewPlayer("X", "TestPlayer")

	if player.GetSymbol() != "X" {
		t.Errorf("Expected symbol X, got %s", player.GetSymbol())
	}

	if player.GetName() != "TestPlayer" {
		t.Errorf("Expected name TestPlayer, got %s", player.GetName())
	}
}

func TestAICreation(t *testing.T) {
	ai := NewAI("O")

	if ai.symbol != "O" {
		t.Errorf("Expected AI symbol O, got %s", ai.symbol)
	}
}
package game

import (
	"fmt"
	"math/rand"
	"time"
)

type AI struct {
	symbol string
}

func NewAI(symbol string) *AI {
	return &AI{symbol: symbol}
}

func (ai *AI) GetMove(board *Board) (int, int) {
	fmt.Println("🧠 AI is thinking strategically...")
	time.Sleep(1 * time.Second)

	// 1. Can the AI win?
	if row, col, canWin := ai.findWinningMove(board, ai.symbol); canWin {
		fmt.Printf("🎯 AI goes for the win: %d,%d\n", row, col)
		return row, col
	}

	// 2. Does it need to block the player?
	playerSymbol := "X"
	if ai.symbol == "X" {
		playerSymbol = "O"
	}
	if row, col, mustBlock := ai.findWinningMove(board, playerSymbol); mustBlock {
		fmt.Printf("🛡️ AI blocks your move: %d,%d\n", row, col)
		return row, col
	}

	// 3. Take the center if it's free
	if board.GetCell(1, 1) == " " {
		fmt.Printf("🎯 AI takes the center: 1,1\n")
		return 1, 1
	}

	// 4. Random move
	emptySpots := board.GetEmptySpots()
	if len(emptySpots) == 0 {
		// No moves available (should not happen)
		fmt.Println("🤖 AI: No moves available")
		return -1, -1
	}
	
	randomIndex := rand.Intn(len(emptySpots))
	chosen := emptySpots[randomIndex]
	fmt.Printf("🤖 AI chooses: %d,%d\n", chosen[0], chosen[1])
	return chosen[0], chosen[1]
}

func (ai *AI) findWinningMove(board *Board, player string) (int, int, bool) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board.GetCell(i, j) == " " {
				// Try this move
				board.SetCell(i, j, player)
				if board.CheckWinner() == player {
					// Undo the move
					board.SetCell(i, j, " ")
					return i, j, true
				}
				// Undo the move
				board.SetCell(i, j, " ")
			}
		}
	}
	return 0, 0, false
}
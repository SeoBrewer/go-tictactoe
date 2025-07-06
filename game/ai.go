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
	fmt.Println("🧠 IA está pensando estratégicamente...")
	time.Sleep(1 * time.Second)

	// 1. ¿Puede ganar la IA?
	if row, col, canWin := ai.findWinningMove(board, ai.symbol); canWin {
		fmt.Printf("🎯 IA va por la victoria: %d,%d\n", row, col)
		return row, col
	}

	// 2. ¿Necesita bloquear al jugador?
	playerSymbol := "X"
	if ai.symbol == "X" {
		playerSymbol = "O"
	}
	if row, col, mustBlock := ai.findWinningMove(board, playerSymbol); mustBlock {
		fmt.Printf("🛡️ IA bloquea tu jugada: %d,%d\n", row, col)
		return row, col
	}

	// 3. Tomar el centro si está libre
	if board.GetCell(1, 1) == " " {
		fmt.Printf("🎯 IA toma el centro: 1,1\n")
		return 1, 1
	}

	// 4. Movimiento aleatorio
	emptySpots := board.GetEmptySpots()
	if len(emptySpots) == 0 {
		// No hay movimientos disponibles (no debería pasar)
		fmt.Println("🤖 IA: No hay movimientos disponibles")
		return -1, -1
	}
	
	randomIndex := rand.Intn(len(emptySpots))
	chosen := emptySpots[randomIndex]
	fmt.Printf("🤖 IA elige: %d,%d\n", chosen[0], chosen[1])
	return chosen[0], chosen[1]
}

func (ai *AI) findWinningMove(board *Board, player string) (int, int, bool) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board.GetCell(i, j) == " " {
				// Probar este movimiento
				board.MakeMove(i, j, player)
				if board.CheckWinner() == player {
					// Deshacer el movimiento
					board.MakeMove(i, j, " ")
					return i, j, true
				}
				// Deshacer el movimiento
				board.MakeMove(i, j, " ")
			}
		}
	}
	return 0, 0, false
}
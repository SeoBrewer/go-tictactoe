package main

import (
	"fmt"
	"math/rand"
	"time"

	"tictactoe/game"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	
	fmt.Println("🎮 ¡Bienvenido al Tic-Tac-Toe!")
	fmt.Println("💡 Juegas como X, la IA inteligente es O")
	fmt.Println("📝 Instrucciones: Ingresa tu movimiento como 'fila,columna' (ej: 1,2)")
	fmt.Println("📍 Las posiciones van del 0 al 2")

	for {
		playGame()
		
		if !askPlayAgain() {
			fmt.Println("🎮 ¡Gracias por jugar!")
			break
		}
	}
}

func playGame() {
	board := game.NewBoard()
	player := game.NewPlayer("X", "Humano")
	ai := game.NewAI("O")
	
	currentPlayerIsHuman := true
	gameOver := false

	for !gameOver {
		board.Print()

		var row, col int
		var err error

		if currentPlayerIsHuman {
			// Human player's turn
			for {
				row, col, err = player.GetMove()
				if err != nil {
					fmt.Printf("❌ Error: %v\n", err)
					continue
				}

				if !board.MakeMove(row, col, player.GetSymbol()) {
					fmt.Println("❌ Posición inválida o ya ocupada. Intenta otra vez.")
					continue
				}
				break
			}
		} else {
			// Check if there are moves available before AI's turn
			if board.IsFull() {
				fmt.Println("🤝 ¡Es un empate!")
				gameOver = true
				break
			}
			
			// AI's turn
			row, col = ai.GetMove(board)
			if row == -1 && col == -1 {
				// No moves available
				fmt.Println("🤝 ¡Es un empate!")
				gameOver = true
				break
			}
			board.MakeMove(row, col, "O")
		}

		// Check for winner
		winner := board.CheckWinner()
		if winner != "" {
			board.Print()
			if winner == "X" {
				fmt.Println("🎉 ¡Increíble! ¡Ganaste a la IA inteligente!")
			} else {
				fmt.Println("🤖 ¡La IA te ha derrotado!")
			}
			gameOver = true
			break
		}

		// Check for tie (after each move)
		if board.IsFull() {
			board.Print()
			fmt.Println("🤝 ¡Es un empate!")
			gameOver = true
			break
		}

		// Switch turn
		currentPlayerIsHuman = !currentPlayerIsHuman
	}
}

func askPlayAgain() bool {
	var input string
	fmt.Print("¿Quieres jugar otra vez? (s/n): ")
	fmt.Scanln(&input)
	
	return input == "s" || input == "S" || input == "si" || input == "Si" || input == "SI"
}
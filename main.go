package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func printBoard(board [3][3]string) {
    fmt.Println("\n  0   1   2")
    for i := 0; i < 3; i++ {
        fmt.Printf("%d ", i)
        for j := 0; j < 3; j++ {
            fmt.Printf(" %s ", board[i][j])
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

func getMove(player string) (int, int) {
    reader := bufio.NewReader(os.Stdin)
    fmt.Printf("Jugador %s, ingresa tu movimiento (fila,columna): ", player)
    
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)
    parts := strings.Split(input, ",")
    
    row, _ := strconv.Atoi(parts[0])
    col, _ := strconv.Atoi(parts[1])
    
    return row, col
}

// Función para verificar si hay ganador
func checkWinner(board [3][3]string) string {
    // Verificar filas
    for i := 0; i < 3; i++ {
        if board[i][0] != " " && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
            return board[i][0]
        }
    }
    
    // Verificar columnas
    for j := 0; j < 3; j++ {
        if board[0][j] != " " && board[0][j] == board[1][j] && board[1][j] == board[2][j] {
            return board[0][j]
        }
    }
    
    // Verificar diagonales
    if board[0][0] != " " && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
        return board[0][0]
    }
    
    if board[0][2] != " " && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
        return board[0][2]
    }
    
    return "" // No hay ganador
}

func main() {
    var board [3][3]string
    currentPlayer := "X"
    
    // Llenar con espacios
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            board[i][j] = " "
        }
    }
    
    // Loop principal del juego
    for {
        printBoard(board)
        
        row, col := getMove(currentPlayer)
        board[row][col] = currentPlayer
        
        // Verificar ganador
        winner := checkWinner(board)
        if winner != "" {
            printBoard(board)
            fmt.Printf("🎉 ¡Jugador %s gana!\n", winner)
            break
        }
        
        // Cambiar jugador
        if currentPlayer == "X" {
            currentPlayer = "O"
        } else {
            currentPlayer = "X"
        }
    }
}
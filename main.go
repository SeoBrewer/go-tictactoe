package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "strings"
    "time"
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

// Verificar si alguien puede ganar en el próximo movimiento
func findWinningMove(board [3][3]string, player string) (int, int, bool) {
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if board[i][j] == " " {
                // Probar este movimiento
                board[i][j] = player
                if checkWinner(board) == player {
                    board[i][j] = " " // Deshacer
                    return i, j, true
                }
                board[i][j] = " " // Deshacer
            }
        }
    }
    return 0, 0, false
}

// IA inteligente
func getSmartAIMove(board [3][3]string) (int, int) {
    fmt.Println("🧠 IA está pensando estratégicamente...")
    time.Sleep(1 * time.Second)
    
    // 1. ¿Puede ganar la IA?
    if row, col, canWin := findWinningMove(board, "O"); canWin {
        fmt.Printf("🎯 IA va por la victoria: %d,%d\n", row, col)
        return row, col
    }
    
    // 2. ¿Necesita bloquear al jugador?
    if row, col, mustBlock := findWinningMove(board, "X"); mustBlock {
        fmt.Printf("🛡️ IA bloquea tu jugada: %d,%d\n", row, col)
        return row, col
    }
    
    // 3. Tomar el centro si está libre
    if board[1][1] == " " {
        fmt.Printf("🎯 IA toma el centro: 1,1\n")
        return 1, 1
    }
    
    // 4. Movimiento aleatorio en posiciones disponibles
    var emptySpots [][2]int
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if board[i][j] == " " {
                emptySpots = append(emptySpots, [2]int{i, j})
            }
        }
    }
    
    randomIndex := rand.Intn(len(emptySpots))
    chosen := emptySpots[randomIndex]
    fmt.Printf("🤖 IA elige: %d,%d\n", chosen[0], chosen[1])
    return chosen[0], chosen[1]
}

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
    
    return ""
}

func main() {
    rand.Seed(time.Now().UnixNano())
    
    var board [3][3]string
    currentPlayer := "X"
    
    // Llenar con espacios
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            board[i][j] = " "
        }
    }
    
    fmt.Println("🎮 Juegas como X, la IA inteligente es O")
    fmt.Println("💡 La IA ahora puede ganar y bloquear tus jugadas")
    
    for {
        printBoard(board)
        
        var row, col int
        
        if currentPlayer == "X" {
            row, col = getMove(currentPlayer)
        } else {
            row, col = getSmartAIMove(board)
        }
        
        board[row][col] = currentPlayer
        
        winner := checkWinner(board)
        if winner != "" {
            printBoard(board)
            if winner == "X" {
                fmt.Println("🎉 ¡Increíble! ¡Ganaste a la IA inteligente!")
            } else {
                fmt.Println("🤖 ¡La IA te ha derrotado!")
            }
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
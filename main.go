package main

import "fmt"

// Función para mostrar el tablero bonito
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

func main() {
    // Crear tablero 3x3 vacío
    var board [3][3]string
    
    // Llenar con espacios
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            board[i][j] = " "
        }
    }
    
    // Mostrar tablero bonito
    printBoard(board)
    
    // Poner una X de prueba
    board[1][1] = "X"
    board[0][2] = "O"
    
    // Mostrar otra vez
    printBoard(board)
}
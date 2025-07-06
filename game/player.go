package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Player struct {
	symbol string
	name   string
}

func NewPlayer(symbol, name string) *Player {
	return &Player{
		symbol: symbol,
		name:   name,
	}
}

func (p *Player) GetMove() (int, int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Jugador %s (%s), ingresa tu movimiento (fila,columna): ", p.name, p.symbol)

	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, 0, err
	}

	input = strings.TrimSpace(input)
	parts := strings.Split(input, ",")

	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("formato inválido. Usa: fila,columna (ej: 1,2)")
	}

	row, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return 0, 0, fmt.Errorf("fila debe ser un número")
	}

	col, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return 0, 0, fmt.Errorf("columna debe ser un número")
	}

	return row, col, nil
}

func (p *Player) GetSymbol() string {
	return p.symbol
}

func (p *Player) GetName() string {
	return p.name
}
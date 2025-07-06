# Tic-Tac-Toe (Go)

A simple and fun Tic-Tac-Toe game written in Go, where you play as **X** against a smart AI (**O**).  
Play in your terminal and try to beat the AI!

## Features

- Play as X, AI plays as O
- Smart AI (tries to win, blocks you, takes center, or picks randomly)
- Friendly terminal interface with emoji feedback
- Replay option after each game

## How to Play

1. **Clone the repo:**
   ```bash
   git clone https://github.com/SeoBrewer/go-tictactoe
   cd go-tictactoe
   ```

2. **Build and run:**
   ```bash
   cd cmd
   go run main.go
   ```

3. **Gameplay:**
   - Enter your move as `row,column` (for example: `1,2`)
   - Rows and columns are numbered from 0 to 2
   - Try to get three in a row before the AI does!

## Example

```
🎮 Welcome to Tic-Tac-Toe!
💡 You play as X, the smart AI is O
📝 Instructions: Enter your move as 'row,column' (e.g. 1,2)
📍 Positions go from 0 to 2

  0   1   2
0   |   |   
  ---|---|---
1   |   |   
  ---|---|---
2   |   |   

Your move: 1,1
...
```

## Requirements

- Go 1.18 or newer

## Project Structure

```
go-tictactoe/
├── cmd/
│   └── main.go        # Entry point
└── game/
    ├── ai.go          # AI logic
    ├── board.go       # Board logic
    └── player.go      # Player logic
```

---

## Credits

**Coded by Army Nougues**, as a Go learning mini project using AI (GitHub Copilot) as coding assistant.

---

## License

MIT

Enjoy the
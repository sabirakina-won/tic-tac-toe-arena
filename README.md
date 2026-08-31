# tic-tac-toe-arena


A terminal tic-tac-toe game written in Go.

The project supports two game modes:

- Two human players can play against each other
- One human player can play against a computer opponent

The game has command-line flags for names, board size, colors, large board rendering, statistics, and the first player.

## Requirements

- Go installed on your computer

## How to run

Clone the repository:

```bash
git clone https://github.com/sabirakina-won/tic-tac-toe-arena.git
```

Open the project folder:

```bash
cd tic-tac-toe-arena
```

Run the game from the repository root:

```bash
go run . --players
```

## Game modes

### Two players

Two human players take turns. X and O enter the number of an empty cell.

```bash
go run . --players
```

### Play against AI

The human player is X. The computer player is O.

```bash
go run . --ai
```

The computer can also start first:

```bash
go run . --ai --first O
```

## Flags

| Flag | Description |
|---|---|
| `--players` | Start a game for two human players |
| `--ai` | Start a game against the computer |
| `--color` | Enable ANSI colors in terminal output |
| `--big` | Render the board with large X and O glyphs |
| `--verbose` | Show extended game statistics |
| `--first X` | X makes the first move |
| `--first O` | O makes the first move |
| `--name A,B` | Set custom player names: X is A and O is B |
| `--size N` | Use an N×N board, where N must be at least 3 |
| `--help` | Show the help message |
| `-h` | Show the help message |

Examples:

```bash
go run . --players --color
```

```bash
go run . --players --big
```

```bash
go run . --players --name Sabira,Zarina
```

```bash
go run . --players --size 4
```

```bash
go run . --players --verbose --first O
```

## Rules

The board cells are numbered from left to right and from top to bottom.

For a 3×3 board:

```text
 1 | 2 | 3
---+---+---
 4 | 5 | 6
---+---+---
 7 | 8 | 9
```

Players enter the number of the cell where they want to place their mark.

- X and O take turns
- A move must be a number inside the board range
- A player cannot choose an occupied cell
- A player wins by filling an entire row, column, or diagonal with the same mark
- A draw happens when the board is full and no player has won
- The program asks whether the players want to play again and keeps session statistics

## AI strategy

In AI mode the computer plays as O.

The AI uses a deterministic rule-based strategy:

1. Win: if O can win on this move, choose that cell
2. Block: if X could win on the next move, block that cell
3. Center: choose cell 5 if it is free
4. Corner: choose the first free corner in this order: 1, 3, 7, 9
5. Side: choose the first free side in this order: 2, 4, 6, 8

The AI does not use minimax or recursion.

## Example game

Command:

```bash
go run . --players
```

Example:

```text
 1 | 2 | 3
---+---+---
 4 | 5 | 6
---+---+---
 7 | 8 | 9

X move: 5
O move: 1
X move: 9
O move: 2
X move: 3
O move: 8
X move: 6

X wins!

=== Stats ===
Games: 1   X: 1   O: 0   Draws: 0
```

## Project structure

```text
tic-tac-toe-arena/
├── go.mod
├── README.md
├── main.go
└── internal/
    ├── ai/
    │   └── ai.go
    ├── board/
    │   └── board.go
    ├── cli/
    │   └── cli.go
    └── game/
        └── game.go
```

## Team

- Sabira Sarybay — GitHub: [@sabirakina-won](https://github.com/sabirakina-won)
- Zarina Temirgali — GitHub: [@zori-ui](https://github.com/zori-ui)
- Ayanat Bakenova — GitHub: # tic-tac-toe-arenа


A terminal tic-tac-toe game written in Go.

The project supports two game modes:

- Two human players can play against each other
- One human player can play against a computer opponent

The game has command-line flags for names, board size, colors, large board rendering, statistics, and the first player.

## Requirements

- Go installed on your computer

## How to run

Clone the repository:

```bash
git clone https://github.com/sabirakina-won/tic-tac-toe-arena.git
```

Open the project folder:

```bash
cd tic-tac-toe-arena
```

Run the game from the repository root:

```bash
go run . --players
```

## Game modes

### Two players

Two human players take turns. X and O enter the number of an empty cell.

```bash
go run . --players
```

### Play against AI

The human player is X. The computer player is O.

```bash
go run . --ai
```

The computer can also start first:

```bash
go run . --ai --first O
```

## Flags

| Flag | Description |
|---|---|
| `--players` | Start a game for two human players |
| `--ai` | Start a game against the computer |
| `--color` | Enable ANSI colors in terminal output |
| `--big` | Render the board with large X and O glyphs |
| `--verbose` | Show extended game statistics |
| `--first X` | X makes the first move |
| `--first O` | O makes the first move |
| `--name A,B` | Set custom player names: X is A and O is B |
| `--size N` | Use an N×N board, where N must be at least 3 |
| `--help` | Show the help message |
| `-h` | Show the help message |

Examples:

```bash
go run . --players --color
```

```bash
go run . --players --big
```

```bash
go run . --players --name Sabira,Zarina
```

```bash
go run . --players --size 4
```

```bash
go run . --players --verbose --first O
```

## Rules

The board cells are numbered from left to right and from top to bottom.

For a 3×3 board:

```text
 1 | 2 | 3
---+---+---
 4 | 5 | 6
---+---+---
 7 | 8 | 9
```

Players enter the number of the cell where they want to place their mark.

- X and O take turns
- A move must be a number inside the board range
- A player cannot choose an occupied cell
- A player wins by filling an entire row, column, or diagonal with the same mark
- A draw happens when the board is full and no player has won
- The program asks whether the players want to play again and keeps session statistics

## AI strategy

In AI mode the computer plays as O.

The AI uses a deterministic rule-based strategy:

1. Win: if O can win on this move, choose that cell
2. Block: if X could win on the next move, block that cell
3. Center: choose cell 5 if it is free
4. Corner: choose the first free corner in this order: 1, 3, 7, 9
5. Side: choose the first free side in this order: 2, 4, 6, 8

The AI does not use minimax or recursion.

## Example game

Command:

```bash
go run . --players
```

Example:

```text
 1 | 2 | 3
---+---+---
 4 | 5 | 6
---+---+---
 7 | 8 | 9

X move: 5
O move: 1
X move: 9
O move: 2
X move: 3
O move: 8
X move: 6

X wins!

=== Stats ===
Games: 1   X: 1   O: 0   Draws: 0
```

## Project structure

```text
tic-tac-toe-arena/
├── go.mod
├── README.md
├── main.go
└── internal/
    ├── ai/
    │   └── ai.go
    ├── board/
    │   └── board.go
    ├── cli/
    │   └── cli.go
    └── game/
        └── game.go
```

## Team

- Sabira Sarybay — GitHub: [@sabirakina-won](https://github.com/sabirakina-won)
- Zarina Temirgali — GitHub: [@zori-ui](https://github.com/zori-ui)
- Ayanat Bakenova — GitHub: [@ayanata-a aya](https://github.com/ayanata-a)
- Anuar Sharipov — GitHub: [@sharp-a](https://github.com/sharp-a)
- Anuar Sharipov — GitHub: 

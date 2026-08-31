package game

import (
	"fmt"
	"io"
	"strings"

	"github.com/sabirakina-won-github-sabwon09/tic-tac-toe-arena/internal/ai"
	"github.com/sabirakina-won-github-sabwon09/tic-tac-toe-arena/internal/board"
	"github.com/sabirakina-won-github-sabwon09/tic-tac-toe-arena/internal/cli"
)

const (
	red   = "\033[91m"
	reset = "\033[0m"
)

type Stats struct {
	Games int
	WinsX int
	WinsO int
	Draws int
}

func Start(config cli.Config) {
	b := board.New(config.Size)
	stats := Stats{}

	for {
		moves, finished := playOneRound(b, config, &stats)

		if !finished {
			return
		}

		printStats(config, stats, moves)

		fmt.Print("Play again? (y/n): ")

		var answer string

		if _, err := fmt.Scan(&answer); err != nil {
			return
		}

		if strings.ToLower(answer) != "y" {
			return
		}

		b.Reset()
	}
}

func playOneRound(
	b *board.Board,
	config cli.Config,
	stats *Stats,
) (int, bool) {

	current := config.First
	moves := 0

	for {
		b.Render(config.Color, config.Big)

		if config.AI && current == "O" {
			cell, reason := ai.ChooseMove(b.Cells())

			if cell == 0 {
				return moves, false
			}

			if config.Verbose {
				fmt.Printf(
					"AI: %s at %d\n",
					reason,
					cell,
				)
			}

			fmt.Printf(
				"%s plays %d\n",
				playerName(config, current),
				cell,
			)

			b.Move(cell, current)
			moves++

		} else {
			fmt.Printf(
				"%s move: ",
				playerName(config, current),
			)

			cell, err := readMove()

			if err == io.EOF {
				return moves, false
			}

			if err != nil {
				printError(
					config.Color,
					fmt.Sprintf(
						"Error: enter a number 1-%d",
						b.Size()*b.Size(),
					),
				)
				continue
			}

			if cell < 1 || cell > b.Size()*b.Size() {
				printError(
					config.Color,
					fmt.Sprintf(
						"Error: enter a number 1-%d",
						b.Size()*b.Size(),
					),
				)
				continue
			}

			if !b.Move(cell, current) {
				printError(
					config.Color,
					fmt.Sprintf(
						"Error: cell %d is taken",
						cell,
					),
				)
				continue
			}

			moves++
		}

		if b.CheckWin(current) {
			b.Render(config.Color, config.Big)

			fmt.Printf(
				"%s wins!\n",
				playerName(config, current),
			)

			stats.Games++

			if current == "X" {
				stats.WinsX++
			} else {
				stats.WinsO++
			}

			return moves, true
		}

		if b.IsFull() {
			b.Render(config.Color, config.Big)

			fmt.Println("Draw!")

			stats.Games++
			stats.Draws++

			return moves, true
		}

		current = switchPlayer(current)
	}
}

func readMove() (int, error) {
	var cell int

	_, err := fmt.Scan(&cell)

	if err == io.EOF {
		return 0, io.EOF
	}

	if err != nil {
		var badInput string
		fmt.Scan(&badInput)

		return 0, err
	}

	return cell, nil
}

func switchPlayer(current string) string {
	if current == "X" {
		return "O"
	}

	return "X"
}

func playerName(config cli.Config, player string) string {
	if player == "X" {
		return config.NameX
	}

	return config.NameO
}

func printStats(
	config cli.Config,
	stats Stats,
	moves int,
) {
	fmt.Println("=== Stats ===")

	fmt.Printf(
		"Games: %d   %s: %d   %s: %d   Draws: %d\n",
		stats.Games,
		config.NameX,
		stats.WinsX,
		config.NameO,
		stats.WinsO,
		stats.Draws,
	)

	if config.Verbose {
		winRateX :=
			float64(stats.WinsX) * 100 /
				float64(stats.Games)

		winRateO :=
			float64(stats.WinsO) * 100 /
				float64(stats.Games)

		fmt.Printf(
			"Moves this game: %d   Win rate — %s: %.0f%%  %s: %.0f%%\n",
			moves,
			config.NameX,
			winRateX,
			config.NameO,
			winRateO,
		)
	}
}

func printError(color bool, message string) {
	if color {
		fmt.Println(red + message + reset)
		return
	}

	fmt.Println(message)
}

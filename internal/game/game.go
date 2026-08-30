
package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ВАШ-GITHUB-АККАУНТ/tic-tac-toe-arena/internal/board"
)

type Stats struct {
	WinsX int
	WinsO int
	Draws int
}

func Start() {
	reader := bufio.NewReader(os.Stdin)
	stats := Stats{}

	for {
		playOneRound(reader, &stats)

		fmt.Printf("\nСтатистика: X = %d, O = %d, Ничьи = %d\n", stats.WinsX, stats.WinsO, stats.Draws)
		fmt.Print("Сыграть ещё раз? (y/n): ")

		answer := readLine(reader)
		if strings.ToLower(strings.TrimSpace(answer)) != "y" {
			fmt.Println("Спасибо за игру!")
			return
		}
	}
}

// playOneRound проводит одну партию от начала до победы/ничьи.
func playOneRound(reader *bufio.Reader, stats *Stats) {
	b := board.New()
	current := "X"

	for {
		b.Render()
		fmt.Printf("Ход игрока %s. Введите строку и столбец (например: 1 2): ", current)

		row, col, err := readMove(reader)
		if err != nil {
			fmt.Println("Ошибка: введите два числа через пробел (строка столбец).")
			continue
		}

		ok := b.Move(row, col, current)
		if !ok {
			fmt.Println("Некорректный ход: клетка занята или координаты вне поля. Попробуйте снова.")
			continue
		}

		if b.CheckWin(current) {
			b.Render()
			fmt.Printf("🎉 Игрок %s победил!\n", current)
			if current == "X" {
				stats.WinsX++
			} else {
				stats.WinsO++
			}
			return
		}

		if b.IsFull() {
			b.Render()
			fmt.Println("Ничья!")
			stats.Draws++
			return
		}

		current = switchPlayer(current)
	}
}

func readMove(reader *bufio.Reader) (int, int, error) {
	line := readLine(reader)
	var row, col int
	n, err := fmt.Sscan(line, &row, &col)
	if err != nil || n != 2 {
		return 0, 0, fmt.Errorf("неверный формат ввода")
	}
	return row, col, nil
}

func readLine(reader *bufio.Reader) string {
	line, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	return line
}

func switchPlayer(current string) string {
	if current == "X" {
		return "O"
	}
	return "X"
}
package board

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	reset     = "\033[0m"
	red       = "\033[91m"
	blue      = "\033[94m"
	greenBold = "\033[1;32m"
	dim       = "\033[2m"
)

type Board struct {
	size        int
	cells       []string
	winningLine []int
}

func New(size int) *Board {
	return &Board{
		size:  size,
		cells: make([]string, size*size),
	}
}

func (b *Board) Size() int {
	return b.size
}

func (b *Board) Cells() []string {
	copyOfCells := make([]string, len(b.cells))
	copy(copyOfCells, b.cells)
	return copyOfCells
}

func (b *Board) Move(cell int, player string) bool {
	if cell < 1 || cell > len(b.cells) {
		return false
	}

	index := cell - 1

	if b.cells[index] != "" {
		return false
	}

	b.cells[index] = player
	b.winningLine = nil

	return true
}

func (b *Board) CheckWin(player string) bool {
	b.winningLine = nil

	for row := 0; row < b.size; row++ {
		line := make([]int, b.size)
		win := true

		for col := 0; col < b.size; col++ {
			index := row*b.size + col
			line[col] = index

			if b.cells[index] != player {
				win = false
			}
		}

		if win {
			b.winningLine = line
			return true
		}
	}

	for col := 0; col < b.size; col++ {
		line := make([]int, b.size)
		win := true

		for row := 0; row < b.size; row++ {
			index := row*b.size + col
			line[row] = index

			if b.cells[index] != player {
				win = false
			}
		}

		if win {
			b.winningLine = line
			return true
		}
	}

	line := make([]int, b.size)
	win := true

	for i := 0; i < b.size; i++ {
		index := i*b.size + i
		line[i] = index

		if b.cells[index] != player {
			win = false
		}
	}

	if win {
		b.winningLine = line
		return true
	}

	line = make([]int, b.size)
	win = true

	for i := 0; i < b.size; i++ {
		index := i*b.size + (b.size - 1 - i)
		line[i] = index

		if b.cells[index] != player {
			win = false
		}
	}

	if win {
		b.winningLine = line
		return true
	}

	return false
}

func (b *Board) IsFull() bool {
	for _, cell := range b.cells {
		if cell == "" {
			return false
		}
	}

	return true
}

func (b *Board) Reset() {
	for i := range b.cells {
		b.cells[i] = ""
	}

	b.winningLine = nil
}

func (b *Board) Render(color, big bool) {
	if big {
		b.renderBig(color)
		return
	}

	b.renderNormal(color)
}

func (b *Board) renderNormal(color bool) {
	width := len(strconv.Itoa(len(b.cells)))

	separatorCell := strings.Repeat("-", width+2)
	separator := strings.TrimSuffix(
		strings.Repeat(separatorCell+"+", b.size),
		"+",
	)

	for row := 0; row < b.size; row++ {
		for col := 0; col < b.size; col++ {
			index := row*b.size + col
			value := b.cells[index]

			if value == "" {
				value = strconv.Itoa(index + 1)
			}

			value = fmt.Sprintf("%*s", width, value)
			value = b.addColor(index, value, color)

			fmt.Printf(" %s ", value)

			if col < b.size-1 {
				fmt.Print("|")
			}
		}

		fmt.Println()

		if row < b.size-1 {
			fmt.Println(separator)
		}
	}
}

func (b *Board) renderBig(color bool) {
	separator := strings.TrimSuffix(
		strings.Repeat("-----+", b.size),
		"+",
	)

	for row := 0; row < b.size; row++ {
		for line := 0; line < 3; line++ {
			for col := 0; col < b.size; col++ {
				index := row*b.size + col

				text := b.bigCell(index, line)
				text = b.addColor(index, text, color)

				fmt.Print(text)

				if col < b.size-1 {
					fmt.Print("|")
				}
			}

			fmt.Println()
		}

		if row < b.size-1 {
			fmt.Println(separator)
		}
	}
}

func (b *Board) bigCell(index, line int) string {
	cell := b.cells[index]

	if cell == "X" {
		glyph := []string{
			"X   X",
			"  X  ",
			"X   X",
		}

		return glyph[line]
	}

	if cell == "O" {
		glyph := []string{
			" OOO ",
			"O   O",
			" OOO ",
		}

		return glyph[line]
	}

	if line == 1 {
		number := strconv.Itoa(index + 1)

		left := (5 - len(number)) / 2
		right := 5 - len(number) - left

		return strings.Repeat(" ", left) +
			number +
			strings.Repeat(" ", right)
	}

	return "     "
}

func (b *Board) addColor(index int, text string, color bool) string {
	if !color {
		return text
	}

	if b.isWinningCell(index) && b.cells[index] != "" {
		return greenBold + text + reset
	}

	if b.cells[index] == "X" {
		return red + text + reset
	}

	if b.cells[index] == "O" {
		return blue + text + reset
	}

	return dim + text + reset
}

func (b *Board) isWinningCell(index int) bool {
	for _, winningIndex := range b.winningLine {
		if winningIndex == index {
			return true
		}
	}

	return false
}

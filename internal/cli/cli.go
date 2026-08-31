package cli

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const Usage = `Usage: go run . (--players | --ai) [options]

Modes (exactly one required):
  --players        two human players take turns
  --ai             play against the computer (you are X)

Options:
  --color          enable colored output (default: plain)
  --big            render the board with large glyphs
  --verbose        show extended statistics
  --first X|O      who moves first (default: X)
  --name A,B       custom names: X=A, O=B (e.g. --name Alice,Bob)
  --size N         board is N×N, win = N in a row (default: 3)
  --help, -h       print this help and exit 0
`

type Config struct {
	Players bool
	AI      bool
	Color   bool
	Big     bool
	Verbose bool

	First string
	NameX string
	NameO string
	Size  int
}

func DefaultConfig() Config {
	return Config{
		First: "X",
		NameX: "X",
		NameO: "O",
		Size:  3,
	}
}

func Parse() (Config, bool, error) {
	config := DefaultConfig()
	args := os.Args[1:]

	if len(args) == 0 {
		return config, false, fmt.Errorf("choose exactly one of --players or --ai")
	}

	for i := 0; i < len(args); i++ {
		arg := args[i]

		switch arg {
		case "--help", "-h":
			return config, true, nil

		case "--players":
			config.Players = true

		case "--ai":
			config.AI = true

		case "--color":
			config.Color = true

		case "--big":
			config.Big = true

		case "--verbose":
			config.Verbose = true

		case "--first":
			if i+1 >= len(args) {
				return config, false, fmt.Errorf("--first needs a value: X or O")
			}

			i++
			first := strings.ToUpper(args[i])

			if first != "X" && first != "O" {
				return config, false, fmt.Errorf("--first must be X or O")
			}

			config.First = first

		case "--name":
			if i+1 >= len(args) {
				return config, false, fmt.Errorf("--name needs two names: A,B")
			}

			i++
			names := strings.Split(args[i], ",")

			if len(names) != 2 {
				return config, false, fmt.Errorf("--name must contain two names separated by a comma")
			}

			nameX := strings.TrimSpace(names[0])
			nameO := strings.TrimSpace(names[1])

			if nameX == "" || nameO == "" {
				return config, false, fmt.Errorf("--name must contain two non-empty names")
			}

			config.NameX = nameX
			config.NameO = nameO

		case "--size":
			if i+1 >= len(args) {
				return config, false, fmt.Errorf("--size needs a number")
			}

			i++
			size, err := strconv.Atoi(args[i])

			if err != nil || size < 3 {
				return config, false, fmt.Errorf("--size must be an integer >= 3")
			}

			config.Size = size

		default:
			return config, false, fmt.Errorf("unknown flag: %s", arg)
		}
	}

	if config.Players == config.AI {
		return config, false, fmt.Errorf("choose exactly one of --players or --ai")
	}

	if config.AI && config.Size != 3 {
		return config, false, fmt.Errorf("--ai and --size cannot be combined (AI is 3×3 only)")
	}

	return config, false, nil
}

func PrintUsage() {
	fmt.Print(Usage)
}

func PrintErrorAndUsage(err error) {
	fmt.Printf("Error: %s\n", err)
	PrintUsage()
}

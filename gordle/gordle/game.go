package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Game holds all the information we need to play a game of gordle.
type Game struct {
	reader *bufio.Reader
}

// errInvalidWordLength is returned when
// the guess has the wrong number of characters.
var errInvalidWordLength = fmt.Errorf("invalid guess, word doesn't have the same number of characters as the solution")

// New returns a Game, which can be used to play.
func New(playerInput io.Reader) *Game {
	g := &Game{
		reader: bufio.NewReader(playerInput),
	}

	return g
}

// Play runs the game.
func (g *Game) Play() {
	fmt.Println("Welcome to Gordle")

	// ask for a valid word
	guess := g.ask()

	fmt.Printf("Your guess is: %s\n", string(guess))
}

const solutionLength = 5

// ask reads the input until a valid suggestion is made (and returned).
func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d-character guess:\n", solutionLength)

	var guess []rune
	for {
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(
				os.Stderr,
				"Gordle failed to read your guess: %s\n",
				err.Error(),
			)
			continue
		}

		guess = []rune(string(playerInput))

		err = g.validateGuess(guess)
		if err != nil {
			_, _ = fmt.Fprintf(
				os.Stderr,
				"Your attempt is invalid with Gordle's solution: %s.\n",
				err.Error(),
			)
			continue
		}
		break
	}

	return guess
}

// validateGuess ensures the guess is valid enough.
func (g *Game) validateGuess(guess []rune) error {
	if len(guess) != solutionLength {
		return fmt.Errorf(
			"expected %d, got %d, %w",
			solutionLength,
			len(guess),
			errInvalidWordLength,
		)
	}
	return nil
}

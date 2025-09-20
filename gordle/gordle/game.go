package gordle

import "fmt"

// Game holds all the information we need to play a game of gordle.
type Game struct{}

// New returns a Game, which can be used to play.
func New() *Game {
	g := &Game{}

	return g
}

func (g *Game) Play() {
	fmt.Println("Welcome to Gordle")

	fmt.Println("Enter a gues:")
}

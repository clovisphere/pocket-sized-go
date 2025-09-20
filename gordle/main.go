package main

import (
	"learn-go-pockets/gordle/gordle"
	"os"
)

func main() {
	g := gordle.New(os.Stdin)
	g.Play()
}

package main

import (
	"github.com/huandu/xstrings"
	_ "math/rand"
)
import "./pocket"

func main() {
	pocket.Digit()
	println(xstrings.Shuffle(pocket.City()))
}

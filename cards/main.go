package main

import (
	"fmt"
)

func main() {
	// cards := newDeck()
	// cards.shuffle()
	// cards.print()
	for i := 0; i < 11; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}
}

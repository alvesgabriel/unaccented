package main

import (
	"fmt"

	"github.com/alvesgabriel/unaccented"
)

func main() {
	str := unaccented.Unaccented("árabe")
	fmt.Println(str)
}

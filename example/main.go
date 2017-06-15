package main

import (
	"fmt"

	"unaccented"
)

func main() {
	str := unaccented.Unaccented("árabe")
	fmt.Println(str)
}

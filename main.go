package main

import (
	"fmt"

	"github.com/bgst009/miniature-octo-guide/tool"
)

func main() {
	fmt.Println("Hello world!")
	fmt.Printf("tool.Plus(1, 2): %v\n", tool.Plus(1, 2))
	fmt.Printf("tool.Mult(4, 5): %v\n", tool.Mult(4, 5))
}

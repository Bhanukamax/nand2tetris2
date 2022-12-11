package main

import (
	"fmt"
	"imbmax.com/n2t/vm/parser"
)

func main() {
	fmt.Println("this is the vm", parser.Parse())
	fmt.Println("this is another test", 2)
}

package main

import (
	"fmt"
	"imbmax.com/n2t/vm/log"
	"imbmax.com/n2t/vm/parser"
	"os"
	"strings"
		"bufio"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("this is the vm", parser.Parse())

	fmt.Println("this is another test", len(os.Args))

	argsCount := len(os.Args)
	if argsCount != 2 {
		panic("Correct useage vm {FileName}.vm)")
	}

	fileName := os.Args[1]
	dat, err := os.ReadFile(fileName)

	log.Out(fmt.Sprintf("Openning file for reading!!!! " + fileName))

	check(err)
	fmt.Println(string(dat))
	fmt.Printf("lenght %d\n", len(string(dat)))

	var lines[]string
	sc := bufio.NewScanner(strings.NewReader(string(dat)))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	fmt.Println(lines, len(lines))

	// Reading line by line
	// reader := bufio.NewReader(dat)
	// b, err := reader.Peek(1)
	// check(err)

	// fmt.Printf("%b", b)
}

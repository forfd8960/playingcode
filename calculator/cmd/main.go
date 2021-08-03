package main

import (
	"fmt"
	"github.com/forfd8960/playingcode/calculator"
)

func main() {
	s := "(1 + 2) * (3 / 4) + 100"
	fmt.Println("run calculation")
	cal := &calculator.Calculator{s: s}
	if err := cal.Exec(); err != nil {
		fmt.Printf("calculate %s has error: %v\n", s, err)
	}

	fmt.Println("result: ", cal.Result())
}

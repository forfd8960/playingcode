package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/forfd8960/playingcode/calculator"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("(enter expression)>>")
		scan.Scan()
		calculate(scan.Text())
	}
}

func calculate(exp string) {
	fmt.Println("run calculation of: ", exp)
	cal := calculator.NewCalculator(exp)
	if err := cal.Exec(); err != nil {
		fmt.Printf("calculate %s has error: %v\n", exp, err)
		return
	}

	result, err := cal.Result()
	fmt.Printf("result: %f, err: %v\n", result, err)
}

package year2022

import (
	"fmt"
	"testing"
)

func TestGen2(t *testing.T) {
	s := gen2()
	fmt.Println(s)
}

func TestGen0(t *testing.T) {
	s := gen0()
	fmt.Println(s)
}

func TestGen2022(t *testing.T) {
	result := Gen2022()
	fmt.Println(result)
}

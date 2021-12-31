package year2022

import (
	"fmt"
	"strings"
)

const (
	horizontal = "-"
	vertical   = "|"
)

func Gen2022() string {
	year2022 := []string{}
	year2022 = append(year2022, gen2())
	year2022 = append(year2022, gen0())
	year2022 = append(year2022, gen2())
	year2022 = append(year2022, gen2())
	return strings.Join(year2022, "   ")
}

func gen2() string {
	two := make([][]string, 9)
	for i := range two {
		two[i] = make([]string, 10)
	}

	for i := 0; i < 10; i++ {
		two[0][i] = horizontal
	}

	for i := 1; i <= 3; i++ {
		for j := 0; j < 10; j++ {
			if j == 9 {
				two[i][j] = vertical
			} else {
				two[i][j] = " "
			}
		}
	}
	for i := 0; i < 10; i++ {
		two[4][i] = horizontal
	}

	for i := 5; i <= 7; i++ {
		for j := 0; j < 10; j++ {
			if j == 0 {
				two[i][j] = vertical
			} else {
				two[i][j] = " "
			}
		}
	}

	for i := 0; i < 10; i++ {
		two[8][i] = horizontal
	}

	// fmt.Printf("two: %+v\n", two)

	result := []string{}
	for i := range two {
		result = append(result, strings.Join(two[i], ""))
		fmt.Println(strings.Join(two[i], ""))
	}

	twoStr := strings.Join(result, "\n")
	return strings.TrimSuffix(twoStr, "\n")
}

func gen0() string {
	zero := make([][]string, 9)
	for i := range zero {
		zero[i] = make([]string, 10)
	}

	for i := 0; i < 10; i++ {
		zero[0][i] = horizontal
	}

	for i := 1; i <= 7; i++ {
		for j := 0; j < 10; j++ {
			if j == 9 || j == 0 {
				zero[i][j] = vertical
			} else {
				zero[i][j] = " "
			}
		}
	}

	for i := 0; i < 10; i++ {
		zero[8][i] = horizontal
	}

	// fmt.Printf("zero: %+v\n", zero)

	result := []string{}
	for i := range zero {
		result = append(result, strings.Join(zero[i], ""))
		fmt.Println(strings.Join(zero[i], ""))
	}

	zeroStr := strings.Join(result, "\n")
	return strings.TrimSuffix(zeroStr, "\n")
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type snafuNrs []string

func snafuToDecimal(s string) int{
	result := 0
	for p := 0; len(s)-1-p >= 0; p++ {
		switch s[len(s)-1-p] {
		case '0':
		case '1':
			result += int(math.Pow(5, float64(p)))
		case '2':
			result += 2 * int(math.Pow(5, float64(p)))
		case '-':
			result -= int(math.Pow(5, float64(p)))
		case '=':
			result -= 2 * int(math.Pow(5, float64(p)))
		}
	}
	return result
}

func decimalToSnafu(nr int) string {
	s := ""
	carry := false
	for nr != 0 || carry {
		digit := nr % 5
		if carry {
			digit++
		}
		nr /= 5
		digStr := strconv.Itoa(digit)
		switch digit{
		case 5:
			s = "0" + s
			carry = true
		case 0, 1, 2:
			s = digStr + s
			carry = false
		case 4:
			s = "-" + s
			carry = true
		case 3:
			s = "=" + s
			carry = true
		}
	}
	return s
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var nrs snafuNrs

	for scanner.Scan() {
			t := scanner.Text()
			nrs = append(nrs, t)
	}

	sum := 0
	for _, nr := range nrs {
		sum += snafuToDecimal(nr)
	}
	fmt.Println("part 1:", decimalToSnafu(sum))
}
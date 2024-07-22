package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var a, action string
	fmt.Scanf("%q%s", &a, &action)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	b := scanner.Text()

	x := []rune(a)
	y := []rune(b)

	if b == "" {
		panic("неправильный ввод первой строки!")
	} else if len(x) <= 10 && len(y) <= 12 {
		c := rune(b[0])
		d := rune(b[len(b)-1])

		switch action {
		case "+":
			if c == '"' && d == '"' {
				b = b[1 : len(b)-1]
				fmt.Printf("\"%s%s\"", a, b)
			} else {
				panic("неправильный ввод второй строки!")
			}
		case "-":
			if c == '"' && d == '"' {
				b = b[1 : len(b)-1]
				a = strings.ReplaceAll(a, b, "")
				fmt.Printf("%q", a)
			} else {
				panic("неправильный ввод первой строки!")
			}
		case "*":
			if b == "1" || b == "2" || b == "3" || b == "4" || b == "5" || b == "6" || b == "7" || b == "8" || b == "9" || b == "10" {
				h, _ := strconv.Atoi(b)
				x = []rune(strings.Repeat(a, h))
				if len(x) > 40 {
					x = append(x[:40])
					fmt.Printf("\"%s...\"", string(x))
				} else {
					fmt.Printf("%q", string(x))
				}
			} else {
				panic("некорректное число!")
			}
		case "/":
			if b == "1" || b == "2" || b == "3" || b == "4" || b == "5" || b == "6" || b == "7" || b == "8" || b == "9" || b == "10" {
				h, _ := strconv.Atoi(b)
				x = x[:len(x)/h]
				fmt.Printf("%q", string(x))
			} else {
				panic("некорректное число!")
			}
		default:
			panic("некорректная операция!")
		}
	} else {
		panic("некоректная длина введенных строк!")
	}
}

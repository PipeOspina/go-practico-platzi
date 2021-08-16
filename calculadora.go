package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

func (*calc) operate(in string, operator string) int {
	cleanEntrace := strings.Split(in, operator)
	operador1 := parse(cleanEntrace[0])
	operador2 := parse(cleanEntrace[1])
	switch operator {
	case "+":
		return (operador1 + operador2)
	case "-":
		return (operador1 - operador2)
	case "*":
		return (operador1 * operador2)
	case "/":
		return (operador1 / operador2)
	default:
		fmt.Println(operator, "no soportado")
		return 0
	}
}

func parse(in string) int {
	operator, err := strconv.Atoi(in)
	if err != nil {
		fmt.Println(err)
	}
	return operator
}

func read() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	in := read()
	operator := read()
	c := calc{}
	total := c.operate(in, operator)
	fmt.Println(total)
}

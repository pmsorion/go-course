package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type calc struct{}

func (calc) operate(entrada string, operador string) int {
	entradaLipia := strings.Split(entrada, operador)
	operador1 := parsear(entradaLipia[0])
	operador2 := parsear(entradaLipia[1])
	switch operador {
	case "+":
		fmt.Println(operador1 + operador2)
		return operador1 + operador2
	case "-":
		fmt.Println(operador1 - operador2)
		return operador1 - operador2
	case "*":
		fmt.Println(operador1 * operador2)
		return operador1 * operador2
	case "/":
		if operador2 == 0 {
			fmt.Println("no se puede dividir por 0")
			return 0
		} else {
			fmt.Println(operador1 / operador2)
			return operador1 / operador2
		}
	default:
		fmt.Println("el operador->", operador, "no esta soportado")
		return 0
	}
}

func parsear(entrada string) int {
	operador, _ := strconv.Atoi(entrada)
	return operador
}

func leerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main()  {
	entrada := leerEntrada()
	operador := leerEntrada()
	c := calc{}
	fmt.Println(c.operate(entrada, operador))
}


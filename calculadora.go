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
	case "-":
		fmt.Println(operador1 - operador2)
	case "*":
		fmt.Println(operador1 * operador2)
	case "/":
		if operador2 == 0 {
			fmt.Println("no se puede dividir por 0")
		} else {
			fmt.Println(operador1 / operador2)
		}
	default:
		fmt.Println("el operador->", operador, "no esta soportado")
	}
	
}

func parsear(entrada string) int {
	operador, _ := strconv.Atoi(entrada)
	return operador
}

func leerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operacion := scanner.Text()
}

func main()  {
	entrada := leerEntrada()
	operador := leerEntrada()

	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(operador1)
	}

	operador2, err2 := strconv.Atoi(valores[1])
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(operador2)
	}

	if err1 == nil && err2 == nil {

	} else {
		fmt.Println("no se pudo hacer la operacion")
	}
}


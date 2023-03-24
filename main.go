package main

import (
	"fmt"
	"projetoCalc/calc"
)

func main() {
	var num1, num2 float32
	var sinal, reiniciar string

	fmt.Println("Bem-vindo(a) a calculadora!")
	fmt.Println("Digite o primeiro valor:")
	fmt.Scanln(&num1)
	fmt.Println("Digite o segundo valor:")
	fmt.Scanln(&num2)

	for sinal == "" {
		fmt.Println("Digite a operação (Ex: +, -, /, *)")
		fmt.Scanln(&sinal)
		x := calc.ValidarCalculo(num1, num2, sinal)
		fmt.Println(x)
	}

	for reiniciar != "S" && reiniciar != "s" && reiniciar != "N" && reiniciar != "n" {

		fmt.Println("Deseja realizar outra operação? (S/N)")
		fmt.Scanln(&reiniciar)

		if reiniciar == "S" || reiniciar == "s" {
			main()
		} else if reiniciar == "N" || reiniciar == "n" {
			fmt.Println("Obrigado por usar a calculadora!")
			break
		} else {
			fmt.Println("Utilize S ou N por gentileza!")
		}
	}

}

package greetings

import "fmt"

// Hello retorna uma mensagem de saudação para a pessoa
func Hello(name string) string {
	/*
	 * Para declarar uma variável de forma rápida, é utilizado o sinal :=
	 *
	 * message := fmt.Sprintf("Hi, %v. Welcome!", name)
	 * return message
	 *
	 * mas também pode ser utilizado como o trecho abaixo
	 */

	var message string
	message = fmt.Sprintf("Hi, %v. Welcome!", name)

	return message
}

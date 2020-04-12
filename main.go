package main

import (
	"bank/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

/*
Uma interface é a definição de um conjunto de métodos comuns a um ou mais tipos. É o que permite a criação de tipos polimórficos em Go.
Podemos reaproveitar métodos ou funções entre tipos utilizando interface.
	Nosso projeto é um exemplo disso. Reutilizamos o método Sacar para pagar os boletos com ambas as contas.
*/
type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 60)

	fmt.Println(contaDoDenis.ObterSaldo())

	contaDaLuiza := contas.ContaCorrente{}
	contaDaLuiza.Depositar(500)
	PagarBoleto(&contaDaLuiza, 400)

	fmt.Println(contaDaLuiza.ObterSaldo())
}

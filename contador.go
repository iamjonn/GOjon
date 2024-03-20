package main

// time é uma biblioteca que contém funções para manipulação de tempo
import (
	"fmt"
	"time"
)

//Função que conta até count que é um número inteiro
func contando(count int) {
	for i := 0; i < count; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
	
}

//Função principal
func main() {
	// funciona sequencialmente a nao ser que seja chamada uma go routine para exucutar o paralelismo
 go contando(5) //rota 2
 go contando(5) // rotina 3 -> 2kb -> 1mb
	contando(2)
}

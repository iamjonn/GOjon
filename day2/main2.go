package main

import "fmt"

// canal é uma forma de comunicação entre goroutines
func main() {
	ch := make(chan int) // cria um canal
	go publish(ch)
	consume(ch)
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i // envia uma mensagem para o canal
	}
	close(ch)
}

// fica sempre verificando se chegou uma nova mensagem no canal
func consume(ch chan int) {
	for i := range ch {
		fmt.Println("cansei", i)
		//depois de ler a mensagem, o canal é esvaiziado
	}
}
package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for i := range data {
		time.Sleep(time.Second) //sleep de 1 segundo
		fmt.Printf("worker %d received %d\n", workerId, i) //%d guarda a posiçao na string para ser substituida por um inteiro que é passado depois da string
	}

}

//usando paralelismo para processar dados
func main() {
	data := make(chan int) //cria um canal	
	qtdWorkers := 10 //quantidade de workers

	for i := 0; i < qtdWorkers; i++ {
		go worker(i, data) //cria uma goroutine para cada worker
	}
	

	for i := 0; i < 100; i++ {
		data <- i //envia uma mensagem para o canal
	}

}
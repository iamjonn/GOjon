package main

import "fmt"

// funcao main é uma thread
func main() {
	canal := make(chan string)
//thread 2
  go func() {
	    canal <- "Olá"
	}()

	msg := <-canal //se o canal estiver cheeio, esvazia ele!!
	fmt.Println(msg)
}
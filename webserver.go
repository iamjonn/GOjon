package main

//http é uma biblioteca que contém funções para criação de servidores web
import (
	"net/http"
)

//Função principal
func main () {
	//http.ListenAndServe é uma função que cria um servidor web
	http.ListenAndServe(":3000", nil)
	//primeiro rodaar o codigo e depois abrir outro terminal e digitar curl localhost:8080
}


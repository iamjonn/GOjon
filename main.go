
//Declaração do pacote
package main

//fmt é uma biblioteca que contém funções para entrada e saída de dados
import "fmt"
//struct é um tipo de dado que pode conter outros tipos de dados
type Person struct {
	name string
	age int
}

//Toda função em Go é um método de alguma coisa
func (p Person) Walk(){
	fmt.Println(p.name, "is walking")
}

//Função principal
func main() {
	//call the function
	var a int 
	a = 10
	fmt.Println(a)
	var pessoa Person
	pessoa.name = "Lucas"
	pessoa.age = 20
	pessoa.Walk()
}
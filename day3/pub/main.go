package main

import (
	"context"
	"flag"

	"cloud.google.com/go/pubsub"
) // bliblioteca do google cloud pubsub

func main() {
	ctx := context.Background() // ctx é o contexto da aplicação

	projectid := flag.String("projectid", "", "project id") // flag para pegar o id do projeto
	topicNAme := flag.String("topic", "", "Nome do tópico") // flag para pegar o nome do tópico
	msg := flag.String("msg", "", "Mensagem a ser enviada") // flag para pegar a mensagem a ser enviada
	flag.Parse() // parse das flags que serve para pegar os valores das flags

	c, err := pubsub.NewClient(ctx, *projectid) // c é o cliente do pubsub
	if err != nil {
		panic(err) // se der erro, o programa para
	}
	defer c.Close() // fecha o cliente

	topic := c.Topic(*topicNAme) // topic é o tópico
	exist, err := topic.Exists(ctx) // exist é uma variável que verifica se o tópico existe
	if err != nil {
		panic(err) // se der erro, o programa para
	}
	if !exist {
		topic, err = c.CreateTopic(ctx, *topicNAme) // cria o tópico
		if err != nil {
			panic(err) // se der erro, o programa para
		}
	}
	result := topic.Publish(ctx, &pubsub.Message{ // publica a mensagem no tópico
		Data: []byte(*msg),
	})

	_, err = result.Get(ctx) // pega o resultado e verifica se deu erro
	if err != nil {
		panic(err)
	}
}

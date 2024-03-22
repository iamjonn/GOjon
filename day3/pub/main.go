package main

import (
	"context"
	"flag"

	"cloud.google.com/go/pubsub"
) // bliblioteca do google cloud pubsub

func main() {
	ctx := context.Background()

	projectid := flag.String("projectid", "", "project id")
	topicNAme := flag.String("topic", "", "Nome do tópico")
	msg := flag.String("msg", "", "Mensagem a ser enviada")
	flag.Parse()

	c, err := pubsub.NewClient(ctx, *projectid)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	topic := c.Topic(*topicNAme)
	exist, err := topic.Exists(ctx)
	if err != nil {
		panic(err)
	}
	if !exist {
		topic, err = c.CreateTopic(ctx, *topicNAme)
		if err != nil {
			panic(err)
		}
	}
	result := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(*msg),
	})

	_, err = result.Get(ctx)
	if err != nil {
		panic(err)
	}
}

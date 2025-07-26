package rpc

import "gfAdmin/internal/client"

type Message struct {
	Client  *client.Client
	Content []byte
}

func NewMessage(client *client.Client, content []byte) *Message {
	return &Message{
		Client:  client,
		Content: content,
	}
}

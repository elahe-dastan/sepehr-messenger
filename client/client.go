package client

import (
	"alibaba/protocol"
	"log"

	"google.golang.org/grpc"
)

func New() protocol.SimpleChatClient {
	conn, err := grpc.Dial("localhost:1373", grpc.WithInsecure())

	if err != nil {
		log.Println(err)
	}

	return protocol.NewSimpleChatClient(conn)
}
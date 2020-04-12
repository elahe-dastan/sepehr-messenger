package client

import (
	"log"

	"github.com/elahe-dastan/interview/protocol"

	"google.golang.org/grpc"
)

func New() protocol.ChatClient {
	conn, err := grpc.Dial("localhost:1373", grpc.WithInsecure())

	if err != nil {
		log.Println(err)
	}

	return protocol.NewChatClient(conn)
}